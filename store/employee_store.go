package store

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mcgtrt/azure-graphql/graph/model"
	"github.com/mcgtrt/azure-graphql/types"
	"github.com/mcgtrt/azure-graphql/util"
	_ "github.com/microsoft/go-mssqldb"
	"golang.org/x/crypto/bcrypt"
)

const pwCost = 12

// These are the standard connection strings. If you provide the .env
// file, they will be automatically changed from the init() function
var (
	server   = "localhost"
	port     = "1433"
	user     = "sa"
	password = "superStrong(!)Password"
	database = "master"
)

type EmployeeStorer interface {
	GetEmployeeByID(context.Context, string) (*model.Employee, error)
	GetEmployeeByEmail(context.Context, string) (*types.AuthEmployee, error)
	GetEmployees(context.Context, *int, *int, *string) ([]*model.Employee, error)
	InsertEmployee(context.Context, *model.CreateEmployeeParams) (*model.Response, error)
	UpdateEmployee(context.Context, *model.UpdateEmployeeParams) (*model.Response, error)
	DeleteEmployee(context.Context, string) (*model.Response, error)
}

type AzureEmployeeStore struct {
	db *sql.DB
}

func NewAzureEmployeeStore() (EmployeeStorer, error) {
	connString := fmt.Sprintf(
		"server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server,
		user,
		password,
		port,
		database,
	)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	return &AzureEmployeeStore{
		db: db,
	}, nil
}

func (s *AzureEmployeeStore) GetEmployeeByID(ctx context.Context, id string) (*model.Employee, error) {
	if err := s.db.PingContext(ctx); err != nil {
		return nil, err
	}

	tsql := fmt.Sprintf(
		`SELECT EmployeeID, FirstName, LastName, Username, Email, DOB, DepartmentID, Position FROM AzureQl.Employee 
		WHERE EmployeeID = '%s';`,
		id,
	)
	rows, err := s.db.QueryContext(ctx, tsql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := &model.Employee{}
	for rows.Next() {
		if err := rows.Scan(
			&e.EmployeeID,
			&e.FirstName,
			&e.LastName,
			&e.Username,
			&e.Email,
			&e.Dob,
			&e.DepartmentID,
			&e.Position,
		); err != nil {
			return nil, err
		}
	}
	return e, nil
}

func (s *AzureEmployeeStore) GetEmployeeByEmail(ctx context.Context, email string) (*types.AuthEmployee, error) {
	if err := s.db.PingContext(ctx); err != nil {
		return nil, err
	}

	tsql := fmt.Sprintf("SELECT EmployeeID, Email, EncryptedPassword FROM AzureQl.Employee WHERE Email = '%s';", email)
	rows, err := s.db.QueryContext(ctx, tsql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	e := &types.AuthEmployee{}
	for rows.Next() {
		if err := rows.Scan(
			&e.ID,
			&e.Email,
			&e.EncryptedPassword,
		); err != nil {
			return nil, err
		}
	}
	return e, nil
}

func (s *AzureEmployeeStore) GetEmployees(ctx context.Context, pLimit *int, pPage *int, sortBy *string) ([]*model.Employee, error) {
	if err := s.db.PingContext(ctx); err != nil {
		return nil, err
	}

	orderBy := ""
	if sortBy != nil {
		if err := util.ValidateSortBy(sortBy); err != nil {
			return nil, err
		} else {
			orderBy = fmt.Sprintf(" ORDER BY %s", *sortBy)
		}
	}

	var (
		limit      = 25
		page       = 1
		pagination = ""
	)
	if pLimit != nil && *pLimit > 0 {
		limit = *pLimit
	}
	if pPage != nil && *pPage > 0 {
		page = *pPage
	}
	if page > 1 {
		offset := fmt.Sprintf(" OFFSET %d ROWS", (page-1)*limit)
		fetch := fmt.Sprintf(" FETCH NEXT %d ROWS ONLY", limit)
		pagination = offset + fetch
	}

	tsql := `
		SELECT EmployeeID, FirstName, LastName, Username, Email, DOB, DepartmentID, Position 
		FROM AzureQl.Employee` + orderBy + pagination + ";"
	rows, err := s.db.QueryContext(ctx, tsql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*model.Employee
	for rows.Next() {
		e := &model.Employee{}
		if err := rows.Scan(
			&e.EmployeeID,
			&e.FirstName,
			&e.LastName,
			&e.Username,
			&e.Email,
			&e.Dob,
			&e.DepartmentID,
			&e.Position,
		); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}

	return employees, nil
}

func (s *AzureEmployeeStore) InsertEmployee(ctx context.Context, params *model.CreateEmployeeParams) (*model.Response, error) {
	if err := util.ValidateCreateEmployeeParams(params); err != nil {
		return nil, err
	}
	if err := s.db.PingContext(ctx); err != nil {
		return nil, err
	}

	tsql := `
		INSERT INTO AzureQl.Employee (FirstName, LastName, Username, EncryptedPassword, Email, DOB, DepartmentID, Position) 
		VALUES (@FirstName, @LastName, @Username, @EncryptedPassword, @Email, @DOB, @DepartmentID, @Position);
		select isNull(SCOPE_IDENTITY(), -1);
	`

	state, err := s.db.Prepare(tsql)
	if err != nil {
		return nil, err
	}
	defer state.Close()

	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), pwCost)
	if err != nil {
		return nil, err
	}

	var (
		row = state.QueryRowContext(
			ctx,
			sql.Named("FirstName", params.FirstName),
			sql.Named("LastName", params.LastName),
			sql.Named("Username", params.Username),
			sql.Named("EncryptedPassword", string(encpw)),
			sql.Named("Email", params.Email),
			sql.Named("DOB", params.Dob),
			sql.Named("DepartmentID", params.DepartmentID),
			sql.Named("Position", params.Position),
		)
		newID int64
	)

	if err := row.Scan(&newID); err != nil {
		return nil, err
	}

	return &model.Response{
		Status: http.StatusOK,
		Msg:    fmt.Sprintf("%d", newID),
	}, nil
}

func (s *AzureEmployeeStore) UpdateEmployee(ctx context.Context, params *model.UpdateEmployeeParams) (*model.Response, error) {
	if err := util.ValidateUpdateEmploteeParams(params); err != nil {
		return nil, err
	}
	if err := s.db.PingContext(ctx); err != nil {
		return nil, err
	}

	var (
		query = util.MakeQueryFromUpdateParams(params)
		tsql  = fmt.Sprintf("UPDATE AzureQl.Employee SET %s WHERE EmployeeID ='%s'", query, params.EmployeeID)
	)

	_, err := s.db.ExecContext(ctx, tsql)
	if err != nil {
		return nil, err
	}

	return &model.Response{
		Status: http.StatusOK,
		Msg:    "updated",
	}, nil
}

func (s *AzureEmployeeStore) DeleteEmployee(ctx context.Context, id string) (*model.Response, error) {
	if err := s.db.PingContext(ctx); err != nil {
		return nil, err
	}

	tsql := "DELETE FROM AzureQl.Employee WHERE EmployeeID = @EmployeeID"
	_, err := s.db.ExecContext(ctx, tsql, sql.Named("EmployeeID", id))
	if err != nil {
		return nil, err
	}

	return &model.Response{
		Status: http.StatusOK,
		Msg:    "deleted",
	}, nil
}

func init() {
	if url := os.Getenv("AZURE_SERVER_URL"); url != "" {
		server = url
	}
	if p := os.Getenv("AZURE_SERVER_PORT"); p != "" {
		port = p
	}
	if u := os.Getenv("AZURE_USERNAME"); u != "" {
		user = u
	}
	if p := os.Getenv("AZURE_PASSWORD"); p != "" {
		password = p
	}
	if db := os.Getenv("AZURE_DBNAME"); db != "" {
		database = db
	}
}

package store

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mcgtrt/azure-graphql/graph/model"
	_ "github.com/microsoft/go-mssqldb"
)

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
	GetEmployees(context.Context) ([]*model.Employee, error)
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

func (s *AzureEmployeeStore) GetEmployeeByID(context.Context, string) (*model.Employee, error) {
	return nil, nil
}

func (s *AzureEmployeeStore) GetEmployees(context.Context) ([]*model.Employee, error) {
	return nil, nil
}

func (s *AzureEmployeeStore) InsertEmployee(context.Context, *model.CreateEmployeeParams) (*model.Response, error) {
	return nil, nil
}

func (s *AzureEmployeeStore) UpdateEmployee(context.Context, *model.UpdateEmployeeParams) (*model.Response, error) {
	return nil, nil
}

func (s *AzureEmployeeStore) DeleteEmployee(context.Context, string) (*model.Response, error) {
	return nil, nil
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

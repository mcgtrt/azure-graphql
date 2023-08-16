package util

import (
	"fmt"
	"strconv"

	"github.com/mcgtrt/azure-graphql/graph/model"
)

const (
	minFirstNameLen = 2
	maxFirstNameLen = 32
	minLastNameLen  = 2
	maxLastNameLen  = 32
	minUsernameLen  = 6
	maxUsernameLen  = 32
	minPasswordLen  = 8
	maxPasswordLen  = 32
	minEmailLen     = 7
	maxEmailLen     = 64
	DOBLen          = 10
	minPositionLen  = 2
	maxPositionLen  = 32
)

// TODO: Instead of validating one field by one, this could return a map[string]string
// with detailed key - value pair with description and a key field
func ValidateCreateEmployeeParams(params *model.CreateEmployeeParams) error {
	if len(params.FirstName) < minFirstNameLen || len(params.FirstName) > maxFirstNameLen {
		return fmt.Errorf("first name should have between %d and %d characters", minFirstNameLen, maxFirstNameLen)
	}
	if len(params.LastName) < minLastNameLen || len(params.LastName) > maxLastNameLen {
		return fmt.Errorf("last name should have between %d and %d characters", minLastNameLen, maxLastNameLen)
	}
	if len(params.Username) < minUsernameLen || len(params.Username) > maxUsernameLen {
		return fmt.Errorf("username should have between %d and %d characters", minUsernameLen, maxUsernameLen)
	}
	if len(params.Email) < minEmailLen || len(params.Email) > maxEmailLen {
		// TODO: Add regex expresion to check if email valid
		return fmt.Errorf("email should have between %d and %d characters", minEmailLen, maxEmailLen)
	}
	if len(params.Password) < minPasswordLen || len(params.Password) > maxPasswordLen {
		// TODO: Add password validator for uppercase, number, special char etc.
		return fmt.Errorf("password should have between %d and %d characters", minPasswordLen, maxPasswordLen)
	}
	if len(params.Dob) != DOBLen {
		// TODO: Add time format validation
		return fmt.Errorf("date of birth should have a format: 1999/09/29")
	}
	id, err := strconv.Atoi(params.DepartmentID)
	if err != nil {
		return fmt.Errorf("invalid department id")
	}
	if id < 1 {
		return fmt.Errorf("department id must be bigger than 0")
	}
	if len(params.Position) < minPositionLen || len(params.Position) > maxPositionLen {
		return fmt.Errorf("position should have between %d and %d characters", minPositionLen, maxPositionLen)
	}

	return nil
}

// TODO: Same map[string]string error handling as for ValidateCreateEmployeeParams
func ValidateUpdateEmploteeParams(params *model.UpdateEmployeeParams) error {
	if len(params.EmployeeID) == 0 {
		return fmt.Errorf("invalid employee id")
	}
	if params.FirstName != nil {
		if len(*params.FirstName) < minFirstNameLen || len(*params.FirstName) > maxFirstNameLen {
			return fmt.Errorf("first name should have between %d and %d characters", minFirstNameLen, maxFirstNameLen)
		}
	}
	if params.LastName != nil {
		if len(*params.LastName) < minLastNameLen || len(*params.LastName) > maxLastNameLen {
			return fmt.Errorf("last name should have between %d and %d characters", minLastNameLen, maxLastNameLen)
		}
	}
	if params.Username != nil {
		if len(*params.Username) < minUsernameLen || len(*params.Username) > maxUsernameLen {
			return fmt.Errorf("username should have between %d and %d characters", minUsernameLen, maxUsernameLen)
		}
	}
	if params.Dob != nil {
		if len(*params.Dob) != DOBLen {
			// TODO: Add time format validation
			return fmt.Errorf("date of birth should have a format: 1999/09/29")
		}
	}
	if params.DepartmentID != nil {
		id, err := strconv.Atoi(*params.DepartmentID)
		if err != nil {
			return fmt.Errorf("invalid department id")
		}
		if id < 1 {
			return fmt.Errorf("department id must be bigger than 0")
		}
	}
	if params.Position != nil {
		if len(*params.Position) < minPositionLen || len(*params.Position) > maxPositionLen {
			return fmt.Errorf("position should have between %d and %d characters", minPositionLen, maxPositionLen)
		}
	}

	return nil
}

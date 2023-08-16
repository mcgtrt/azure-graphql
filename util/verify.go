package util

import (
	"fmt"

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
func VerifyCreateEmployeeParams(params model.CreateEmployeeParams) error {
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
		return fmt.Errorf("email should have between %d and %d characters", minEmailLen, maxEmailLen)
	}
	if len(params.Password) < minPasswordLen || len(params.Password) > maxPasswordLen {
		return fmt.Errorf("password should have between %d and %d characters", minPasswordLen, maxPasswordLen)
	}
	if len(params.Dob) != DOBLen {
		// TODO: Add time format validation
		return fmt.Errorf("Date of birth should have a format: 1999/09/29")
	}
	if params.DepartmentID < 1 {
		return fmt.Errorf("department id must be bigger than 0")
	}
	if len(params.Position) < minPositionLen || len(params.Position) > maxPositionLen {
		return fmt.Errorf("position should have between %d and %d characters", minPositionLen, maxPositionLen)
	}

	return nil
}

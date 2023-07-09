package validator_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yoas0bi/micro-toolkit/validator"
)

var (
	testStruct = Test{}
)

type Test struct {
	RequiredString   string   `validate:"required"`
	RequiredNumber   int      `validate:"required"`
	RequiredMultiple []string `validate:"required"`
}

func TestValidator(t *testing.T) {
	should := assert.New(t)

	if should.NoError(validator.Init()) {
		err := validator.Validate(testStruct)
		fmt.Println(err)
	}

}

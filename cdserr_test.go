package cdserr_test

import (
	"errors"
	"fmt"
	"github.com/cirrusdata/cdserr"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func dummyFn() error {
	return cdserr.NewInvalidParameterError(nil, "Test Message")
}

func TestInvalidParameter(t *testing.T) {
	err := dummyFn()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Test Message")
	assert.True(t, cdserr.IsInvalidParameterError(err))

}
func TestWrapErrors(t *testing.T) {
	err := dummyFn()
	err = fmt.Errorf("haha: %w", err)

	assert.Equal(t, "haha: Test Message", err.Error())
	assert.Equal(t, "Test Message", errors.Unwrap(err).Error())
	assert.True(t, cdserr.IsInvalidParameterError(err))
}

func TestFatalError(t *testing.T) {
	underlyingError := errors.New("base")
	fatal := cdserr.NewFatalError(underlyingError, "thisisafatalerror")

	assert.Equal(t, "thisisafatalerror: base", fatal.Error())
	assert.True(t, cdserr.IsFatalError(fatal))
}

func TestWellKnownErrorsChecking(t *testing.T) {
	tests := []struct {
		err           error
		checkFunction func(err error) bool
	}{
		{cdserr.NewFatalError(nil, "error message"), cdserr.IsFatalError},
		{cdserr.NewAlreadyExistsError(nil, "error message"), cdserr.IsAlreadyExistsError},
		{cdserr.NewAuthenticationError(nil, "error message"), cdserr.IsAuthenticationError},
		{cdserr.NewNotExistsError(nil, "error message"), cdserr.IsNotExistsError},
		{cdserr.NewAuthenticationError(nil, "error message"), cdserr.IsAuthenticationError},
		{cdserr.NewAuthorizationError(nil, "error message"), cdserr.IsAuthorizationError},
		{cdserr.NewBusyError(nil, "error message"), cdserr.IsBusyError},
	}
	for _, test := range tests {
		t.Run(reflect.TypeOf(test.err).String(), func(t *testing.T) {
			assert.True(t, test.checkFunction(test.err))
		})
	}
}

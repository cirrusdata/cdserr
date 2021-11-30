package cdserr

import (
	"errors"
	"fmt"
	"net"
)

type CDSError interface {
	ErrorID() string
}

func wrapErrorIfNeeded(err error, msg string) error {
	if err == nil {
		if msg == "" {
			msg = "unknown error"
		}
		return fmt.Errorf(msg)
	}

	// if err is provided
	if msg == "" {
		return err
	} else {
		return fmt.Errorf("%s: %w", msg, err)
	}
}

type FatalError struct {
	error
}

func (e *FatalError) ErrorID() string {
	return "FATAL_ERROR"
}

func NewFatalError(err error, msg string) *FatalError {
	return &FatalError{error: wrapErrorIfNeeded(err, msg)}
}

func IsFatalError(err error) bool {
	var e *FatalError
	return errors.As(err, &e)
}

type InvalidParameterError struct {
	error
}

func (e *InvalidParameterError) ErrorID() string {
	return "INVALID_PARAMETER"
}

func NewInvalidParameterError(err error, msg string) *InvalidParameterError {
	return &InvalidParameterError{error: wrapErrorIfNeeded(err, msg)}
}

func IsInvalidParameterError(err error) bool {
	//if _, ok := err.(validation.Errors); ok {
	//	return true
	//}
	var e *InvalidParameterError
	return errors.As(err, &e)
}

type NotExistsError struct {
	error
}

func (e *NotExistsError) ErrorID() string {
	return "NOT_EXISTS"
}

func NewNotExistsError(err error, msg string) *NotExistsError {
	return &NotExistsError{error: wrapErrorIfNeeded(err, msg)}
}

func IsNotExistsError(err error) bool {
	var e *NotExistsError
	return errors.As(err, &e)
}

type AlreadyExistsError struct {
	error
}

func (e *AlreadyExistsError) ErrorID() string {
	return "ALREADY_EXISTS"
}

func NewAlreadyExistsError(err error, msg string) *AlreadyExistsError {
	return &AlreadyExistsError{error: wrapErrorIfNeeded(err, msg)}
}

func IsAlreadyExistsError(err error) bool {
	var e *AlreadyExistsError
	return errors.As(err, &e)
}

type InsufficientResourceError struct {
	error
}

func (e *InsufficientResourceError) ErrorID() string {
	return "INSUFFICIENT_ERROR"
}

func NewInsufficientResourceError(err error, msg string) *InsufficientResourceError {
	return &InsufficientResourceError{error: wrapErrorIfNeeded(err, msg)}
}

func IsInsufficientResourceError(err error) bool {
	var e *InsufficientResourceError
	return errors.As(err, &e)
}

type IllegalOperationError struct {
	error
}

func (e *IllegalOperationError) ErrorID() string {
	return "ILLEGAL_OPERATION"
}

func NewIllegalOperationError(err error, msg string) *IllegalOperationError {
	return &IllegalOperationError{error: wrapErrorIfNeeded(err, msg)}
}

func IsIllegalOperationError(err error) bool {
	var e *IllegalOperationError
	return errors.As(err, &e)
}

type BusyError struct {
	error
}

func (e *BusyError) ErrorID() string {
	return "BUSY"
}

func NewBusyError(err error, msg string) *BusyError {
	return &BusyError{error: wrapErrorIfNeeded(err, msg)}
}

func IsBusyError(err error) bool {
	var e *BusyError
	return errors.As(err, &e)
}

type AuthenticationError struct {
	error
}

func (e *AuthenticationError) ErrorID() string {
	return "AUTHENTICATION_ERROR"
}

func NewAuthenticationError(err error, msg string) *AuthenticationError {
	return &AuthenticationError{error: wrapErrorIfNeeded(err, msg)}
}

func IsAuthenticationError(err error) bool {
	var e *AuthenticationError
	return errors.As(err, &e)
}

type AuthorizationError struct {
	error
}

func (e *AuthorizationError) ErrorID() string {
	return "AUTHORIZATION_ERROR"
}

func NewAuthorizationError(err error, msg string) *AuthorizationError {
	return &AuthorizationError{error: wrapErrorIfNeeded(err, msg)}
}

func IsAuthorizationError(err error) bool {
	var e *AuthorizationError
	return errors.As(err, &e)
}

type TimeoutError struct {
	error
}

func (e *TimeoutError) ErrorID() string {
	return "TIMEOUT_ERROR"
}

func NewTimeoutError(err error, msg string) *TimeoutError {
	return &TimeoutError{error: wrapErrorIfNeeded(err, msg)}
}

func IsTimeoutError(err error) bool {
	// is system timeout error
	e, isNetError := err.(net.Error)
	if isNetError && e.Timeout() {
		return true
	}

	var er *TimeoutError
	return errors.As(err, &er)
}

type LicenseError struct {
	error
}

func (e *LicenseError) ErrorID() string {
	return "LICENSE_ERROR"
}

func NewLicenseError(err error, msg string) *LicenseError {
	return &LicenseError{error: wrapErrorIfNeeded(err, msg)}
}

func IsLicenseError(err error) bool {
	var e *LicenseError
	return errors.As(err, &e)
}

type NetworkError struct {
	error
}

func (e *NetworkError) ErrorID() string {
	return "NETWORK_ERROR"
}

func NewNetworkError(err error, msg string) *NetworkError {
	return &NetworkError{error: wrapErrorIfNeeded(err, msg)}
}

func IsNetworkError(err error) bool {
	var e *NetworkError
	return errors.As(err, &e)
}

type IOError struct {
	error
}

func (e *IOError) ErrorID() string {
	return "IO_ERROR"
}

func NewIOError(err error, msg string) *IOError {
	return &IOError{error: wrapErrorIfNeeded(err, msg)}
}

func IsIOError(err error) bool {
	var e *IOError
	return errors.As(err, &e)
}

type UnsupportedError struct {
	error
}

func NewUnsupportedError(err error, msg string) *UnsupportedError {
	return &UnsupportedError{error: wrapErrorIfNeeded(err, msg)}
}

func IsUnsupportedError(err error) bool {
	var e *UnsupportedError
	return errors.As(err, &e)
}

/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

	    All rights reserved
	EMail : <support@mwat.de>
*/
package codeerror

import (
	"fmt"
	"runtime"
)

//lint:file-ignore ST1017 - I prefer Yoda conditions

const (
	// The constant error message of the `ErrCodeLocation` error type.
	StrCodeLocation = "error in code"

	// How to build the string representation:
	strPattern = "File: %q, Line: %d, Function: %q"
)

type ErrCodeLocation struct {
	err      error
	File     string
	Function string
	Line     int
}

// `Error()` returns a string representation of the error message
// along with the error location.
//
// It includes the file name, line number, and function name where
// the error occurred, along with original error's text.
//
// Returns:
// - `string`: a string representation of the error message and location.
func (cl ErrCodeLocation) Error() string {
	return fmt.Sprintf("%s %s; %v", StrCodeLocation, cl.String(), cl.err)
} // Error()

// `String()` implements the `Stringer` interface and returns a string
// representation of the error location.
//
// It includes the file name, line number, and function name where
// the error occurred.
//
// Returns:
// - `string`: a string representation of the error location.
func (cl ErrCodeLocation) String() string {
	return fmt.Sprintf(strPattern, cl.File, cl.Line, cl.Function)
} // String()

// `Unwrap()` returns the original error that was wrapped by
// `ErrCodeLocation`.
//
// Returns:
// - `error`: the original error.
func (cl ErrCodeLocation) Unwrap() error {
	return cl.err
} // Unwrap()

// `CodeError()` is a function that wraps an error with additional
// information about the location where the error occurred. It uses
// certain `runtime` functions to determine the file- and function-names,
// as well as the code line. The `aLines` parameter allows for adjusting
// the reported line number by subtracting the specified number of lines
// from the actual line number.
//
// Parameters:
// - aErr: The error to be wrapped.
// - aLines: The number of lines to subtract from the caller's line number.
//
// Returns:
// - `error`: a `ErrCodeLocation` instance that contains the original error,
// file, function, and adjusted line number of the code causing `aErr`.
func CodeError(aErr error, aLines int) error {
	// Get program counter, file, line number, and status of the caller.
	pc, eFile, eLine, ok := runtime.Caller(1)
	if !ok {
		// it was not possible to recover the information
		return aErr
	}

	// Adjust the line number if `aLines` is greater than zero and
	// the calculated line number is not less than `aLines`.
	if 0 < aLines && eLine >= aLines {
		eLine -= aLines
	}

	// Get the name of the function for the program counter.
	eFunction := runtime.FuncForPC(pc).Name()

	// Return a new instance of `ErrCodeLocation` with the provided error,
	// file, function, and adjusted line number.
	return &ErrCodeLocation{
		err:      aErr,
		File:     eFile,
		Function: eFunction,
		Line:     eLine,
	}
} // CodeError()

/* _EoF_ */

/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

	    All rights reserved
	EMail : <support@mwat.de>
*/
package sourceerror

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

//lint:file-ignore ST1017 - I prefer Yoda conditions

const (
	// The constant error message of the `ErrSourceLocation` error type.
	StrCodeLocation = "error in source"

	// How to build the string representation:
	strPattern = "File: %q, Line: %d, Function: %q. Stack: %v"
)

type ErrSourceLocation struct {
	err      error
	File     string
	Function string
	Line     int
	Stack    string
}

var (
	// If set `true`, the `SourceError()` function basically becomes a NoOp.
	NODEBUG bool
)

// `Error()` returns a string representation of the error message
// along with the error location.
//
// It includes the file name, line number, and function name where
// the error occurred, along with original error's text.
//
// Returns:
// - `string`: a string representation of the error message and location.
func (se ErrSourceLocation) Error() string {
	return fmt.Sprintf("%s %s; %v",
		StrCodeLocation, se.String(), se.err)
} // Error()

// `String()` implements the `Stringer` interface and returns a string
// representation of the error location.
//
// It includes the file name, line number, and function name where
// the error occurred.
//
// Returns:
// - `string`: a string representation of the error location.
func (se ErrSourceLocation) String() string {
	return fmt.Sprintf(strPattern,
		se.File, se.Line, se.Function, se.Stack)
} // String()

// `Unwrap()` returns the original error that was wrapped by
// `ErrSourceLocation`.
//
// Returns:
// - `error`: the original error.
func (se ErrSourceLocation) Unwrap() error {
	return se.err
} // Unwrap()

// `SourceError()` is a function that wraps an error with additional
// information about the location where the error occurred. It uses
// certain `runtime` functions to determine the file- and function-names,
// as well as the code line and the call stack.
// The `aLines` parameter allows for adjusting the reported line number by
// subtracting the specified number of lines from the actual line number.
//
// NOTE: If the global `NODEBUG` flag is `true`, this function simply
// returns he given `aErr`, skipping the error location investigation.
//
// Parameters:
// - `aErr`: The error to be wrapped.
// - `aLines`: The number of lines to subtract from the caller's line number.
//
// Returns:
// - `error`: A new `ErrSourceLocation` instance that contains the original
// error, file, function, and adjusted line number of the code causing `aErr`.
func SourceError(aErr error, aLines int) error {
	if NODEBUG {
		return aErr
	}

	// Get program counter, file, line number, and status of the caller.
	pc, eFile, eLine, ok := runtime.Caller(1)
	if !ok {
		return aErr // not possible to recover the information
	}

	// Adjust the line number if `aLines` is greater than zero and
	// the calculated line number is not less than `aLines`.
	if 0 < aLines && eLine >= aLines {
		eLine -= aLines
	}

	// Get the name of the function for the program counter.
	eFunction := runtime.FuncForPC(pc).Name()

	// Get a formatted stack trace of the goroutine that calls it.
	eStack := string(debug.Stack())

	// Return a new instance of `ErrSourceLocation` with the provided error,
	// file, function, and adjusted line number.
	return &ErrSourceLocation{
		err:      aErr,
		File:     eFile,
		Function: eFunction,
		Line:     eLine,
		Stack:    eStack,
	}
} // SourceError()

/* _EoF_ */

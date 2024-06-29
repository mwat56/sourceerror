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
	StringSourceLocation = "error in source"

	// How to build the string representation:
	stringPattern = "Error: %v\nFile: %q\nLine: %d\nFunction: %q\nStack: %s"
)

// `ErrSource` is an error type that wraps another error with the
// location of where that other error was encountered.
// All public fields should be considered R/O (there really isn't any
// reason to modify those fields apart from confusing yourself).
//
// The fields are as follows:
// - `File`: The source file where the error was encountered.
// - `Function`: The function wherein the error was encountered
// - `Line`: The code line within the `File`.
// - `Stack`: The call stack to where the error was created.
type ErrSource struct {
	err      error  // 16 bytes
	File     string // 16 bytes
	Function string // dito
	Line     int    // 8 bytes
	Stack    []byte // 24 bytes
}

var (
	// If set `true`, the `Wrap()` function will skip the error
	// location investigation.
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
func (se ErrSource) Error() string {
	return fmt.Sprintf("%q\n%s", StringSourceLocation, se.primStr())
} // Error()

// `init()` is a special function in Go that is automatically called when
// a package is imported. It is used to initialise variables or perform
// any necessary setup for the package.
//
// Here, this function is used to ensure that the `ErrSourceLocation` type
// implements the `error` and `fmt.Stringer` interfaces.
// The `error` interface is required for any type that represents an error
// in Go, and the `fmt.Stringer` interface is used to provide a `String()`
// method for formatting the error message.
func init() {
	var (
		_ error        = ErrSource{}
		_ error        = (*ErrSource)(nil)
		_ fmt.Stringer = ErrSource{}
		_ fmt.Stringer = (*ErrSource)(nil)
	)
} // init()

// The `primStr()` method is an internal helper function that constructs a
// string representation of the error message along with the error location.
// The method's purpose is twofold: firstly it avoids implicit recursions
// between the `Error()` and `String()` methods, and secondly is serves
// as a helper for the unit-tests.
func (se ErrSource) primStr() string {
	return fmt.Sprintf(stringPattern,
		se.err, se.File, se.Line, se.Function, se.Stack)
} // primStr()

// `String()` implements the `Stringer` interface and returns a string
// representation of the error location.
//
// It includes the file name, line number, and function name where
// the error occurred ass well as a call stack.
//
// Returns:
// - `string`: a string representation of the error location.
func (se ErrSource) String() string {
	return se.primStr()
} // String()

// `Unwrap()` returns the original error that was wrapped by
// `ErrSourceLocation`.
//
// Returns:
// - `error`: the original error.
func (se ErrSource) Unwrap() error {
	return se.err
} // Unwrap()

// --------------------------------------------------------------------------

// `Wrap()` is a function that wraps an error with additional
// information about the location where the error occurred. It uses
// certain `runtime` functions to determine the file- and function-names,
// as well as the code line and the call stack.
// The `aLines` parameter allows for adjusting the reported line number by
// subtracting the specified number of lines from the actual line number.
//
// NOTE: If the global `NODEBUG` flag is `true`, this function returns an
// instance with the given `aErr`, while file, function, line number,
// stacktrace fields remain empty.
//
// Parameters:
// - `aErr`: The error to be wrapped.
// - `aLines`: The number of lines to subtract from the caller's line number.
//
// Returns:
// - `error`: A new `ErrSourceLocation` instance that contains `aErr`, as well
// as file, function, and adjusted line number of the code causing the error.
func Wrap(aErr error, aLines int) error {
	if NODEBUG {
		// Return a new instance of `ErrSource` with the provided error,
		// while file, function, line, and stack-trace remain empty.
		return &ErrSource{
			err: aErr,
		}
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

	// Return a new instance of `ErrSourceLocation` with the provided error,
	// file, function, adjusted line number, and stack trace.
	return &ErrSource{
		err:      aErr,
		File:     eFile,
		Function: eFunction,
		Line:     eLine,
		Stack:    debug.Stack(),
	}
} // Wrap()

/* _EoF_ */

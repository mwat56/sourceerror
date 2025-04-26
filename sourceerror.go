/*
Copyright © 2024, 2025  M.Watermann, 10247 Berlin, Germany

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

// --------------------------------------------------------------------------
// internally used constants:

const (
	// `StringErrSource` is the error message text of the
	// `ErrSource` error type.
	StringErrSource = "error in source"

	// `stringPattern` is the pattern used to build the string
	// representation.
	stringPattern = `Error: %v\nFile: "%s:%d"\nLine: %d\nFunction: %q\nStack: %s`
)

// --------------------------------------------------------------------------
// `ErrSource` type:

// `ErrSource` is an error type that wraps another error with the
// location of where that other error was encountered. It includes the
// file name, line number, and function name where the initial error
// occurred, along with original error's message text and call stack.
//
// All public fields should be considered R/O (there really isn't any
// reason to modify those fields apart from confusing yourself).
//
// The fields are as follows:
//   - `File`: The source file where the error was encountered.
//   - `Function`: The function wherein the error was encountered.
//   - `Line`: The code line within the `File`.
//   - `Stack`: The call stack to where the error was created.
type ErrSource struct {
	err      error  // 16 bytes
	File     string // 16 bytes
	Function string // dito
	Line     int    // 8 bytes
	Stack    []byte // 24 bytes
}

// --------------------------------------------------------------------------
// Public/exported variables:

var (
	// `NODEBUG` is a toggle used by [New] to either skip the error's
	// location investigation or include it.
	NODEBUG bool

	// `NOSTACK` is a toggle used by [New] to either skip the error's
	// call-stack investigation or include it.
	NOSTACK bool
)

// --------------------------------------------------------------------------

// `init()` is a special function in Go that is automatically called when
// a package is imported. It is used to initialise variables or perform
// any necessary setup for the package.
//
// Here, this function is used to ensure that the `ErrSource` type
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

// --------------------------------------------------------------------------
// `ErrSource` methods:

// `Error()` returns a string representation of the error message
// along with the error location.
//
// It includes the file name, line number, and function name where
// the error occurred, along with original error's message text and
// call stack.
//
// Returns:
//   - `string`: A string representation of the error message and location.
func (se ErrSource) Error() string {
	return fmt.Sprintf("%q\n%s", StringErrSource, se.primStr())
} // Error()

// The `primStr()` method is an internal helper function that constructs a
// string representation of the error message along with the error location.
// The method's purpose is twofold: firstly it avoids implicit recursions
// between the `Error()` and `String()` methods, and secondly is serves
// as a helper for the unit-tests.
//
// Returns:
//   - `string`: A string representation of the error instance.
func (se ErrSource) primStr() string {
	return fmt.Sprintf(stringPattern,
		se.err, se.File, se.Line, se.Line, se.Function, se.Stack)
} // primStr()

// `String()` implements the `Stringer` interface and returns a string
// representation of the error instance.
//
// It includes the file name, line number, and function name where
// the error occurred as well as a call stack.
//
// Returns:
//   - `string`: A string representation of the `ErrSource` instance.
func (se ErrSource) String() string {
	return se.primStr()
} // String()

// `Unwrap()` returns the original error that was wrapped by
// `ErrSource`.
//
// Returns:
//   - `error`: The original error.
func (se ErrSource) Unwrap() error {
	return se.err
} // Unwrap()

// --------------------------------------------------------------------------
// Public/exported function (i.e. constructor):

// `New()` returns a new `ErrSource` instance that wraps `aErr` with
// additional information about the location where the initial error
// occurred. It uses certain `runtime` functions to determine the file-
// and function-names, as well as the code line and the call stack.
//
// The `aLines` parameter allows for adjusting the reported line number by
// subtracting the specified number of lines from the actual line number
// to point to the code line where the initial error actually occurred.
//
// If the global `NODEBUG` flag is `true`, this function returns just
// the given `aErr`, without any memory overhead.
//
// If the global `NOSTACK` flag is `true`, this function returns does
// not add the initial error's call stack.
//
// Parameters:
//   - `aErr`: The error to be wrapped.
//   - `aLines`: The number of lines to subtract from the caller's line number.
//
// Returns:
//   - `error`: A new `ErrSource` instance that contains `aErr`, as well as
//     file, function, and adjusted line number of the code causing the error.
func New(aErr error, aLines int) error {
	if NODEBUG {
		// Return the provided error without any wrapping at all.
		return aErr
	}

	// Get program counter, file, line number, and status of the caller.
	pc, eFile, eLine, ok := runtime.Caller(1)
	if !ok {
		// not possible to recover the information
		return &ErrSource{
			err: aErr,
		}
	}

	// Adjust the line number if `aLines` is greater than zero and
	// the calculated line number is not less than `aLines`.
	if 0 < aLines && eLine >= aLines {
		eLine -= aLines
	}

	// Get the name of the function for the program counter.
	eFunction := runtime.FuncForPC(pc).Name()

	var eStack []byte
	if !NOSTACK {
		eStack = debug.Stack()
	}

	// Return a new instance of `ErrSource` with the provided error,
	// file, function, adjusted line number, and stack trace.
	return &ErrSource{
		err:      aErr,
		File:     eFile,
		Function: eFunction,
		Line:     eLine,
		Stack:    eStack,
	}
} // New()

// `Wrap()` wraps an error with additional information.
//
// Deprecated: `Wrap()` is deprecated and will be removed in the future. Call [New] instead.
func Wrap(aErr error, aLines int) error {
	return New(aErr, aLines)
} // Wrap()

/* _EoF_ */

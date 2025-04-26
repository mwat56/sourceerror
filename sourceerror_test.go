/*
Copyright © 2024, 2025  M.Watermann, 10247 Berlin, Germany

	    All rights reserved
	EMail : <support@mwat.de>
*/
package sourceerror

import (
	"errors"
	"fmt"
	"testing"
)

//lint:file-ignore ST1017 - I prefer Yoda conditions

func Test_ErrSource_Error(t *testing.T) {
	w0 := "some first error"
	e := errors.New(w0)
	cl1 := New(e, 1)
	w1 := fmt.Sprintf("%s", cl1) // use Stringer interface

	cl2 := New(nil, 0)
	w2 := fmt.Sprintf("%s", cl2)

	cl3 := New(cl1, 0)
	w3 := fmt.Sprintf("%s", cl3)

	cl4 := New(cl3, 0)
	w4 := fmt.Sprintf("%s", cl4)

	tests := []struct {
		name string
		err  error // i.e. ErrSource
		want string
	}{
		{"0", e, w0},
		{"1", cl1, w1},
		{"2", cl2, w2},
		{"3", cl3, w3},
		{"4", cl4, w4},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := tt.err
			if got := se.Error(); got != tt.want {
				t.Errorf("%q: ErrSource.Error() =\n%q,\nwant %q",
					tt.name, got, tt.want)
				return
			}
			if oe, ok := se.(*ErrSource); ok {
				t.Logf("%s\n", oe)
				return
			}

			t.Logf("%s\n", tt.want)
		})
	}
} // Test_ErrSource_Error()

func Test_ErrSource_String(t *testing.T) {
	var (
		w1, w2, w3, w4 string
		se             ErrSource
		ok             bool
	)

	e := errors.New("some first error")
	cl1 := New(e, 1)
	if se, ok = cl1.(ErrSource); ok {
		w1 = se.primStr()
	}
	cl2 := New(nil, 0)
	if se, ok = cl2.(ErrSource); ok {
		w2 = se.primStr()
	}
	cl3 := New(cl2, 4)
	if se, ok = cl3.(ErrSource); ok {
		w3 = se.primStr()
	}
	cl4 := New(cl3, 4)
	if se, ok = cl4.(ErrSource); ok {
		w4 = se.primStr()
	}

	tests := []struct {
		name string
		err  error // ErrSource
		want string
	}{
		{"1", cl1, w1},
		{"2", cl2, w2},
		{"3", cl3, w3},
		{"4", cl4, w4},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se, ok = tt.err.(ErrSource)
			if !ok {
				return
			}

			if got := se.String(); got != tt.want {
				t.Errorf("%q: ErrSource.String() =\n%s,\n>>>> want >>>>\n%s",
					tt.name, got, tt.want)
				return
			}

			t.Logf("\n%s\n", tt.want)
		})
	}
} // Test_ErrSource_String()

func Test_ErrSource_StringNODEBUG(t *testing.T) {
	NODEBUG = true
	defer func() {
		NODEBUG = false
	}()

	Test_ErrSource_String(t)
} // Test_ErrSource_StringNODEBUG()

func Test_ErrSource_StringNOSTACK(t *testing.T) {
	NOSTACK = true
	defer func() {
		NOSTACK = false
	}()

	Test_ErrSource_String(t)
} // Test_ErrSource_StringNOSTACK()

func Test_ErrSource_Unwrap(t *testing.T) {
	e1 := errors.New("some first error")
	cl1 := New(e1, 1)
	cl2 := New(nil, 0)
	cl3 := New(cl2, 1)
	cl4 := New(cl3, 1)

	tests := []struct {
		name    string
		err     error
		wantErr error
	}{
		{"1", cl1, e1},
		{"2", cl2, nil},
		{"3", cl3, cl2},
		{"4", cl4, cl3},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := errors.Unwrap(tt.err)
			if got != tt.wantErr {
				t.Errorf("%q: ErrSource.Unwrap() error =\n%v,\nwant:\n %v",
					tt.name, got, tt.wantErr)
				return
			}
			if e, ok := got.(*ErrSource); ok {
				t.Log("\t", e, "\n")
				return
			}

			t.Log("\t", tt.wantErr, "\n")
		})
	}
} // Test_ErrSource_Unwrap()

func Test_New(t *testing.T) {
	// Test cases
	tests := []struct {
		name      string
		err       error
		upLines   int
		wantNil   bool
		checkFile bool
	}{
		{"nil error", nil, 0, false, true},
		{"basic error", errors.New("test error"), 0, false, true},
		{"line adjustment", errors.New("line test"), 2, false, true},
		{"with NODEBUG", errors.New("nodebug test"), 4, false, false},
		{"with NOSTACK", errors.New("nostack test"), 4, false, false},
		{"nested error", New(errors.New("nested"), 0), 1, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save and restore NODEBUG state
			origNODEBUG := NODEBUG
			defer func() { NODEBUG = origNODEBUG }()

			// Set NODEBUG for specific test
			if tt.name == "with NODEBUG" {
				NODEBUG = true
			} else {
				NODEBUG = false
			}

			// Save and restore NOSTACK state
			origNOSTACK := NOSTACK
			defer func() { NOSTACK = origNOSTACK }()

			// Set NOSTACK for specific test
			if tt.name == "with NOSTACK" {
				NOSTACK = true
			} else {
				NOSTACK = false
			}

			// Call the function
			got := New(tt.err, tt.upLines)

			// Check if result is nil when expected
			if (nil == got) != tt.wantNil {
				t.Errorf("New() returned nil: %v, want nil: %v",
					nil == got, tt.wantNil)
				return
			}

			// For NODEBUG case, check that original error is returned
			if tt.name == "with NODEBUG" {
				if got != tt.err {
					t.Errorf("New() with NODEBUG = true didn't return original error")
				}
				return
			}

			// For NOSTACK case, the return error has an empty `Stack` field
			if tt.name == "with NOSTACK" {
				if 0 < len(got.(*ErrSource).Stack) {
					t.Errorf("New() with NOSTACK = true returned non-empty Stack")
				}
				return
			}

			// Type assertion
			se, ok := got.(*ErrSource)
			if !ok {
				t.Errorf("New() didn't return *ErrSource, got %T", got)
				return
			}

			// Check wrapped error
			if !errors.Is(got, tt.err) && (nil != tt.err) {
				t.Errorf("New() error chain doesn't contain original error")
			}

			// Check file and function info is populated
			if tt.checkFile {
				if "" == se.File {
					t.Errorf("New() returned empty File")
				}
				if "" == se.Function {
					t.Errorf("New() returned empty Function")
				}
				if 0 >= se.Line {
					t.Errorf("New() returned invalid Line: %d", se.Line)
				}
			}
		})
	}
} // Test_New()

/* _EoF_ */

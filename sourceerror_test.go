/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

	    All rights reserved
	EMail : <support@mwat.de>
*/
package sourceerror

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
)

//lint:file-ignore ST1017 - I prefer Yoda conditions

func TestErrSourceLocation_Error(t *testing.T) {
	runtime.GOMAXPROCS(1)

	cl1 := fmt.Errorf("dummy")
	e1 := fmt.Errorf("some first error")
	if nil != e1 {
		cl1 = SourceError(e1, 2)
	}
	w1 := fmt.Sprintf("%s", cl1) // use Stringer interface

	cl2 := SourceError(nil, 2)
	w2 := fmt.Sprintf("%s", cl2)

	tests := []struct {
		name string
		err  error // i.e. ErrSourceLocation
		want string
	}{
		{"1", cl1, w1},
		{"2", cl2, w2},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := tt.err
			if got := se.Error(); got != tt.want {
				t.Errorf("%q: ErrSourceLocation.Error() =\n%q,\nwant %q",
					tt.name, got, tt.want)
				return
			}

			t.Log(tt.want, "\n")
		})
	}
} // TestErrSourceLocation_Error()

func TestErrSourceLocation_String(t *testing.T) {
	runtime.GOMAXPROCS(1)

	var (
		cl1 ErrSourceLocation
	)
	e1 := fmt.Errorf("some first error")
	cl11 := SourceError(e1, 1).(*ErrSourceLocation)
	if errors.Is(cl11, cl1) {
		cl1 = *cl11
	}
	w1 := fmt.Sprintf(strPattern, cl1.File, cl1.Line, cl1.Function, cl1.Stack)

	cl2 := SourceError(nil, 0).(*ErrSourceLocation)
	w2 := fmt.Sprintf(strPattern, cl2.File, cl2.Line, cl2.Function, cl2.Stack)

	tests := []struct {
		name string
		err  ErrSourceLocation
		want string
	}{
		{"1", cl1, w1},
		{"2", *cl2, w2},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := tt.err
			if got := se.String(); got != tt.want {
				t.Errorf("%q: ErrSourceLocation.String() =\n%q,\nwant %q",
					tt.name, got, tt.want)
				return
			}

			t.Log(tt.want, "\n")
		})
	}
} // TestErrSourceLocation_String()

func TestErrSourceLocation_Unwrap(t *testing.T) {
	runtime.GOMAXPROCS(1)

	e1 := fmt.Errorf("some first error")
	cl1 := SourceError(e1, 1)
	cl2 := SourceError(nil, 0)

	tests := []struct {
		name    string
		err     error
		wantErr error
	}{
		{"1", cl1, e1},
		{"2", cl2, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := tt.err.(*ErrSourceLocation)
			if got := se.Unwrap(); got != tt.wantErr {
				t.Errorf("%q: ErrSourceLocation.Unwrap() error = %v, wantErr %v",
					tt.name, got, tt.wantErr)
				return
			}

			t.Log(tt.wantErr, "\n")
		})
	}
} // TestErrSourceLocation_Unwrap()

/* _EoF_ */

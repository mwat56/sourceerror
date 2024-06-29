/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

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

func TestErrSourceLocation_Error(t *testing.T) {
	w0 := "some first error"
	e := errors.New(w0)
	cl1 := Wrap(e, 1)
	w1 := fmt.Sprintf("%s", cl1) // use Stringer interface

	cl2 := Wrap(nil, 0)
	w2 := fmt.Sprintf("%s", cl2)

	cl3 := Wrap(cl1, 0)
	w3 := fmt.Sprintf("%s", cl3)

	cl4 := Wrap(cl3, 0)
	w4 := fmt.Sprintf("%s", cl4)

	tests := []struct {
		name string
		err  error // i.e. ErrSourceLocation
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
				t.Errorf("%q: ErrSourceLocation.Error() =\n%q,\nwant %q",
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
} // TestErrSourceLocation_Error()

func TestErrSourceLocation_String(t *testing.T) {
	e := errors.New("some first error")
	cl1 := Wrap(e, 1).(*ErrSource)
	w1 := cl1.primStr()

	cl2 := Wrap(nil, 0).(*ErrSource)
	w2 := cl2.primStr()

	cl3 := Wrap(cl2, 3).(*ErrSource)
	w3 := cl3.primStr()

	cl4 := Wrap(cl3, 3).(*ErrSource)
	w4 := cl4.primStr()

	tests := []struct {
		name string
		err  error // ErrSource
		want string
	}{
		{"1", *cl1, w1},
		{"2", *cl2, w2},
		{"3", *cl3, w3},
		{"4", *cl4, w4},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := tt.err.(ErrSource)
			if got := se.String(); got != tt.want {
				t.Errorf("%q: ErrSourceLocation.String() =\n%s,\n>>>> want >>>>\n%s",
					tt.name, got, tt.want)
				return
			}

			t.Logf("\n%s\n", tt.want)
		})
	}
} // TestErrSourceLocation_String()

func TestErrSourceLocation_Unwrap(t *testing.T) {
	e1 := errors.New("some first error")
	cl1 := Wrap(e1, 1)
	cl2 := Wrap(nil, 0)
	cl3 := Wrap(cl2, 1)
	cl4 := Wrap(cl3, 1)

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
			got := errors.Unwrap((tt.err))
			if got != tt.wantErr {
				t.Errorf("%q: ErrSourceLocation.Unwrap() error =\n%v,\nwant:\n %v",
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
} // TestErrSourceLocation_Unwrap()

/* _EoF_ */

/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

	    All rights reserved
	EMail : <support@mwat.de>
*/
package codeerror

import (
	"errors"
	"fmt"
	"testing"
)

//lint:file-ignore ST1017 - I prefer Yoda conditions

func TestErrCodeLocation_Error(t *testing.T) {
	cl1 := fmt.Errorf("dummy")
	e1 := fmt.Errorf("some first error")
	if nil != e1 {
		cl1 = CodeError(e1, 2)
	}
	w1 := fmt.Sprintf("%s", cl1) // use Stringer interface

	cl2 := CodeError(nil, 2)
	w2 := fmt.Sprintf("%s", cl2)

	tests := []struct {
		name   string
		fields error // i.e. ErrCodeLocation
		want   string
	}{
		{"1", cl1, w1},
		{"2", cl2, w2},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := tt.fields
			if got := cl.Error(); got != tt.want {
				t.Errorf("%q: ErrCodeLocation.Error() =\n%q,\nwant %q",
					tt.name, got, tt.want)
			}

			t.Log(tt.want, "\n")
		})
	}
} // testErrCodeLocation()

func TestErrCodeLocation_String(t *testing.T) {
	var (
		cl1 ErrCodeLocation
	)
	e1 := fmt.Errorf("some first error")
	cl11 := CodeError(e1, 1).(*ErrCodeLocation)
	if errors.Is(cl11, cl1) {
		cl1 = *cl11
	}
	w1 := fmt.Sprintf(strPattern, cl1.File, cl1.Line, cl1.Function)

	cl2 := CodeError(nil, 0).(*ErrCodeLocation)
	w2 := fmt.Sprintf(strPattern, cl2.File, cl2.Line, cl2.Function)

	tests := []struct {
		name   string
		fields ErrCodeLocation
		want   string
	}{
		{"1", cl1, w1},
		{"2", *cl2, w2},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := tt.fields
			if got := cl.String(); got != tt.want {
				t.Errorf("%q: ErrCodeLocation.String() =\n%q,\nwant %q",
					tt.name, got, tt.want)
			}

			t.Log(tt.want, "\n")
		})
	}
} // TestErrCodeLocation_String()

func TestErrCodeLocation_Unwrap(t *testing.T) {
	e1 := fmt.Errorf("some first error")
	cl1 := CodeError(e1, 1)
	cl2 := CodeError(nil, 0)

	tests := []struct {
		name    string
		fields  error
		wantErr error
	}{
		{"1", cl1, e1},
		{"2", cl2, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := tt.fields.(*ErrCodeLocation)
			if got := cl.Unwrap(); got != tt.wantErr {
				t.Errorf("%q: ErrCodeLocation.Unwrap() error = %v, wantErr %v",
					tt.name, got, tt.wantErr)
			}

			t.Log(tt.wantErr, "\n")
		})
	}
} // TestErrCodeLocation_Unwrap()

/* _EoF_ */

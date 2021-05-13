package mfs

import (
	"fmt"

	"github.com/Orest-2/imfk/lib"
)

// Trapezoidal ...
type Trapezoidal struct {
	a, b, c, d float64
}

// NewTrapezoidal ...
func NewTrapezoidal(params []float64) (*Trapezoidal, error) {

	t := &Trapezoidal{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t Trapezoidal) GetParams() []float64 {

	return []float64{t.a, t.b, t.c, t.d}

}

// SetParams ...
func (t Trapezoidal) SetParams(params []float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t Trapezoidal) Eval(x float64) float64 {

	a, b, c, d := t.a, t.b, t.c, t.d

	if a <= x && x < b {
		f := x - a
		s := b - a

		return f / s
	}

	if b <= x && x < c {
		return 1
	}

	if c <= x && x < d {
		f := d - x
		s := d - c

		return f / s
	}

	return 0

}

// validateAndSetParams ...
func (t *Trapezoidal) validateAndSetParams(params []float64) error {

	var a, b, c, d float64

	mp := 4

	if len(params) < mp {
		return fmt.Errorf("мінімум %v параметра: (a, b, c, d)", mp)
	}

	lib.UnpackFloat64(params, &a, &b, &c, &d)

	if !(a < b && b < c && c < d) {
		return fmt.Errorf(
			"має виконуватися співвідношенням a < b < c < d: (a=%v, b=%v, c=%v, d=%v)",
			a, b, c, d,
		)
	}

	t.a, t.b, t.c, t.d = a, b, c, d

	return nil

}

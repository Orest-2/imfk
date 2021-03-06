package mfs

import (
	"fmt"
	"math"

	"github.com/Orest-2/imfk/lib"
)

// SquareSSpline ...
type SquareSSpline struct {
	a, b float64
}

// NewSquareSSpline ...
func NewSquareSSpline(params []float64) (*SquareSSpline, error) {

	t := &SquareSSpline{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t SquareSSpline) GetParams() []float64 {

	return []float64{t.a, t.b}

}

// SetParams ...
func (t SquareSSpline) SetParams(params []float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t SquareSSpline) Eval(x float64) float64 {

	a, b, sumab := t.a, t.b, t.a+t.b

	if x <= a {
		return 0
	}

	if a < x && x <= sumab/2 {
		f := x - a
		s := b - a

		return 2 * math.Pow((f/s), 2)
	}

	if sumab/2 < x && x < b {
		f := b - x
		s := b - a

		return 1 - 2*math.Pow((f/s), 2)
	}

	return 1

}

// validateAndSetParams ...
func (t *SquareSSpline) validateAndSetParams(params []float64) error {

	var a, b float64

	mp := 2

	if len(params) < mp {
		return fmt.Errorf("мінімум %v параметра: (a, b)", mp)
	}

	lib.UnpackFloat64(params, &a, &b)

	if !(a < b) {
		return fmt.Errorf("має виконуватися співвідношенням a < b: (a=%v, b=%v)", a, b)
	}

	t.a, t.b = a, b

	return nil

}

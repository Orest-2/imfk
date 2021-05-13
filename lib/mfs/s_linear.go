package mfs

import (
	"fmt"

	"github.com/Orest-2/imfk/lib"
)

// SLinear ...
type SLinear struct {
	a, b float64
}

// NewSLinear ...
func NewSLinear(params []float64) (*SLinear, error) {

	t := &SLinear{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t SLinear) GetParams() []float64 {

	return []float64{t.a, t.b}

}

// SetParams ...
func (t SLinear) SetParams(params []float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t SLinear) Eval(x float64) float64 {

	a, b := t.a, t.b

	if x <= a {
		return 0
	}

	if a < x && x <= b {
		f := x - a
		s := b - a

		return f / s
	}

	return 1

}

// validateAndSetParams ...
func (t *SLinear) validateAndSetParams(params []float64) error {

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

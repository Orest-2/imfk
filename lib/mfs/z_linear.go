package mfs

import (
	"fmt"

	"github.com/Orest-2/imfk/lib"
)

// ZLinear ...
type ZLinear struct {
	a, b float64
}

// NewZLinear ...
func NewZLinear(params []float64) (*ZLinear, error) {

	t := &ZLinear{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t ZLinear) GetParams() []float64 {

	return []float64{t.a, t.b}

}

// SetParams ...
func (t ZLinear) SetParams(params []float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t ZLinear) Eval(x float64) float64 {

	a, b := t.a, t.b

	if x <= a {
		return 1
	}

	if a < x && x <= b {
		f := b - x
		s := b - a

		return f / s
	}

	return 0

}

// validateAndSetParams ...
func (t *ZLinear) validateAndSetParams(params []float64) error {

	var a, b float64

	mp := 2

	if len(params) < mp {
		return fmt.Errorf("Мінімум %v параметра: (a, b)", mp)
	}

	lib.UnpackFloat64(params, &a, &b)

	if !(a < b) {
		return fmt.Errorf("має виконуватися співвідношенням a < b: (a=%v, b=%v)", a, b)
	}

	t.a, t.b = a, b

	return nil

}

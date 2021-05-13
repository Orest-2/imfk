package mfs

import (
	"fmt"
	"math"

	"github.com/Orest-2/imfk/lib"
)

// Gauss ...
type Gauss struct {
	a, b float64
}

// NewGauss ...
func NewGauss(params []float64) (*Gauss, error) {

	t := &Gauss{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t Gauss) GetParams() []float64 {

	return []float64{t.a, t.b}

}

// SetParams ...
func (t Gauss) SetParams(params []float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t Gauss) Eval(x float64) float64 {

	a, b := t.a, t.b

	f := math.Pow(x-b, 2)
	s := 2 * math.Pow(a, 2)

	return math.Exp(-f / s)

}

// validateAndSetParams ...
func (t *Gauss) validateAndSetParams(params []float64) error {

	var a, b float64

	mp := 2

	if len(params) < mp {
		return fmt.Errorf("мінімум %v параметра: (a, b)", mp)
	}

	lib.UnpackFloat64(params, &a, &b)

	if a <= 0 {
		return fmt.Errorf("має виконуватися співвідношенням a > 0: (a=%v)", a)
	}

	if !(a < b) {
		return fmt.Errorf("має виконуватися співвідношенням a < b: (a=%v, b=%v)", a, b)
	}

	t.a, t.b = a, b

	return nil

}

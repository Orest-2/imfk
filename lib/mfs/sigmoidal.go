package mfs

import (
	"fmt"
	"math"

	"github.com/Orest-2/imfk/lib"
)

// Sigmoidal ...
type Sigmoidal struct {
	a, b float64
}

// NewSigmoidal ...
func NewSigmoidal(params []float64) (*Sigmoidal, error) {

	t := &Sigmoidal{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t Sigmoidal) GetParams() []float64 {

	return []float64{t.a, t.b}

}

// SetParams ...
func (t Sigmoidal) SetParams(params []float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t Sigmoidal) Eval(x float64) float64 {

	a, b := t.a, t.b

	return 1 / (1 + math.Exp(-a*(x-b)))

}

// validateAndSetParams ...
func (t *Sigmoidal) validateAndSetParams(params []float64) error {

	var a, b float64

	mp := 2

	if len(params) < mp {
		return fmt.Errorf("мінімум %v параметра: (a, b)", mp)
	}

	lib.UnpackFloat64(params, &a, &b)

	// if !(a < b) {
	// 	return fmt.Errorf("має виконуватися співвідношенням a < b: (a=%v, b=%v)", a, b)
	// }

	t.a, t.b = a, b

	return nil

}

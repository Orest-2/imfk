package mfs

import (
	"fmt"
	"math"

	"github.com/Orest-2/imfk/lib"
)

// GeneralizedSigmoid ...
type GeneralizedSigmoid struct {
	x0, h []float64
}

// NewGeneralizedSigmoid ...
func NewGeneralizedSigmoid(params [][]float64) (*GeneralizedSigmoid, error) {

	t := &GeneralizedSigmoid{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t GeneralizedSigmoid) GetParams() [][]float64 {

	return [][]float64{t.x0, t.h}

}

// SetParams ...
func (t GeneralizedSigmoid) SetParams(params [][]float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t GeneralizedSigmoid) Eval(xv []float64) float64 {

	x0, h := t.x0, t.h

	e := 0.0

	for i, x := range xv {

		e -= (x0[i] * -1) * (x - h[i])

	}

	r := 1 / (1 + math.Exp(e))

	return r

}

// validateAndSetParams ...
func (t *GeneralizedSigmoid) validateAndSetParams(params [][]float64) error {

	var a, b []float64

	mp := 2

	if len(params) < mp {
		return fmt.Errorf("мінімум %v параметра: (a, b)", mp)
	}

	lib.UnpackFloat64Slice(params, &a, &b)

	if len(a) < 2 {
		return fmt.Errorf("параметер a має містит два елемента: (a=%v)", a)
	}

	if len(b) < 2 {
		return fmt.Errorf("параметер b має містит два елемента: (b=%v)", b)
	}

	t.x0, t.h = b, a

	return nil

}

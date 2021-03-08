package mfs

import (
	"fmt"
	"math"

	"github.com/Orest-2/imfk/lib"
)

// Pyramidal ...
type Pyramidal struct {
	x0, h []float64
}

// NewPyramidal ...
func NewPyramidal(params [][]float64) (*Pyramidal, error) {

	t := &Pyramidal{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t Pyramidal) GetParams() [][]float64 {

	return [][]float64{t.x0, t.h}

}

// SetParams ...
func (t Pyramidal) SetParams(params [][]float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t Pyramidal) Eval(xv []float64) float64 {

	x0, h := t.x0, t.h

	r := 0.0

	for i, x := range xv {

		f := math.Abs(x - x0[i])
		s := h[i]

		r += f / s

	}

	return math.Max(1-r, 0)

}

// validateAndSetParams ...
func (t *Pyramidal) validateAndSetParams(params [][]float64) error {

	var a, b []float64

	mp := 2

	if len(params) < mp {
		return fmt.Errorf("Мінімум %v параметра: (a, b)", mp)
	}

	lib.UnpackFloat64Slice(params, &a, &b)

	if len(a) < 2 {
		return fmt.Errorf("параметер a має містит два елемента: (a=%v)", a)
	}

	if len(b) < 2 {
		return fmt.Errorf("параметер b має містит два елемента: (b=%v)", b)
	}

	t.x0, t.h = a, b

	return nil

}

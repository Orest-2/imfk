package mfs

import (
	"fmt"
	"math"

	"github.com/Orest-2/imfk/lib"
)

// TrapezoidalPyramidal ...
type TrapezoidalPyramidal struct {
	x0, h, a []float64
}

// NewTrapezoidalPyramidal ...
func NewTrapezoidalPyramidal(params [][]float64) (*TrapezoidalPyramidal, error) {

	t := &TrapezoidalPyramidal{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t TrapezoidalPyramidal) GetParams() [][]float64 {

	return [][]float64{t.x0, t.h, t.a}

}

// SetParams ...
func (t TrapezoidalPyramidal) SetParams(params [][]float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t TrapezoidalPyramidal) Eval(xv []float64) float64 {

	x0, h, a := t.x0, t.h[0]+t.h[1], t.a

	r := 0.0

	for i, x := range xv {

		f := math.Abs(x - x0[i])
		s := a[i]

		r += (f / s)
	}

	if h-r >= 0 && h-r <= 1 {
		return h - r
	}

	if h-r > 1 {
		return 1
	}

	return 0

}

// validateAndSetParams ...
func (t *TrapezoidalPyramidal) validateAndSetParams(params [][]float64) error {

	var a, b, c []float64

	mp := 3

	if len(params) < mp {
		return fmt.Errorf("Мінімум %v параметра: (a, b)", mp)
	}

	lib.UnpackFloat64Slice(params, &a, &b, &c)

	if len(a) < 2 {
		return fmt.Errorf("параметер a має містит два елемента: (a=%v)", a)
	}

	if len(b) < 2 {
		return fmt.Errorf("параметер b має містит два елемента: (b=%v)", b)
	}

	if len(c) < 2 {
		return fmt.Errorf("параметер b має містит два елемента: (c=%v)", c)
	}

	t.x0, t.h, t.a = a, b, c

	return nil

}

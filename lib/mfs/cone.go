package mfs

import (
	"fmt"
	"math"

	"github.com/Orest-2/imfk/lib"
)

// Cone ...
type Cone struct {
	x0, h []float64
}

// NewCone ...
func NewCone(params [][]float64) (*Cone, error) {

	t := &Cone{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t Cone) GetParams() [][]float64 {

	return [][]float64{t.x0, t.h}

}

// SetParams ...
func (t Cone) SetParams(params [][]float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t Cone) Eval(xv []float64) float64 {

	x0, h := t.x0, t.h

	e := 0.0

	for i, x := range xv {

		f := math.Pow(x-x0[i], 2)
		s := math.Pow(h[i], 2)

		e += f / s

	}

	r := math.Sqrt(e)

	if r <= 1 {
		return 1 - r
	}

	return 0.00

}

// validateAndSetParams ...
func (t *Cone) validateAndSetParams(params [][]float64) error {

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

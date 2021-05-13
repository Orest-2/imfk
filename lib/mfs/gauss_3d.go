package mfs

import (
	"fmt"
	"math"

	"github.com/Orest-2/imfk/lib"
)

// Gauss3D ...
type Gauss3D struct {
	a, b []float64
}

// NewGauss3D ...
func NewGauss3D(params [][]float64) (*Gauss3D, error) {

	t := &Gauss3D{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t Gauss3D) GetParams() [][]float64 {

	return [][]float64{t.a, t.b}

}

// SetParams ...
func (t Gauss3D) SetParams(params [][]float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t Gauss3D) Eval(xv []float64) float64 {

	a, b := t.a, t.b

	e := 0.0

	for i, x := range xv {

		e1 := x - a[i]

		e += math.Pow(e1/b[i], 2)

	}

	return math.Exp(-e)

}

// validateAndSetParams ...
func (t *Gauss3D) validateAndSetParams(params [][]float64) error {

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

	t.a, t.b = a, b

	return nil

}

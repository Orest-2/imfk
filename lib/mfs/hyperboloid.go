package mfs

import (
	"fmt"
	"math"

	"github.com/Orest-2/imfk/lib"
)

// Hyperboloid ...
type Hyperboloid struct {
	a, b []float64
}

// NewHyperboloid ...
func NewHyperboloid(params [][]float64) (*Hyperboloid, error) {

	t := &Hyperboloid{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t Hyperboloid) GetParams() [][]float64 {

	return [][]float64{t.a, t.b}

}

// SetParams ...
func (t Hyperboloid) SetParams(params [][]float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t Hyperboloid) Eval(xv []float64) float64 {

	a, b := t.a, t.b

	e := 1.0

	for i, x := range xv {

		e1 := math.Pow(x-a[i], 2)
		e2 := math.Pow(b[i], 2)

		e += e1 / e2

	}

	r := math.Sqrt(e)

	if r < 2 {
		return 2 - r
	}

	return 0

}

// validateAndSetParams ...
func (t *Hyperboloid) validateAndSetParams(params [][]float64) error {

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

	t.a, t.b = a, b

	return nil

}

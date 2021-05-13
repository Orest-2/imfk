package mfs

import (
	"fmt"

	"github.com/Orest-2/imfk/lib"
)

// Triangular ...
type Triangular struct {
	a, b, c float64
}

// NewTriangular ...
func NewTriangular(params []float64) (*Triangular, error) {

	t := &Triangular{}

	if err := t.validateAndSetParams(params); err != nil {
		return nil, err
	}

	return t, nil

}

// GetParams ...
func (t Triangular) GetParams() []float64 {

	return []float64{t.a, t.b, t.c}

}

// SetParams ...
func (t Triangular) SetParams(params []float64) error {

	if err := t.validateAndSetParams(params); err != nil {
		return err
	}

	return nil

}

// Eval ...
func (t Triangular) Eval(x float64) float64 {

	a, b, c := t.a, t.b, t.c

	if a < x && x <= b {
		f := x - a
		s := b - a

		return f / s
	}

	if b < x && x < c {
		f := c - x
		s := c - b

		return f / s
	}

	return 0

}

// validateAndSetParams ...
func (t *Triangular) validateAndSetParams(params []float64) error {

	var a, b, c float64

	mp := 3

	if len(params) < mp {
		return fmt.Errorf("мінімум %v параметра: (a, b, c)", mp)
	}

	lib.UnpackFloat64(params, &a, &b, &c)

	if !(a < b && b < c) {
		return fmt.Errorf("має виконуватися співвідношенням a < b < c: (a=%v, b=%v, c=%v)", a, b, c)
	}

	t.a, t.b, t.c = a, b, c

	return nil

}

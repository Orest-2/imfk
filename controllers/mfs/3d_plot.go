package mfsController

import (
	"errors"
	"net/http"

	"github.com/Orest-2/imfk/lib"
	"github.com/Orest-2/imfk/lib/mfs"
	"github.com/Orest-2/imfk/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type make3DPlotRequest struct {
	MembershipFuncID int         `json:"membership_func_id" binding:"required"`
	FuncParams       [][]float64 `json:"func_params" binding:"required"`
	PlotParams       []float64   `json:"plot_params"`
}

type make3DPlotResponse struct {
	make3DPlotRequest
	X []float64   `json:"x"`
	Y []float64   `json:"y"`
	Z [][]float64 `json:"z"`
}

// Make3DPlot ...
func Make3DPlot(c *gin.Context) {
	json := make3DPlotRequest{}
	mf := models.MembershipFunc{}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.DB.First(&mf, json.MembershipFuncID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := make3DPlotResponse{
		make3DPlotRequest: make3DPlotRequest{
			MembershipFuncID: json.MembershipFuncID,
			PlotParams:       json.PlotParams,
		},
	}

	res.X, res.Y, _, res.Z, res.make3DPlotRequest.FuncParams, err = get3DPlotData(mf, json.FuncParams, nil, json.PlotParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": res,
		},
	)

}

func get3DPlotData(
	mf models.MembershipFunc, funcParams, inData [][]float64, plotParams []float64,
) (
	[]float64, []float64, []float64, [][]float64, [][]float64, error,
) {
	vx := []float64{}
	vy := []float64{}
	vz := []float64{}
	mz := [][]float64{}
	mfParams := [][]float64{}

	p := 0.1
	m := 10.0
	s := 1

	if len(plotParams) >= 2 {
		lib.UnpackFloat64(plotParams, &m, &p)

		if p <= 0 {
			p = m / 100
		}

		s = int(m / p)

		for i := 0; i <= s; i++ {

			k := float64(i)
			vx = append(vx, k*p)
			vy = append(vy, k*p)

		}
	}

	getInData := func() [][]float64 {

		if inData != nil {
			return inData
		}

		res := [][]float64{}

		x := 0
		y := 0

		for i := 0; i <= s*s; i++ {
			x = i / s

			if i%s == 0 {
				y = 0
			}

			res = append(res, []float64{float64(y) * p, float64(x) * p})

			y++
		}

		return res

	}

	switch mf.Code {
	case mfs.ConeMembershipFuncCode:

		tmf, err := mfs.NewCone(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	case mfs.PyramidalMembershipFuncCode:

		tmf, err := mfs.NewPyramidal(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	case mfs.TrapezoidalPyramidalMembershipFuncCode:

		tmf, err := mfs.NewTrapezoidalPyramidal(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	case mfs.GeneralizedSigmoidMembershipFuncCode:

		tmf, err := mfs.NewGeneralizedSigmoid(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	case mfs.Bell3DShapedMembershipFuncCode:

		tmf, err := mfs.NewBellShaped3D(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	case mfs.Gauss3DMembershipFuncCode:

		tmf, err := mfs.NewGauss3D(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	case mfs.HyperboloidMembershipFuncCode:

		tmf, err := mfs.NewHyperboloid(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	case mfs.Ellipsoid3DMembershipFuncCode:

		tmf, err := mfs.NewEllipsoid(funcParams)
		if err != nil {
			return vx, vy, vz, mz, mfParams, err
		}

		mfParams = tmf.GetParams()

		vxy := getInData()

		for _, xy := range vxy {
			z := tmf.Eval(xy)
			vz = append(vz, z)
		}

		mz = lib.Vector2Matrix(vz, s)

	}

	return vx, vy, vz, mz, mfParams, nil
}

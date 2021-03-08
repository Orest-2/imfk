package controllers

import (
	"errors"
	"math"
	"net/http"

	"github.com/Orest-2/imfk/lib/mfs"
	"github.com/Orest-2/imfk/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type make2DPlotRequest struct {
	MembershipFuncID int       `json:"membership_func_id" binding:"required"`
	FuncParams       []float64 `json:"func_params" binding:"required"`
}

type make2DPlotResponse struct {
	make2DPlotRequest
	X []float64 `json:"x"`
	Y []float64 `json:"y"`
}

// Make2DPlot ...
func Make2DPlot(c *gin.Context) {
	json := make2DPlotRequest{}
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

	res := make2DPlotResponse{
		make2DPlotRequest: make2DPlotRequest{
			MembershipFuncID: json.MembershipFuncID,
		},
	}

	res.X, res.Y, res.make2DPlotRequest.FuncParams, err = get2DPlotData(mf, json.FuncParams, nil)
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

func get2DPlotData(
	mf models.MembershipFunc, params []float64, inData []float64,
) (
	[]float64, []float64, []float64, error,
) {

	vx := []float64{}
	vy := []float64{}
	mfParams := []float64{}

	getVectorX := func(p []float64, i int) []float64 {

		if inData != nil {

			return inData

		}

		x := []float64{}

		fp := p[i]
		fpabs := math.Abs(fp)
		lp := p[len(p)-1]

		fromX := fp - fpabs
		toX := lp + fpabs
		for i := fromX; i <= toX; i++ {
			x = append(x, i)
		}

		return x

	}

	switch mf.Code {
	case mfs.TriangularMembershipFuncCode:

		tmf, err := mfs.NewTriangular(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.TrapezoidalMembershipFuncCode:

		tmf, err := mfs.NewTrapezoidal(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.SquareZSplineMembershipFuncCode:

		tmf, err := mfs.NewSquareZSpline(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.HarmoniousZSplineMembershipFuncCode:

		tmf, err := mfs.NewHarmoniousZSpline(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case
		mfs.SSigmoidalMembershipFuncCode,
		mfs.ZSigmoidalMembershipFuncCode:

		tmf, err := mfs.NewSigmoidal(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 1)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.ZLinearMembershipFuncCode:

		tmf, err := mfs.NewZLinear(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.SquareSSplineMembershipFuncCode:

		tmf, err := mfs.NewHarmoniousSSpline(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.HarmoniousSSplineMembershipFuncCode:

		tmf, err := mfs.NewHarmoniousSSpline(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.SLinearMembershipFuncCode:

		tmf, err := mfs.NewSLinear(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 0)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.BellShapedMembershipFuncCode:

		tmf, err := mfs.NewBellShaped(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 2)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	case mfs.GaussMembershipFuncCode:

		tmf, err := mfs.NewGauss(params)
		if err != nil {
			return vx, vy, mfParams, err
		}

		mfParams = tmf.GetParams()

		vx = getVectorX(mfParams, 1)

		for _, x := range vx {
			y := tmf.Eval(x)
			vy = append(vy, y)
		}

	}

	return vx, vy, mfParams, nil

}

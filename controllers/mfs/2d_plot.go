package mfsController

import (
	"errors"
	"net/http"

	"github.com/Orest-2/imfk/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type make2DPlotRequest struct {
	MembershipFuncID int       `json:"membership_func_id" binding:"required"`
	FuncParams       []float64 `json:"func_params" binding:"required"`
	PlotParams       []float64 `json:"plot_params"`
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
			PlotParams:       json.PlotParams,
		},
	}

	res.X, res.Y, res.make2DPlotRequest.FuncParams, err = get2DPlotData(mf, json.FuncParams, nil, json.PlotParams)
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

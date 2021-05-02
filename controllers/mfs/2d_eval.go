package mfsController

import (
	"errors"
	"net/http"

	"github.com/Orest-2/imfk/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type eval2DRequest struct {
	MembershipFuncID int       `json:"membership_func_id" binding:"required"`
	FuncParams       []float64 `json:"func_params" binding:"required"`
	InData           []float64 `json:"in_data" binding:"required"`
}

type eval2DResponse struct {
	eval2DRequest
	InData []float64 `json:"in_data"`
	Result []float64 `json:"result"`
}

// Eval2D ...
func Eval2D(c *gin.Context) {
	json := eval2DRequest{}
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

	res := eval2DResponse{
		eval2DRequest: eval2DRequest{
			MembershipFuncID: json.MembershipFuncID,
		},
	}

	res.InData, res.Result, res.eval2DRequest.FuncParams, err = get2DPlotData(mf, json.FuncParams, json.InData, nil)
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

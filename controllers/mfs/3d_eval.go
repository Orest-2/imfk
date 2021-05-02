package mfsController

import (
	"errors"
	"net/http"

	"github.com/Orest-2/imfk/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type eval3DRequest struct {
	MembershipFuncID int         `json:"membership_func_id" binding:"required"`
	FuncParams       [][]float64 `json:"func_params" binding:"required"`
	InData           [][]float64 `json:"in_data" binding:"required"`
}

type eval3DResponse struct {
	eval3DRequest
	InData [][]float64 `json:"in_data"`
	Result []float64   `json:"result"`
}

// Eval3D ...
func Eval3D(c *gin.Context) {
	json := eval3DRequest{}
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

	res := eval3DResponse{
		eval3DRequest: eval3DRequest{
			MembershipFuncID: json.MembershipFuncID,
		},
		InData: json.InData,
	}

	_, _, res.Result, _, res.eval3DRequest.FuncParams, err = get3DPlotData(mf, json.FuncParams, json.InData, nil)
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

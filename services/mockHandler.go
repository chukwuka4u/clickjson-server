package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// represents a single data property in the requirest schema
type Field struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

// represents the overall schema definition payload
type MockRequest struct {
	Count  int     `json:"count"`
	Fields []Field `json:"fields" binding:"required,gt=0"`
}

func HandleMockGeneration(ctx *gin.Context) {
	var req MockRequest

	// validates incoming json body against our MockRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// sets a count limit if the developer did not set one
	if req.Count <= 0 {
		req.Count = 10
	}
	if req.Count > 100 {
		req.Count = 100
	}

	mockDataList := make([]map[string]interface{}, req.Count)

	//traverse through the count as a loop to build each row
	for i := 0; i < req.Count; i++ {
		row := make(map[string]interface{})
		for _, field := range req.Fields {
			row[field.Name] = generateValue(field.Type)
		}
		mockDataList[i] = row
	}

	ctx.JSON(http.StatusOK, mockDataList)
}

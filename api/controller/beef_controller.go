package controller

import (
	"7solution/api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BeefController struct {
	WordCounter *service.WordCounter
	Baconipsum  service.BaconipsumGetter
}

// SummaryHandler fetch data from baconipsum and extract meat out of the content
func (ctrl *BeefController) SummaryHandler(c *gin.Context) {
	// to query meat-and-filler or all-meat based on ?type= query string
	// default to meat-and-filler
	beefType := "meat-and-filler"
	if c.Query("type") == "all-meat" {
		beefType = "all-meat"
	}

	// retrieve content from baconipsum
	text, err := ctrl.Baconipsum.Get(beefType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error retrieving data from baconipsum.com api",
		})
		return
	}

	// processing text
	summary := ctrl.WordCounter.CountAllWord(text)

	c.JSON(http.StatusOK, gin.H{
		"beef": summary,
	})
}

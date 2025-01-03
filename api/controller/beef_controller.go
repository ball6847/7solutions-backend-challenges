package controller

import (
	"7solution/api/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	// excludes, after extraction we still got some irrelevant words
	// let's define and exclude some here
	// for baconipsum default words see - https://github.com/petenelson/wp-any-ipsum/blob/trunk/lib/default-custom.txt
	excludes = []string{
		"kevin", // Why does kevin end up here in the bacon list ? Kevin Bacon maybe ?
		"tail",  // eatable but, does this consider meat ?
		"turkey",
		"chicken",
		"cow",
		"pig",
	}

	// fillers - see https://github.com/petenelson/wp-any-ipsum/blob/trunk/lib/default-filler.txt
	fillers = []string{
		"lorem", "ipsum", "consectetur", "adipisicing", "elit", "sed",
		"do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et",
		"dolore", "magna", "aliqua", "ut", "enim", "ad", "minim", "veniam",
		"quis", "nostrud", "exercitation", "ullamco", "laboris", "nisi", "ut",
		"aliquip", "ex", "ea", "commodo", "consequat", "duis", "aute", "irure",
		"dolor", "in", "reprehenderit", "in", "voluptate", "velit", "esse",
		"cillum", "dolore", "eu", "fugiat", "nulla", "pariatur", "excepteur",
		"sint", "occaecat", "cupidatat", "non", "proident", "sunt", "in",
		"culpa", "qui", "officia", "deserunt", "mollit", "anim", "id", "est",
		"laborum",
	}

	// minOccurrences - minimum occurrences of each word
	minOccurrences = 20
)

type BeefController struct {
	extractor  *service.Extractor
	baconipsum *service.BaconipsumClient
}

func NewBeefController(
	extractor *service.Extractor,
	baconipsum *service.BaconipsumClient,
) *BeefController {
	return &BeefController{
		extractor:  extractor,
		baconipsum: baconipsum,
	}
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
	text, err := ctrl.baconipsum.Get(beefType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error retrieving data from baconipsum.com api",
		})
		return
	}

	// extract words
	text = strings.ToLower(text)
	words := ctrl.extractor.Extract(text, minOccurrences, append(excludes, fillers...))

	// build response body
	result := make(map[string]int)
	for _, word := range words {
		result[word.Phrase] = word.Count
	}

	c.JSON(http.StatusOK, gin.H{
		"beef": result,
	})
}

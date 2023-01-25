package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestSample struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Headers string `json:"header"`
	Body    string `json:"body"`
	Method  string `json:"method"`
	Url     string `json:"url"`
}

var testSamples = []TestSample{
	{Id: "CodeName",
		Name:    "Example API Call",
		Headers: "Hello Jacob and Paul",
		Body:    "",
		Method:  "",
		Url:     ""},
}

func getTests(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, testSamples)
}

func postTest(c *gin.Context) {
	var newTest TestSample

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newTest); err != nil {
		return
	}

	// Add the new album to the slice.
	testSamples = append(testSamples, newTest)
	c.IndentedJSON(http.StatusCreated, testSamples)
}

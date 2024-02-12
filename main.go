package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

type FormData struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	N      int64  `json:"n"`
	Size   string `json:"size"`
	Style  string `json:"style"`
}

type Result struct {
	Created int64      `json:"created"`
	Data    []DataItem `json:"data"`
}

type DataItem struct {
	URL string `json:"url"`
}

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file ", err)
		return
	}

	// load OpenAI key
	apiKey := "Bearer " + os.Getenv("OPEN_AI_KEY")

	// create a Gin router
	router := gin.Default()

	// Web pages route
	router.LoadHTMLGlob("templates/*")

	router.Static("/assets", "./assets")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"description": "A test representation of making use of the OpenAI image generator API",
		})
	})

	// Post to OpenAI endpoint
	router.POST("/api/getImage", func(ctx *gin.Context) {

		// Get payload from the body
		var data FormData

		// Bind the Jsom from the request body
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//call the openai endpoint
		client := resty.New()

		// Make POST request to OPEN AI
		response, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("Authorization", apiKey).
			SetBody(data).
			Post("https://api.openai.com/v1/images/generations")

		// If there is an error, return 500
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong", "error": err})
			return
		}

		r := response.String()

		var result Result

		err1 := json.Unmarshal([]byte(r), &result)

		if err1 != nil {
			fmt.Println("Error:", err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": result})

	})

	router.Run(":8080")

}

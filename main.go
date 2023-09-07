package main

import (
	"application/middleware"
	"application/pkg/setting"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
}

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.GET("/ready", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.GET("/uselessfact", func(c *gin.Context) {
		resp, err := forwardRequest(*setting.AppSetting)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if value, exists := resp["value"]; exists {
			c.JSON(http.StatusOK, value)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Request is not supported"})
		}

	})

	r.GET("/funnyfact", func(c *gin.Context) {
		resp, err := forwardRequest(*setting.AppSetting)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if text, exists := resp["text"]; exists {
			c.JSON(http.StatusOK, text)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Request is not supported"})
		}

	})

	err := r.Run()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func forwardRequest(app setting.App) (map[string]interface{}, error) {
	resp, err := http.Get(app.RequestUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, err
}

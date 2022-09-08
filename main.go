package main

import (
	"net/http"

	"food/env"
	"food/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	baseUrl := env.GetInstance().BaseURL

	r.GET("/", (func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	}))

	r.GET(baseUrl+"/recipes", handlers.GetRecipes)
	r.GET(baseUrl+"/recipeById", handlers.GetRecipe)

	r.Run(env.GetInstance().Port)
}

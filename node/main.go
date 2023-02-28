package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorgonia/gorgonia"
)

// Define a struct to hold the machine learning model
type Model struct {
	ID           string
	Name         string
	Description  string
	Architecture string
	WeightsHash  string
	Graph        *gorgonia.ExprGraph
	Model        *gorgonia.Node
}

// Define a slice to hold all the models
var models []Model

// Define a function to add a new model to the slice
func addModel(id string, name string, description string, architecture string, weightsHash string, graph *gorgonia.ExprGraph, model *gorgonia.Node) {
	newModel := Model{
		ID:           id,
		Name:         name,
		Description:  description,
		Architecture: architecture,
		WeightsHash:  weightsHash,
		Graph:        graph,
		Model:        model,
	}
	models = append(models, newModel)
}

// Define a function to handle the root endpoint
func handleRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to MeshAI!",
	})
}

// Define a function to handle the models endpoint
func handleModels(c *gin.Context) {
	c.JSON(http.StatusOK, models)
}

func main() {
	// Create a new router
	router := gin.Default()

	// Define the root endpoint
	router.GET("/", handleRoot)

	// Define the models endpoint
	router.GET("/models", handleModels)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

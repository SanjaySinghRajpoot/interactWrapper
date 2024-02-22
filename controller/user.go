package controller

import (
	"net/http"

	"github.com/SanjaySinghRajpoot/interactWrapper/utils"
	"github.com/gin-gonic/gin"
)

// GetURLTemp holds the URL retrieved from some source
var GetURLTemp string

// GetData holds the interactions data for each URL
var GetData = make(map[string][]string)

// interactionPayload represents the structure of interaction payload received
type InteractionPayload struct {
	URL       string `json:"url"`        // URL of the interaction
	TimeStamp string `json:"time_stamp"` // Timestamp of the interaction
}

// GetURL handles the GET request for retrieving URL
func GetURL(ctx *gin.Context) {

	if GetURLTemp == "" {
		ctx.JSON(http.StatusNotFound, "URL not found")

	}

	ctx.JSON(http.StatusOK, GetURLTemp)
}

// GetInteractions handles the POST request for retrieving interactions
func GetInteractions(ctx *gin.Context) {
	var payload InteractionPayload

	// Bind the JSON payload to InteractionPayload struct
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Retrieve interactions data for the given URL
	interactions := GetData[payload.URL]

	// If interactions data exist for the URL
	if len(interactions) > 0 {
		if payload.TimeStamp != "" {
			// Filter interactions based on the provided timestamp
			filteredInteractions := utils.FilterByTimestamp(interactions, payload.TimeStamp)

			// Respond with filtered interactions
			ctx.JSON(http.StatusOK, filteredInteractions)
			return
		}

		// Respond with all interactions for the URL
		ctx.JSON(http.StatusOK, interactions)
		return
	}

	// No interactions data found for the URL
	ctx.JSON(http.StatusNotFound, "Interactions not found")
}

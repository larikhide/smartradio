package vote_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VoteCategory struct {
	ID   string `json:"id"`
	Name string `json:"name" binding:"required"`
}

// RetrieveVotecategories returns music categories which user can choose
func (h *VoteAPIHandler) RetrieveVoteCategories(c *gin.Context) {
	// TODO call service method to fetch data from storage

	c.JSON(http.StatusOK, []VoteCategory{
		{ID: "c100", Name: "cheerful"},
		{ID: "c101", Name: "relaxed"},
		{ID: "c102", Name: "lyrical"},
	})
}

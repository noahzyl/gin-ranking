/*
 * Handle HTTP requests about ranking
 */

package controllers

import (
	"github.com/gin-gonic/gin"
)

type RankingController struct{}

type RankingSearch struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (r *RankingController) GetRankingList(ctx *gin.Context) {
	search := &RankingSearch{}
	err := ctx.BindJSON(search) // Parse parameters in the request
	if err != nil {
		ReturnErrorJson(ctx, 4001, gin.H{"error": err.Error()})
		return
	}
	ReturnSuccessJson(ctx, 0, search.Name, search.Cid, 1)
}

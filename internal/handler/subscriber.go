package handler

import (
	"fmt"
	"github.com/6a6ydoping/deeznuts_push_service/api"
	"github.com/6a6ydoping/deeznuts_push_service/internal/entity"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *Handler) subscriber(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")

	htmlContent, err := ioutil.ReadFile("subscriber/home.html")
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "HTML file not found",
		})
		return
	}

	ctx.String(http.StatusOK, string(htmlContent))
}

func (h *Handler) sendPushMessage(ctx *gin.Context) {
	var payload entity.PushPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}
	msg, err := h.messageService.SendMessage(ctx, &payload)
	if err != nil {
		log.Printf("send message error: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}
	fmt.Println(msg)
}

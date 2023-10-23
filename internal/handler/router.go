package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	// сервим js файл для регистрации веб клиента
	router.GET("/service-worker.js", func(c *gin.Context) {
		c.File("subscriber/service-worker.js")
	})

	apiV1 := router.Group("/api/v1")

	apiV1.GET("/service", h.subscriber)
	apiV1.POST("/send", h.sendPushMessage)

	return router
}

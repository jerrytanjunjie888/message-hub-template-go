package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jerry/message-hub-template-go/template/delivery"
	"github.com/jerry/message-hub-template-go/template/service"
)

func (ds *dserver) MapRoutes() {
	// Group routes for v1

	apiV1 := ds.router.Group("api/v1")
	ds.eventRoutes(apiV1)
}

func (ds *dserver) eventRoutes(api *gin.RouterGroup) {
	eventRoutes := api.Group("/events")
	{
		var eventSvc service.Service
		ds.cont.Invoke(func(u service.Service) {
			eventSvc = u
		})

		event := delivery.NewEventController(eventSvc)

		eventRoutes.GET("/", event.GetAll)
		eventRoutes.POST("/", event.Create)
		eventRoutes.GET("/:id", event.GetByID)
		eventRoutes.PUT("/:id", event.Update)
		eventRoutes.DELETE("/:id", event.Delete)
	}
}

package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerry/message-hub-template-go/template/model"
	"github.com/jerry/message-hub-template-go/template/service"
	uuid "github.com/satori/go.uuid"
)

type eventController struct {
	// put logs here
	svc service.Service
}

func NewEventController(svc service.Service) *eventController {
	return &eventController{svc}
}

func (e *eventController) GetAll(ctx *gin.Context) {
	events, err := e.svc.GetAll()
	if len(events) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.SecureJSON(http.StatusOK, events)
}

func (e *eventController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	event, err := e.svc.GetByID(id)
	if event == nil || err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.SecureJSON(http.StatusOK, event)
}

func (e *eventController) Create(ctx *gin.Context) {
	var event model.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	e.svc.Create(&event)
	ctx.Status(http.StatusCreated)
}

// update func goes here
func (e *eventController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var event model.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.ID = id
	e.svc.Update(&event)
	ctx.Status(http.StatusOK)
}

func (e *eventController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	e.svc.Delete(id)
	ctx.Status(http.StatusNoContent)
}

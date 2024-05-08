package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/quminzhi/go-examples/gin-example/models"
	"sync"
)

var g *generator = &generator{}

type VideoController interface {
	GetAll(context *gin.Context)
	Update(context *gin.Context)
	Create(context *gin.Context)
	Delete(context *gin.Context)
}

type controllerImpl struct {
	// Dummy data
	Videos []models.Video
}

func NewVideoController() VideoController {
	return &controllerImpl{make([]models.Video, 0)}
}

func (c *controllerImpl) GetAll(context *gin.Context) {
	context.JSON(200, c.Videos)
}

func (c *controllerImpl) Update(context *gin.Context) {
	var videoToUpdate models.Video
	if err := context.ShouldBindUri(&videoToUpdate); err != nil {
		context.String(400, "bad request %v", err)
		return
	}
	if err := context.ShouldBindJSON(&videoToUpdate); err != nil {
		context.String(400, "bad request %v", err)
		return
	}
	for idx, video := range c.Videos {
		if video.Id == videoToUpdate.Id {
			c.Videos[idx] = videoToUpdate
			context.JSON(200, c.Videos)
			return
		}
	}
	context.String(400, "bad request: not found")
}

func (c *controllerImpl) Create(context *gin.Context) {
	id := g.getNext()
	video := models.Video{Id: id}
	if err := context.ShouldBind(&video); err != nil {
		context.String(400, "bad request %v", err)
		return
	}
	c.Videos = append(c.Videos, video)
	context.String(200, "Successfully created")
}

func (c *controllerImpl) Delete(context *gin.Context) {
	var videoToDelete models.Video
	if err := context.ShouldBindUri(&videoToDelete); err != nil {
		context.String(400, "bad request %v", err)
		return
	}
	for idx, video := range c.Videos {
		if video.Id == videoToDelete.Id {
			c.Videos = append(c.Videos[:idx], c.Videos[idx+1:]...)
			context.String(200, "Successfully deleted")
			return
		}
	}
	context.String(400, "bad request: failed to delete")
}

type generator struct {
	counter int
	sync.Mutex
}

func (g *generator) getNext() int {
	g.Lock()
	defer g.Unlock()
	g.counter++
	return g.counter
}

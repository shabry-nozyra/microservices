package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shabry-nozyra/microservices/helpers/log"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

type Context struct {
	DB  *gorm.DB
	Log *log.AppLog
	Gin *gin.Engine
	Client *resty.Client
}


func (ctx *Context) Register(group string) {
	public := ctx.Gin.Group(group)
	{
		public.GET("/suara", ctx.getAllSuara)
		public.GET("/suara/:id", ctx.getSuara)
		public.POST("/suara/add", ctx.createSuara)
		public.PUT("/suara/update", ctx.updateSuara)
		public.DELETE("/suara/delete/:id", ctx.deleteSuara)
	}
}




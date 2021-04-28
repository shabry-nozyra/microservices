package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shabry-nozyra/microservices/helpers/log"
	"github.com/go-resty/resty/v2"
	"github.com/shabry-nozyra/microservices/notifier"
	"gorm.io/gorm"
)

type Context struct {
	DB  *gorm.DB
	Log *log.AppLog
	Gin *gin.Engine
	Client *resty.Client
	NotifierClient *notifier.Notifier
}


func (ctx *Context) Register(group string) {
	public := ctx.Gin.Group(group)
	{
		public.GET("/suara", ctx.getAllSuaraByTPS)
		public.GET("/totalsuarakec1", ctx.TotalSuaraByKec1)
		public.GET("/totalsuarakec2", ctx.TotalSuaraByKec2)
		public.GET("/totalsuarakec3", ctx.TotalSuaraByKec3)
		public.GET("/totalsuarakec4", ctx.TotalSuaraByKec4)
		public.GET("/totalsuarakec5", ctx.TotalSuaraByKec5)
		public.GET("/totalsuarakec6", ctx.TotalSuaraByKec6)
		public.GET("/totalsuarakec7", ctx.TotalSuaraByKec7)
		public.GET("/suaratotal", ctx.getTotalTpsMasuk)
		public.GET("/suaras", ctx.getAllSuara)
		public.GET("/suara/:id", ctx.getSuara)
		public.POST("/suara/add", ctx.createSuara)
		public.PUT("/suara/update", ctx.updateSuara)
		public.DELETE("/suara/delete/:id", ctx.deleteSuara)
	}
}




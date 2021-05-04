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
		public.GET("/suara", ctx.getAllSuaraByTPS)
		public.GET("/totalsuarakec/:id", ctx.TotalSuaraByKec1)
		public.GET("/suaras", ctx.getAllSuara)
		public.GET("/suaratotal", ctx.getTotalTpsMasuk)
		public.GET("/suara/:id", ctx.getSuara)
		public.GET("/suaraByKec/:kec", ctx.getByKec)
		public.GET("/detailkec/:kec", ctx.DetailKec)
		public.GET("/NamaKec", ctx.getKec)
		public.GET("/NamaNagari", ctx.getNagari)
		public.POST("/suara/add", ctx.createSuara)
		public.PUT("/suara/update", ctx.updateSuara)
		public.DELETE("/suara/delete/:id", ctx.deleteSuara)
	}
}




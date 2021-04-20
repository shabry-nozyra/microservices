package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shabry-nozyra/microservices/handlers"
	"github.com/shabry-nozyra/microservices/helpers/env"
	"github.com/shabry-nozyra/microservices/helpers/log"
	"github.com/shabry-nozyra/microservices/models"
	"github.com/go-resty/resty/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func main()  {
	l := log.NewLog("Log", "")
	defer l.Close()

	cs := env.Get().ConnectionString()
	db, err := gorm.Open(postgres.Open(cs), &gorm.Config{})

	if err != nil{
		l.Fatal("Koneksi Ke Database Gagal")

		return
	}


	err = models.MigrateModel(db)
	if err != nil{
		l.Error("AutoMigrate Gagal")
		return
	}

	client := resty.New()
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()

	c := cors.DefaultConfig()
	c.AllowWildcard = true
	c.AllowOrigins = []string{"*"}
	c.AddAllowHeaders("Authorization", "Content-Type")
	c.AddExposeHeaders("Authorization", "Content-Type")
	g.Use(cors.New(c))

	h := handlers.Context{Gin: g, DB: db, Log:l, Client: client}
	h.Register("")

	port := os.Getenv("AppPort")

	l.Infof("start listen and serve at %v", port)
	s := &http.Server{Addr: "0.0.0.0:" + port, Handler: g}
	err = s.ListenAndServe()
	if err != nil {
		l.Fatal("failed to connect to serv")
		return
	}
}

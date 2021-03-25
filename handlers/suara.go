package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shabry-nozyra/microservices/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)



//getAll
func (ctx *Context) getAllSuara(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p := models.TPS{}
	err := p.GetTPS(ctx.Client, id)
	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}

	m := models.Suaras{}
	err = m.All(ctx.DB, id)

	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	c.JSON(http.StatusOK, &m)
}


func (ctx *Context) getSuara(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p := models.Suara{}

	err := p.Get(ctx.DB, id)

	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	c.JSON(http.StatusOK, p)
}

func (ctx *Context) createSuara(c *gin.Context) {
	p := models.Suara{}
	err := c.ShouldBindJSON(&p)

	err = p.Create(ctx.DB)
	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusBadRequest, errorm)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	c.JSON(http.StatusCreated, res)

}

func (ctx *Context) deleteSuara(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p := models.Suara{}

	err := p.Get(ctx.DB, id)

	if err != nil{
		ctx.Log.Warn("Bad Request")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err = p.Delete(ctx.DB, id)
	if err != nil{
		errorm := "Gagal Menjalankan Query"
		ctx.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, errorm)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	c.JSON(http.StatusOK, res)
}

func (ctx *Context) updateSuara(c *gin.Context) {
	cl := models.Suara{}
	err := c.ShouldBindJSON(&cl)
	if err != nil{
		ctx.Log.Warn(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	p := models.Suara{}
	err = p.Get(ctx.DB, int(cl.ID))

	if err != nil{
		if err == gorm.ErrRecordNotFound{
			ctx.Log.Warn(err)
			c.JSON(http.StatusNotFound, nil)
			return
		}else{
			ctx.Log.Warn(err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
	}

	p.Suara1 = cl.Suara1
	p.Suara2 = cl.Suara2
	p.Suara3 = cl.Suara3
	p.SuaraSah = cl.SuaraSah
	p.JPL = cl.JPL
	p.SuaraTidakSah = cl.SuaraTidakSah
	p.JumlahGolput = cl.JumlahGolput
	p.FotoC1 =  cl.FotoC1
	p.Status = cl.Status
	p.WaktuInput = cl.WaktuInput
	p.DataKe = cl.DataKe


	err = p.Update(ctx.DB)
	if err != nil{
		errorm := "Gagal Menjalankan Query"
		ctx.Log.Error(err)
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	c.JSON(http.StatusCreated, res)
}


package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pusher/pusher-http-go"
	"github.com/shabry-nozyra/microservices/models"
	"github.com/shabry-nozyra/microservices/helpers/log"
	"gorm.io/gorm"
	"math"
	"net/http"
	"strconv"
	"strings"
)

var l = log.NewLog("Log", "")

var client = pusher.Client{
	AppID:   "1192937",
	Key:     "d78a56b5e5e7ed44eeb4",
	Secret:  "a137aa3643740f6b1945",
	Cluster: "ap1",
	Secure:  true,
}





//getAll
func (ctx *Context) getAllSuaraByTPS(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	p := models.TPS{}
	err := p.GetTPS(ctx.Client, id)
	if err != nil{
		if strings.Contains(err.Error(), "404"){
			ctx.Log.Warn(err)
			errorm := "Data Not Found"
			c.JSON(http.StatusNotFound, errorm)
			return
		}else{
			ctx.Log.Error(err)
			errorm := "Gagal Menjalankan Query"
			c.JSON(http.StatusInternalServerError, errorm)
			return
		}

	}

	m := models.Suaras{}
	err = m.AllByTPS(ctx.DB, id)

	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	c.JSON(http.StatusOK, m)
}

func (ctx *Context) getAllSuara(c *gin.Context) {
	s := models.Suaras{}
	err := s.All(ctx.DB)

	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	c.JSON(http.StatusOK, s)
}

func (ctx *Context) getKec(c *gin.Context) {
	s := models.Kecs{}
	err := s.GetKec(ctx.DB)

	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	//kecamatan := float64(s.Count(ctx.DB))

	c.JSON(http.StatusOK, s)
}

func (ctx *Context) getNagari(c *gin.Context) {
	s := models.Nagaris{}
	err := s.GetNagari(ctx.DB)

	if err != nil{
		ctx.Log.Error(err.Error())
		errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, errorm)
		return
	}
	c.JSON(http.StatusOK, s)
}


func (ctx *Context) getSuara(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p := models.Suara{}

	err := p.Get(ctx.DB, id)

	if err != nil{
		ctx.Log.Error(err.Error())
		//errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, p)
}

func (ctx *Context) getByKec(c *gin.Context) {
	Kec := c.Param("kec")
	p := models.Suaras{}

	err := p.GetByKec(ctx.DB, Kec)

	if err != nil{
		ctx.Log.Error(err.Error())
		//errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, p)
}

func (ctx *Context) TotalSuaraByKec1(c *gin.Context) {
	r := models.Result{}
	err := r.TotalSuaraByKec1(ctx.DB)
	if err != nil{
		ctx.Log.Error(err.Error())
		//errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Log.Info(r)
	c.JSON(http.StatusOK, r)
}
func (ctx *Context) TotalSuaraByKec2(c *gin.Context) {
	r := models.Result{}
	err := r.TotalSuaraByKec2(ctx.DB)
	if err != nil{
		ctx.Log.Error(err.Error())
		//errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Log.Info(r)
	c.JSON(http.StatusOK, r)
}
func (ctx *Context) TotalSuaraByKec3(c *gin.Context) {
	r := models.Result{}
	err := r.TotalSuaraByKec3(ctx.DB)
	if err != nil{
		ctx.Log.Error(err.Error())
		//errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Log.Info(r)
	c.JSON(http.StatusOK, r)
}
func (ctx *Context) TotalSuaraByKec4(c *gin.Context) {
	r := models.Result{}
	err := r.TotalSuaraByKec4(ctx.DB)
	if err != nil{
		ctx.Log.Error(err.Error())
		//errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Log.Info(r)
	c.JSON(http.StatusOK, r)
}
func (ctx *Context) TotalSuaraByKec5(c *gin.Context) {
	r := models.Result{}
	err := r.TotalSuaraByKec5(ctx.DB)
	if err != nil{
		ctx.Log.Error(err.Error())
		//errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Log.Info(r)
	c.JSON(http.StatusOK, r)
}
func (ctx *Context) TotalSuaraByKec6(c *gin.Context) {
	r := models.Result{}
	err := r.TotalSuaraByKec6(ctx.DB)
	if err != nil{
		ctx.Log.Error(err.Error())
		//errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Log.Info(r)
	c.JSON(http.StatusOK, r)
}
func (ctx *Context) TotalSuaraByKec7(c *gin.Context) {
	r := models.Result{}
	err := r.TotalSuaraByKec7(ctx.DB)
	if err != nil{
		ctx.Log.Error(err.Error())
		//errorm := "Gagal Menjalankan Query"
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Log.Info(r)
	c.JSON(http.StatusOK, r)
}
func (ctx *Context) getTotalTpsMasuk(c *gin.Context) {
	p := models.Suaras{}
	count := float64(p.Count(ctx.DB))
	suara1 := float64(p.Sum(ctx.DB))
	suara2 := float64(p.Sum2(ctx.DB))
	suara3 := float64(p.Sum3(ctx.DB))
	totalSuara := float64(suara1 + suara2 + suara3)
	totalTps := float64(p.TotalTps(ctx.DB))
	persentaseSuaraMasuk := float64(math.Round(count*100/totalTps))
	persenSuara1 := float64(math.Round(float64(suara1*100/totalSuara)))
	persenSuara2 := float64(math.Round(float64(suara2*100/totalSuara)))
	persenSuara3 := float64(math.Round(float64(suara3*100/totalSuara)))

	res := map[string]float64{
		"total_tps_masuk": count,
		"total_suara_masuk": totalSuara,
		"total_tps": totalTps,
		"persentase_suara_masuk": persentaseSuaraMasuk,
		"suara1": suara1,
		"suara2": suara2,
		"suara3": suara3,
		"persensuara1": persenSuara1,
		"persensuara2": persenSuara2,
		"persensuara3": persenSuara3,
	}
	c.JSON(http.StatusOK, res)
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

	m := models.Suaras{}
	count := float64(m.Count(ctx.DB))
	suara1 := float64(m.Sum(ctx.DB))
	suara2 := float64(m.Sum2(ctx.DB))
	suara3 := float64(m.Sum3(ctx.DB))
	totalSuara := float64(suara1 + suara2 + suara3)
	totalTps := float64(m.TotalTps(ctx.DB))
	persentaseSuaraMasuk := float64(math.Round(count*100/totalTps))
	persenSuara1 := float64(math.Round(float64(suara1*100/totalSuara)))
	persenSuara2 := float64(math.Round(float64(suara2*100/totalSuara)))
	persenSuara3 := float64(math.Round(float64(suara3*100/totalSuara)))

	result := map[string]float64{
		"total_tps_masuk": count,
		"total_suara_masuk": totalSuara,
		"total_tps": totalTps,
		"persentase_suara_masuk": persentaseSuaraMasuk,
		"suara1": suara1,
		"suara2": suara2,
		"suara3": suara3,
		"persensuara1": persenSuara1,
		"persensuara2": persenSuara2,
		"persensuara3": persenSuara3,
	}

	client.Trigger("results", "results", result)
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
func (ctx *Context) DetailKec(c *gin.Context) {
	Kec := c.Param("kec")
	p := models.Suaras{}
	nagari := float64(p.JumlahNagariByKec(ctx.DB,Kec))
	tps := float64(p.JumlahTPSByKec(ctx.DB,Kec))
	jpl := float64(p.TotalJplByKec(ctx.DB, Kec))

	res := map[string]float64{
		"total_nagari": nagari,
		"total_tps": tps,
		"total_jpl": jpl,

	}
	c.JSON(http.StatusOK, res)
}


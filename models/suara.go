package models

import (
	"gorm.io/gorm"
)


type (
	Suara struct {
		ID uint 				`json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		Id_tps int64 			`json:"id_tps" gorm:"column:id_tps;"`
		Suara1 int64 			`json:"suara1" gorm:"column:suara1;"`
		Suara2 int64 			`json:"suara2" gorm:"column:suara2;"`
		Suara3 int64 			`json:"suara3" gorm:"column:suara3;"`
		SuaraSah int64 			`json:"suara_sah" gorm:"column:suara_sah;"`
		JPL int64 				`json:"jpl" gorm:"column:jpl;"`
		SuaraTidakSah int64 	`json:"suara_tidak_sah" gorm:"column:suara_tidak_sah;"`
		JumlahGolput int64 		`json:"jumlah_golput" gorm:"column:jumlah_golput;"`
		FotoC1 string 			`json:"c1" gorm:"column:c1;"`
		Status int64			`json:"status" gorm:"column:status;"`
		WaktuInput int64		`json:"waktu_input" gorm:"column:waktu_input;"`
		DataKe int64 			`json:"data_ke" gorm:"column:data_ke;"`
		Lokasi string 			`json:"lokasi"`
		Nagari string			`json:"nagari"`
		No_tps int64			`json:"no_tps"`
	}
)
func (Suara) TableName() string{
	return "suara_tps" //nama table di database
}
func (p *Suara) Get(db *gorm.DB, ID int) error {
	return db.Model(Suara{}).Where("id = ?", ID).First(p).Error
}


func (p *Suara) Create(db *gorm.DB) error {
	return db.Model(Suara{}).Create(p).Error
}

func (p *Suara) Update(db *gorm.DB) error {
	return db.Model(Suara{}).Where("id = ?", p.ID).Updates(p).Error
}

func (p *Suara) Delete(db *gorm.DB, ID int) error {
	return db.Model(Suara{}).Where("id = ?", ID).Delete(p).Error
}

func (p *Suaras) Count(db *gorm.DB) int64{
		var result int64
		db.Model(&Suaras{}).Distinct("id").Count(&result)
		return result
}
func (p *Suaras) Sum(db *gorm.DB) int64{
	var sum int64
	db.Table("suara_tps").Select("sum(suara1)").Row().Scan(&sum)
	return sum
}

func (p *Suaras) Sum2(db *gorm.DB) int64{
	var sum int64
	db.Table("suara_tps").Select("sum(suara2)").Row().Scan(&sum)
	return sum
}
func (p *Suaras) Sum3(db *gorm.DB) int64{
	var sum int64
	db.Table("suara_tps").Select("sum(suara3)").Row().Scan(&sum)
	return sum
}

func  (p *Suaras) TotalTps(db *gorm.DB) int64{
	var tps int64
	db.Table("tps").Distinct("id").Count(&tps)
	return tps
}

type Result struct {
		Suara1  int64 `json:"suara1"`
		Suara2 int64 `json:"suara2"`
		Suara3 int64 `json:"suara3"`
}

func  (r *Result) TotalSuaraByKec1 (db *gorm.DB) error {
	return db.Raw("Select sum(suara_tps.suara1) as suara1, sum(suara_tps.suara2) as suara2, sum(suara_tps.suara3) as suara3 from suara_tps JOIN tps ON tps.id = suara_tps.id_tps where tps.kecamatan = 'koto parik gadang diateh'").Scan(&r).Error
}
func  (r *Result) TotalSuaraByKec2 (db *gorm.DB) error {
	return db.Raw("Select sum(suara_tps.suara1) as suara1, sum(suara_tps.suara2) as suara2, sum(suara_tps.suara3) as suara3 from suara_tps JOIN tps ON tps.id = suara_tps.id_tps where tps.kecamatan = 'sungai pagu'").Scan(&r).Error
}
func  (r *Result) TotalSuaraByKec3 (db *gorm.DB) error {
	return db.Raw("Select sum(suara_tps.suara1) as suara1, sum(suara_tps.suara2) as suara2, sum(suara_tps.suara3) as suara3 from suara_tps JOIN tps ON tps.id = suara_tps.id_tps where tps.kecamatan = 'pauh duo'").Scan(&r).Error
}
func  (r *Result) TotalSuaraByKec4 (db *gorm.DB) error {
	return db.Raw("Select sum(suara_tps.suara1) as suara1, sum(suara_tps.suara2) as suara2, sum(suara_tps.suara3) as suara3 from suara_tps JOIN tps ON tps.id = suara_tps.id_tps where tps.kecamatan = 'sangir'").Scan(&r).Error
}
func  (r *Result) TotalSuaraByKec5 (db *gorm.DB) error {
	return db.Raw("Select sum(suara_tps.suara1) as suara1, sum(suara_tps.suara2) as suara2, sum(suara_tps.suara3) as suara3 from suara_tps JOIN tps ON tps.id = suara_tps.id_tps where tps.kecamatan = 'sangir jujuan'").Scan(&r).Error
}
func  (r *Result) TotalSuaraByKec6 (db *gorm.DB) error {
	return db.Raw("Select sum(suara_tps.suara1) as suara1, sum(suara_tps.suara2) as suara2, sum(suara_tps.suara3) as suara3 from suara_tps JOIN tps ON tps.id = suara_tps.id_tps where tps.kecamatan = 'sangir balai janggo'").Scan(&r).Error
}
func  (r *Result) TotalSuaraByKec7 (db *gorm.DB) error {
	return db.Raw("Select sum(suara_tps.suara1) as suara1, sum(suara_tps.suara2) as suara2, sum(suara_tps.suara3) as suara3 from suara_tps JOIN tps ON tps.id = suara_tps.id_tps where tps.kecamatan = 'sangir batang hari'").Scan(&r).Error
}


type Suaras []Suara

func (p *Suaras) AllByTPS(db *gorm.DB, Id_tps int) error {
	return db.Model(Suara{}).Where("id_tps = ?", Id_tps).Find(p).Error
}

func (p *Suaras) All(db *gorm.DB) error {
	return db.Model(Suara{}).Find(p).Error
}

func (p *Suaras) GetByKec(db *gorm.DB, Kec string) error {
	return db.Table("suara_tps").Select("*").Joins("join tps on suara_tps.id_tps = tps.id").Where("tps.kecamatan = ?", Kec).Scan(&p).Error

	//return db.Model(Suara{}).Where("id_tps = ?", Id_tps).Find(p).Error
}


func (p *Suaras) JumlahNagariByKec(db *gorm.DB, Kec string) int64 {
	var nagari int64
	db.Table("tps").Distinct("nagari").Where("tps.kecamatan = ?", Kec).Count(&nagari)
	return nagari
}
func (p *Suaras) JumlahTPSByKec(db *gorm.DB, Kec string) int64 {
	var tps int64
	db.Table("tps").Distinct("id").Where("tps.kecamatan = ?", Kec).Count(&tps)
	return tps
}
func  (p *Suaras) TotalJplByKec (db *gorm.DB, Kec string) int64 {
	var JPL int64
	db.Raw("Select sum(jpl) as jpl from tps").Where("tps.kecamatan = ?", Kec).Scan(&JPL)
	return JPL
}


type Kec struct {
	Kecamatan  string `json:"kecamatan"`
}

type Kecs []Kec


type Nagari struct {
	Nagari string `json:"nagari"`
}

type Nagaris []Nagari

func (k *Kecs) GetKec(db *gorm.DB) error {
	return db.Table("tps").Distinct("kecamatan").Select("kecamatan").Find(k).Error	//return db.Model(Suara{}).Where("id_tps = ?", Id_tps).Find(p).Error
}

func (n *Nagaris) GetNagari(db *gorm.DB) error {
	return db.Table("tps").Distinct("nagari").Select("nagari").Find(n).Error	//return db.Model(Suara{}).Where("id_tps = ?", Id_tps).Find(p).Error
}







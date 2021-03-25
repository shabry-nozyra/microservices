package models

import (
	"gorm.io/gorm"
)

type (
	Suara struct {
		ID uint 				`json:"id" gorm:"column:id;primaryKey;autoIncrement"`
		IdTps uint 				`json:"idtps" gorm:"column:idtps;"`
		Suara1 uint8 			`json:"suara1" gorm:"column:suara1;"`
		Suara2 uint8 			`json:"suara2" gorm:"column:suara2;"`
		Suara3 uint8 			`json:"suara3" gorm:"column:suara3;"`
		SuaraSah uint8 			`json:"suarasah" gorm:"column:suarasah;"`
		JPL uint8 				`json:"jpl" gorm:"column:jpl;"`
		SuaraTidakSah uint8 	`json:"suaratidaksah" gorm:"column:suaratidaksah;"`
		JumlahGolput uint 		`json:"jumlahgolput" gorm:"column:jumlahgolput;"`
		FotoC1 string 			`json:"fotoc1" gorm:"column:fotoc1;"`
		Status uint8			`json:"status" gorm:"column:status;"`
		WaktuInput uint8		`json:"waktuinput" gorm:"column:waktuinput;"`
		DataKe uint8 			`json:"datake" gorm:"column:datake;"`
	}
)
func (Suara) TableName() string{
	return "suara" //nama table di database
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

type Suaras []Suara

func (p *Suaras) All(db *gorm.DB, IdTps int) error {
	return db.Model(Suara{}).Where("id = ?", IdTps).Find(p).Error
}






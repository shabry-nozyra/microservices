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

type Suaras []Suara

func (p *Suaras) All(db *gorm.DB) error {
	return db.Model(Suara{}).Find(p).Error
}






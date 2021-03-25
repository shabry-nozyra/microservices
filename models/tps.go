package models

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/go-resty/resty/v2"
	"net/http"
)
type (
	TPS struct {
		ID uint `json:"id"`
		NoTPS uint8 `json:"no_tps"`
		Lokasi string `json:"lokasi"`
		Kecamatan string `json:"kecamatan"`
		Nagari string `json:"nagari"`
		Jorong string `json:"jorong"`
		JPL uint8 `json:"jpl"`
		IsActive uint8 `json:"is_active"`
	}
)

func (p *TPS) GetTPS(client *resty.Client, ID int) error {
	const URL =  "http://localhost:8080/tps/id"
	res, err := client.R().Get(fmt.Sprintf(URL, ID))
	if err != nil{
		return err
	}

	if res.StatusCode() != http.StatusOK{
		return fmt.Errorf(res.Status())
	}
	err = jsoniter.Unmarshal(res.Body(), p)
	if err != nil{
		return err
	}

	return nil

}

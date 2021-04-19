package models

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)
type (
	TPS struct {
		ID int64 `json:"id"`
		NoTPS int64 `json:"no_tps"`
		Lokasi string `json:"lokasi"`
		Kecamatan string `json:"kecamatan"`
		Nagari string `json:"nagari"`
		Jorong string `json:"jorong"`
		JPL int64 `json:"jpl"`
		IsActive int64 `json:"is_active"`
	}
)

func (p *TPS) GetTPS(client *resty.Client, ID int) error {
	const URL =  "http://localhost:8080/tps/%d"
	res, err := client.R().Get(fmt.Sprintf(URL, 75))
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

package griddy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"net/http"
)

type Griddy struct {
	Url        string `json:"-" toml"url"`
	Meter      string `json:"meterID" toml:"meterid"`
	Member     string `json:"memberID" toml:"memberid"`
	Settlement string `json:"settlement_point" toml:"settlement"`
}

func (g *Griddy) LoadConfig(path string) error {
	var config struct {
		Griddy *Griddy
	}
	config.Griddy = g
	_, err := toml.DecodeFile(path, &config)
	return err
}

type GriddyResponse struct {
	Now GriddyStatus `json:"now"`
	TTL uint32       `json:"seconds_until_refresh,string"`
}

type GriddyStatus struct {
	Date             string  `json:"date"`
	Hour             uint32  `json:"hour_num,string"`
	Minute           uint32  `json:"min_num,string"`
	Settlement       string  `json:"settlement_point"`
	PriceType        string  `json:"price_type"`
	Price            float64 `json:"price_ckwh,string"`
	ValueScore       uint32  `json:"value_score,string"`
	MeanPrice        float64 `json:"mean_price_ckwh,string"`
	DiffMean         float64 `json:"diff_mean_ckwh,string"`
	High             float64 `json:"high_ckwh,string"`
	Low              float64 `json:"low_ckwh,string"`
	PriceDisplay     float64 `json:"price_display,string"`
	PriceDisplaySign string  `json:"price_display_sign"`
	DateLocalTZ      string  `json:"date_local_tz"`
}

func (g *Griddy) Get() (GriddyResponse, error) {
	response := GriddyResponse{}
	if g.Url == "" || g.Meter == "" || g.Member == "" || g.Settlement == "" {
		return response, fmt.Errorf("griddy: Url, Meter, Member and Settlement fields are required")
	}
	jsonMember, err := json.Marshal(g)
	if err != nil {
		return response, err
	}
	resp, err := http.Post(g.Url, "application/json", bytes.NewBuffer(jsonMember))
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func New() Griddy {
	griddy := Griddy{}
	griddy.Url = "https://app.gogriddy.com/api/v1/insights/getnow"
	return griddy
}

package model

type CountriesJSON struct {
	Errors []string `json:"errors"`
	Get    string   `json:"get"`
	Paging struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Parameters []struct {
		Name   string `json:"name"`
		Code   string `json:"code"`
		Search string `json:"search"`
	}
	Response []struct {
		Code string `json:"code"`
		Flag string `json:"flag"`
		Name string `json:"name"`
	}
	Results int `json:"results"`
}

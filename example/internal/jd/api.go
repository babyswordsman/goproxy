package jd

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Price struct {
	Price struct {
		P string `json:"p"`
	} `json:"price"`
}

func Api(body []byte, source, method string) *Jd {
	var jd = new(Jd)
	if strings.Contains(source, "api.m.jd.com/?appid=pc-item-soa") && method == http.MethodGet {
		var price Price
		_ = json.Unmarshal(body, &price)
		jd.ProductPrice = price.Price.P
	}
	return jd
}

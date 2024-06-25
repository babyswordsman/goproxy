package jd

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Product struct {
	Price struct {
		P string `json:"p"`
	} `json:"price"`
}

type Review struct {
	ImageListCount          int                       `json:"imageListCount"`
	ProductCommentSummary   ProductCommentSummary     `json:"productCommentSummary"`
	HotCommentTagStatistics []HotCommentTagStatistics `json:"hotCommentTagStatistics"`
}

func Api(body []byte, source, method string) *Jd {
	var jd = new(Jd)
	// 商品价格
	if strings.Contains(source, "api.m.jd.com/?appid=pc-item-soa") && method == http.MethodGet {
		var product Product
		_ = json.Unmarshal(body, &product)
		jd.ProductPrice = product.Price.P
	}
	// 商品评价
	if strings.Contains(source, "api.m.jd.com/?appid=item-v3&functionId=pc_club_productPageComments") && method == http.MethodGet {
		var review Review
		_ = json.Unmarshal(body, &review)
		jd.ProductCommentSummary = &review.ProductCommentSummary
		jd.ProductCommentSummary.HotCommentTagStatistics = review.HotCommentTagStatistics
		jd.ProductCommentSummary.ImageListCount = review.ImageListCount
	}
	return jd
}

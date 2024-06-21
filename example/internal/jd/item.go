package jd

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	logger "github.com/sirupsen/logrus"
)

func Item(body []byte, source string) *Jd {
	var jd = new(Jd)
	switch {
	case strings.Contains(source, "item.jd.com"):
		logger.Infof(" -----> %s", source)
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
		if err != nil {
			logger.Warnf("No url found")
		} else {
			// 品牌
			doc.Find("#parameter-brand li a").Each(func(index int, item *goquery.Selection) {
				jd.Brand = strings.TrimSpace(item.Text())
			})
			// 商品名称
			doc.Find(".sku-name").Each(func(index int, item *goquery.Selection) {
				jd.ProductName = strings.TrimSpace(item.Text())
			})
			// 商品编号
			doc.Find(".parameter2.p-parameter-list .shieldShopInfo").Each(func(index int, item *goquery.Selection) {
				if strings.Contains(item.Text(), "商品编号") {
					jd.ProductNumber = strings.TrimSpace(item.Text())
				}
			})
			// todo: 商品价格
			// 商品链接
			doc.Find("link[rel='canonical']").Each(func(index int, item *goquery.Selection) {
				href, exists := item.Attr("href")
				if exists {
					jd.ProductLink = fmt.Sprintf("https:%s\n", href)
				}
			})
			// todo: 一级分类
			// todo: 二级分类
			// todo: 三级分类
			// 所属店铺
			doc.Find(".parameter2.p-parameter-list .shieldShopInfo a").Each(func(index int, item *goquery.Selection) {
				jd.Store = strings.TrimSpace(item.Text())
			})
			// 商品图片
			doc.Find("#spec-img").Each(func(index int, item *goquery.Selection) {
				imgSrc, exists := item.Attr("data-origin")
				if exists {
					jd.ProductPicture.Main = append(jd.ProductPicture.Main, fmt.Sprintf("https:%s\n", imgSrc))
				}
			})
			// todo: 商品评价
			// todo: 扩展属性
			// 售后保障
		}
	}
	return jd
}

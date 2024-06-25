package jd

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	logger "github.com/sirupsen/logrus"
)

var PROPERTIES_PATTERN = regexp.MustCompile(`\S+：\S+`)

func Item(body []byte, source string) *Jd {
	var jd = new(Jd)
	switch {
	case strings.Contains(source, "item.jd.com"):
		logger.Infof("-----> %s \n", source)
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
			// todo 商品价格
			//doc.Find("body div.summary-price.J-summary-price").Each(func(index int, item *goquery.Selection) {
			//	jd.ProductPrice = strings.TrimSpace(item.Text())
			//	logger.Infof("GoodsAttr: ProductPrice %s", jd.ProductPrice)
			//})
			// 商品链接
			doc.Find("link[rel='canonical']").Each(func(index int, item *goquery.Selection) {
				href, exists := item.Attr("href")
				if exists {
					jd.ProductLink = fmt.Sprintf("https:%s\n", href)
				}
			})
			// 一级分类
			doc.Find("div.crumb.fl.clearfix div.item.first a").Each(func(index int, item *goquery.Selection) {
				jd.FirstLevel = strings.TrimSpace(item.Text())
			})
			// 二级分类
			doc.Find("div.crumb.fl.clearfix div:nth-child(3) a").Each(func(index int, item *goquery.Selection) {
				jd.SecondLevel = strings.TrimSpace(item.Text())
			})
			// 三级分类
			doc.Find("div.crumb.fl.clearfix div:nth-child(5) a").Each(func(index int, item *goquery.Selection) {
				jd.ThirdLevel = strings.TrimSpace(item.Text())
			})
			// 所属店铺
			doc.Find(".parameter2.p-parameter-list .shieldShopInfo a").Each(func(index int, item *goquery.Selection) {
				jd.Store = strings.TrimSpace(item.Text())
			})
			// 商品图片
			doc.Find("#spec-img").Each(func(index int, item *goquery.Selection) {
				imgSrc, exists := item.Attr("data-origin")
				if exists {
					jd.ProductPicture = new(ProductPicture)
					jd.ProductPicture.Main = append(jd.ProductPicture.Main, fmt.Sprintf("https:%s\n", imgSrc))
				}
			})
			// todo: 商品评价
			// 扩展属性
			doc.Find("div.tab-con div:nth-child(1) div.p-parameter").Each(func(index int, item *goquery.Selection) {
				text := strings.TrimSpace(item.Text())
				matches := PROPERTIES_PATTERN.FindAllString(text, -1)
				attrs := make(map[string]string)
				for _, match := range matches {
					parts := strings.Split(match, "：")
					attrs[parts[0]] = parts[1]

				}
				jd.ExtendedProperties = attrs
			})

			// 售后保障
			doc.Find("div.mc div div.serve-agree-bd dl").Each(func(index int, item *goquery.Selection) {
				jd.SalesGuarantee = strings.TrimSpace(item.Text())
				logger.Infof("GoodsAttr: SalesGuarantee %s", jd.SalesGuarantee)
			})

		}
	}
	return jd
}

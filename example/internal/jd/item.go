package jd

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	logger "github.com/sirupsen/logrus"
)

type Jd struct {
	Brand          string `json:"brand"`          // 品牌
	ProductName    string `json:"product_name"`   // 商品名称
	ProductNumber  string `json:"product_number"` // 商品编号
	ProductPrice   int    `json:"product_price"`  // 商品价格
	ProductLink    string `json:"product_link"`   // 商品链接
	FirstLevel     string `json:"first_level"`    // 一级分类
	SecondLevel    string `json:"second_level"`   // 二级分类
	ThirdLevel     string `json:"third_level"`    // 三级分类
	Store          string `json:"store"`          // 所属店铺
	ProductPicture struct {
		Main  []string `json:"main"`  // 主图
		Other []string `json:"other"` // 其他
	} `json:"product_picture"` // 商品图片
	ProductReview struct {
		Reviews        int      `json:"reviews"`         // 评价数
		PraiseExtent   int      `json:"praise_extent"`   // 好评度
		PraiseNum      int      `json:"praise_num"`      // 好评数
		MidNum         int      `json:"mid_num"`         // 中评数
		NegativeNum    int      `json:"negative_num"`    // 差评数
		PrintsNum      int      `json:"prints_num"`      // 晒图数
		VideoPostsNum  int      `json:"video_posts_num"` // 视频晒单数
		EvaluationTags []string `json:"evaluation_tags"` // 评价标签
	} `json:"product_review"` // 商品评价
	SalesGuarantee     string      `json:"sales_guarantee"` // 售后保障
	ExtendedProperties interface{} `json:"extended_properties"`
	//  "扩展属性": {  //扩展属性根据不同产品类别确定，此示例为电视的扩展属性
	//    "屏幕尺寸": "32英寸及以下",
	//    "电视类型": [
	//      "适老电视",
	//      "智能电视"
	//    ],
	//    "推荐观看距离": "2m以下",
	//    "护眼电视": "护眼电视",
	//    "摄像头": "无摄像头",
	//    "能效等级": "三级能效",
	//    "语音控制": "不支持语音功能",
	//    "组套类型": "无组套",
	//    "刷屏率": "60Hz",
	//    "电视初始内容源": "爱奇艺"
	//  },
	RelevantProduct []struct {
		Name    string `json:"name"`  // 商品名
		Price   int    `json:"price"` // 商品价格
		Link    string `json:"link"`  // 商品链接
		Picture string `json:"picture"`
	} `json:"relevant_product"` //  "相关推荐
}

func Item(body []byte, source string) {
	switch {
	case strings.Contains(source, "item.jd.com"):
		logger.Infof(" -----> %s", source)
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
		if err != nil {
			logger.Warnf("No url found")
		} else {
			var jd Jd
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
			fmt.Printf("-------->> %+v\n", jd)
		}
	}
}

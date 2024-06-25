package jd

type Jd struct {
	Brand                 string                 `json:"brand"`                 // 品牌
	ProductName           string                 `json:"productName"`           // 商品名称
	ProductNumber         string                 `json:"productNumber"`         // 商品编号
	ProductPrice          string                 `json:"productPrice"`          // 商品价格
	ProductLink           string                 `json:"productLink"`           // 商品链接
	FirstLevel            string                 `json:"firstLevel"`            // 一级分类
	SecondLevel           string                 `json:"secondLevel"`           // 二级分类
	ThirdLevel            string                 `json:"thirdLevel"`            // 三级分类
	Store                 string                 `json:"store"`                 // 所属店铺
	ProductPicture        *ProductPicture        `json:"productPicture"`        // 商品图片
	ProductCommentSummary *ProductCommentSummary `json:"productCommentSummary"` // 商品评价
	SalesGuarantee        string                 `json:"sales_guarantee"`       // 售后保障
	ExtendedProperties    interface{}            `json:"extendedProperties"`
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

type ProductPicture struct {
	Main  []string `json:"main"`  // 主图
	Other []string `json:"other"` // 其他
}

type ProductCommentSummary struct {
	CommentCountStr         string                    `json:"commentCountStr"`         // 评价数
	GoodRateShow            int                       `json:"goodRateShow"`            // 好评度
	GoodCountStr            string                    `json:"goodCountStr"`            // 好评数
	GeneralCountStr         string                    `json:"generalCountStr"`         // 中评数
	PoorCountStr            string                    `json:"poorCountStr"`            // 差评数
	ImageListCount          int                       `json:"imageListCount"`          // 晒图数
	VideoCountStr           string                    `json:"videoCountStr"`           // 视频晒单数
	HotCommentTagStatistics []HotCommentTagStatistics `json:"hotCommentTagStatistics"` // 评价标签
}

type HotCommentTagStatistics struct {
	Name string `json:"name"`
}

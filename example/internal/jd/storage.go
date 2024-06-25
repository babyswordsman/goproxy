package jd

import (
	"encoding/json"
	"os"

	logger "github.com/sirupsen/logrus"
)

var JddChan chan *Jd

func init() {
	JddChan = make(chan *Jd)
	go func() {
		// 打开文件，如果文件不存在则创建一个新文件
		file, err := os.OpenFile("jd.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logger.Fatal(err)
		}
		defer file.Close()

		//go func() {
		//	// 确保文件内容被同步到磁盘
		//	ticker := time.NewTicker(1 * time.Second)
		//
		//	for {
		//		select {
		//		case _ = <-ticker.C:
		//			err = file.Sync()
		//			if err != nil {
		//				logger.Fatal(err)
		//			}
		//		}
		//	}
		//}()

		var jdd Jd
		for {
			select {
			case j := <-JddChan:
				if len(j.Brand) > 0 {
					if len(jdd.Brand) > 0 {
						js, _ := json.Marshal(jdd)
						// 写入内容
						_, err = file.Write(js)
						if err != nil {
							logger.Fatal(err)
						}
						_, err = file.WriteString("\n\n")
						if err != nil {
							logger.Fatal(err)
						}
						// 确保文件内容被同步到磁盘
						err = file.Sync()
						if err != nil {
							logger.Fatal(err)
						}
					}
					j.ProductCommentSummary = jdd.ProductCommentSummary
					jdd = Jd{}
					jdd = *j
				} else {
					JtoJdd(j, &jdd)
				}
			}
		}
	}()

}

func JtoJdd(j, jdd *Jd) {
	if len(j.ProductPrice) > 0 {
		jdd.ProductPrice = j.ProductPrice
	}

	if j.ProductCommentSummary != nil {
		jdd.ProductCommentSummary = j.ProductCommentSummary
	}
}

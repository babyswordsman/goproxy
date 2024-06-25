package jd

import "fmt"

var JddChan chan *Jd

func init() {
	JddChan = make(chan *Jd)
	go func() {
		var jdd Jd
		for {
			select {
			case j := <-JddChan:
				if len(j.Brand) > 0 {
					if len(jdd.Brand) > 0 {
						fmt.Printf("-->> %+v %+v \n",
							jdd, jdd.ProductCommentSummary)
					}
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

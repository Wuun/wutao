package engine

import (
	"fmt"
	"strconv"
	"sync"
	"time"
	"wutao/conf"
	"wutao/parser"
	"wutao/serializer"
)

var TODY_SUBGCRIBE_NUM *SUBGCRIBE_NUM

type SUBGCRIBE_NUM struct {
	mux   *sync.Mutex
	Count int
}

func init() {
	TODY_SUBGCRIBE_NUM = &SUBGCRIBE_NUM{
		mux:   &sync.Mutex{},
		Count: 0,
	}
}

func ShopIDInformationScheduler() {
	var (
		done   chan interface{}
		shopId chan serializer.CustomerInformation
	)
	shopId = make(chan serializer.CustomerInformation, 30)
	done = make(chan interface{})
	startPage, err := strconv.Atoi(conf.G_Conf.StartPage)
	if err != nil {
		panic("转换初始爬取页面失败")
	}
	for i := 1; i < 9; i++ {
		url := buildURL(conf.G_Conf.Shop, strconv.Itoa(i*100+startPage), strconv.Itoa(100))
		fmt.Print(url + "\n")
		shopIDS := parser.ShopIDParser(url)
		go func() {
			for i := 0; i < len(shopIDS); i++ {
				select {
				case shopId <- shopIDS[i]:
					fmt.Println(i, "::爬取到用户shopid", shopIDS[i])
				case <-done:
					return
				}
			}
		}()

		go func() {
			for {
				select {
				case in := <-shopId:
					parser.SubscriptCustomerParser(in.ShopID)
				case <-done:
					return
				}
			}
		}()
	}
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			TODY_SUBGCRIBE_NUM.mux.Lock()
			if TODY_SUBGCRIBE_NUM.Count >= 800 {
				fmt.Println("今日任务已完成")
				close(done)
			}
		}
	}()
	<-done
	fmt.Print("今天结束于：", startPage+800, "请记住哦。")
}

func buildURL(shop string, offset string, limit string) string {
	url := "https://shopee.tw/shop/" + shop + "/followers/?offset=" + offset + "&limit=" + limit + "&offset_of_offset=0&_=1567397955988&__classic__=1"
	return url
}

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

func ShopIDInformationScheduler() {
	var (
		done   chan interface{}
		shopId chan serializer.CustomerInformation
		w      *sync.WaitGroup
	)
	shopId = make(chan serializer.CustomerInformation, 30)
	done = make(chan interface{})
	startPage, err := strconv.Atoi(conf.G_Conf.StartPage)
	if err != nil {
		panic("转换初始爬取页面失败")
	}

	w = &sync.WaitGroup{}
	w.Add(1)
	go func() {
		defer w.Done()
		for i := 1; i < 9; i++ {
			w.Add(1)
			url := buildURL(conf.G_Conf.Shop, strconv.Itoa(i*100+startPage), strconv.Itoa(100))
			fmt.Print(url + "\n")
			shopIDS := parser.ShopIDParser(url)
			go func(done chan interface{}) {
				defer w.Done()
				for i := 0; i < len(shopIDS); i++ {
					select {
					case shopId <- shopIDS[i]:
						fmt.Println(i, "::爬取到用户shopid", shopIDS[i])
					case <-done:
						return
					}
				}
			}(done)
		}
		time.Sleep(time.Second * 3)
		close(done)
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		for i := 0; i < 5; i++ {
			w.Add(1)
			go func(done chan interface{}) {
				defer w.Done()
				for {
					select {
					case in := <-shopId:
						parser.SubscriptCustomerParser(in.ShopID)
					case <-done:
						return
					}
				}
			}(done)
		}
	}()

	w.Wait()
	fmt.Println("***********************************************")
	fmt.Println("***********************************************")
	fmt.Println("***********************************************")
	fmt.Println("   ")
	fmt.Print("今天结束于：", startPage+800, "请记住哦。")
	fmt.Println("   ")
	fmt.Println("***********************************************")
	fmt.Println("***********************************************")
	fmt.Println("***********************************************")
}

func buildURL(shop string, offset string, limit string) string {
	url := "https://shopee.tw/shop/" + shop + "/followers/?offset=" + offset + "&limit=" + limit + "&offset_of_offset=0&_=1567397955988&__classic__=1"
	return url
}

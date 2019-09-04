package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
	"wutao/conf"
	"wutao/serializer"
	"wutao/spider"
)

func ShopIDParser(url string) []serializer.CustomerInformation {
	body := spider.ShopIDSpider(url)
	var result []serializer.CustomerInformation
	document, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		fmt.Print("试图解析shopID网页出现错误，正在跳过此网页")
		return nil
	}

	document.Find(".btn-follow").Each(func(i int, selection *goquery.Selection) {
		shopID, _ := selection.Attr("shopid")
		status := selection.ChildrenFiltered("span").Text()
		subResult := serializer.CustomerInformation{
			ShopID: shopID,
			Status: status,
		}
		result = append(result, subResult)
	})

	if result == nil {
		fmt.Print("这个页面没有爬取到内容")
	}
	return result
}

func SubscriptCustomerParser(id string) bool {
	url := buildCustomerSubcribeURL(id)
	req, err := http.NewRequest("POST", url, strings.NewReader("csrfmiddlewaretoken=7uYA1gpCLYbRHp6V6fh4F8xjOgyONxSq"))
	if err != nil {
		fmt.Print("请求关注用户失败,用户ID：", id)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("cookie", conf.G_Conf.Cookie)
	req.Header.Add("if-none-match-", "55b03-09fdd33ed85a18ccf15d53602abe54c4")
	req.Header.Add("origin", "https://shopee.tw")
	req.Header.Add("referer", "https://shopee.tw/shop/2390664/followers/?__classic__=1")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("请求失败")
		return false
	}

	if resp.StatusCode != 200 {
		fmt.Println("关注用户失败，shopid为：" + id + resp.Status)
		return false
	}

	fmt.Print("关注用户成功,shopid为："+id, "\n")
	return true
}

func buildCustomerSubcribeURL(id string) (url string) {
	url = `https://shopee.tw/buyer/follow/shop/` + id + "/"
	return url
}

package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

const html = `<div class="first hello" shopid="10"><div>关注1</div></div><div class="first hello" shopid="11"><div>关注2</div></div><div class="first"><div>关注3</div></div>`

func main() {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Print(err)
	}

	document.Find(".first").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("div").Text()
		fmt.Print(text + "\n")
		shopID, _ := selection.Attr("shopid")
		fmt.Print(shopID + "\n")
	})
}

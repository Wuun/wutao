package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := `https://shopee.tw/shop/2390664/followers/?offset=100&limit=20&offset_of_offset=0&_=1567397955988&__classic__=1`

	req, err := http.NewRequest("GET", url, nil)
	errPrint(err)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("cache-control", "max-age=0,no-cache")
	req.Header.Add("cookie", `_gcl_au=1.1.2105406535.1567351453; _fbp=fb.1.1567351456754.1191833235; __BWfp=c1567351496879x37f418b66; cto_lwid=a5fe580c-d1ef-4c2c-81a6-3bf803a49298; _ga=GA1.2.535834295.1567351499; _gid=GA1.2.740442407.1567351499; SPC_IA=-1; SPC_F=pVexD8Iik3rrMFUDlpV6hd3qnScoCqfe; REC_T_ID=a7719ffe-cccc-11e9-aada-b49691377852; SPC_SI=1rvy81ta49ekzbv31whxsbpxqd4tpn2h; SPC_EC="oLK/6QFs0GM9DuLam0GK/S3dg8gmkFWILGZ2duT6SfTrTe08KVh0eirf+LKQXweMsHLAXhC4CzNCYI0oebveY8HEeC9QijVUmNfO/kPUS2K0Q+W/u76Rr0Udja2CuloBdBSSSwRt6pS8BquOlMVYolkI2krHst23nL15puoWDAs="; SPC_U=172631206; SPC_SC_TK=eb02664dbecc6c197156af615b249bd4; SPC_SC_UD=172631206; csrftoken=fWPjhfOkcIIZIu2dJwJdBm62qEokKzRu; welcomePkgShown=true; SPC_RW_HYBRID_ID=18; AMP_TOKEN=%24NOT_FOUND; cto_bundle=0JDhYl94TEI5RVBIM2RFbU1hbFluaTJEQ3dzRU5qb2hlTWpEM01oOGxybEY1TU42aWpNR3ElMkIwOXk2bHpjQldZWHRYRVN6V0tvUTF3ciUyRk9ndGRjJTJGVU55dFQ3Z0tPeHBCUEdxT09oSW90Sk9mRm5yWXN4VWFTalhHVjZkSTlqSHpqZFZMY0ozRndIWWJqdmZsdzZTT2FWYUklMkJaQSUzRCUzRA; SPC_T_IV="AyBr+zeelFaXDixt1CKSiA=="; SPC_T_ID="JHIRnf/qnGIQUYYae/E3E35YXk2HMS1qXh9K542JW5qHJWKbcEDJw4t+YoT1GuBENRQlqkKVErzIzQWOU7PflV7AOmoqMxAeaIu028qxL4A="`)
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "none")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36")
	req.Header.Add("Postman-Token", "a38c607f-dd79-4a0f-b9eb-278a6540c110,ab6257cf-7e61-4050-9210-d69ba1601668")
	req.Header.Add("Host", "shopee.tw")
	req.Header.Add("Connection", "keep-alive")

	res, err := http.DefaultClient.Do(req)
	errPrint(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	errPrint(err)

	fmt.Println(res)
	fmt.Println(string(body))

}


func errPrint(err error){
	if err !=nil {
		fmt.Print(err)
		return
	}
}
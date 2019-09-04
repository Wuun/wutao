package conf

import "flag"

type Conf struct {
	Shop      string
	StartPage string
	Cookie    string
}

var G_Conf *Conf

const cookieDefault = `_gcl_au=1.1.2105406535.1567351453; _fbp=fb.1.1567351456754.1191833235; __BWfp=c1567351496879x37f418b66; cto_lwid=a5fe580c-d1ef-4c2c-81a6-3bf803a49298; _ga=GA1.2.535834295.1567351499; _gid=GA1.2.740442407.1567351499; SPC_IA=-1; SPC_F=pVexD8Iik3rrMFUDlpV6hd3qnScoCqfe; REC_T_ID=a7719ffe-cccc-11e9-aada-b49691377852; SPC_SI=1rvy81ta49ekzbv31whxsbpxqd4tpn2h; SPC_EC="oLK/6QFs0GM9DuLam0GK/S3dg8gmkFWILGZ2duT6SfTrTe08KVh0eirf+LKQXweMsHLAXhC4CzNCYI0oebveY8HEeC9QijVUmNfO/kPUS2K0Q+W/u76Rr0Udja2CuloBdBSSSwRt6pS8BquOlMVYolkI2krHst23nL15puoWDAs="; SPC_U=172631206; SPC_SC_TK=eb02664dbecc6c197156af615b249bd4; SPC_SC_UD=172631206; csrftoken=7uYA1gpCLYbRHp6V6fh4F8xjOgyONxSq; welcomePkgShown=true; SPC_RW_HYBRID_ID=71; cto_idcpy=029636fc-4b42-43f7-8f18-6bcdffee2a04; AMP_TOKEN=%24NOT_FOUND; _gali=shop-followers; _dc_gtm_UA-61915057-6=1; cto_bundle=J6RW7194TEI5RVBIM2RFbU1hbFluaTJEQ3dyJTJGZ3NJRXRVNXVvOHl3RUJDamE1VzVmSkxpMDZpTXN6N1ZoTmUzTERqck5yUVFlTUVGZXN6bVZDQjJDekZkNGtrV3J1RjVJdHd2VUZydEtOS0IwWE0zc0NWNURnaVlmNmJ3RGZLOW1TR2pIYlJVNlpyREowYlUzaFE3VEc0V3Q2QSUzRCUzRA; SPC_T_IV="irSs8OUb8OzRy15rqQ3W9w=="; SPC_T_ID="ZllFyPeen/xcjEZICGnxRhJifaD73SYR1UqbGX01E04aTLv16/KKJRvsIcm5HzwnHHgRSgKHM7MVsiCFg+TEWgjnaV2etUmc1n4EfH+ati8="`

func init() {
	var (
		shop      string
		cookie    string
		startPage string
	)
	flag.StringVar(&shop, "shop", "2390664", "请输入商店的id")
	flag.StringVar(&cookie, "cookie", cookieDefault, "请更新cookie")
	flag.StringVar(&startPage, "start-page", "3200", "要开始爬的页面")
	flag.Parse()
	G_Conf = &Conf{
		Shop:      shop,
		Cookie:    cookie,
		StartPage: startPage,
	}
}

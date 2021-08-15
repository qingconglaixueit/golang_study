package data_svr

import (
	"data_service/svr_common"
	"log"
)

type Finance struct{}

func (f Finance) SendInfo(link string) {
	hotList := GetHotList(link)
	log.Printf(hotList)

	makeData(&SMail, hotList, svr_common.XueQiuSub)

	if err := SMail.SendMail(); err != nil {
		log.Println("SMail.SendMail error: ", err)
	}

	return
}
func GetHotList(link string) string {
	//获取当前年月日

	log.Printf("Chrome visit page %s\n", link)
	htmlContent, err := Hdb.GetHttpHtmlContent(link, "#app > div.AnonymousHome_container_2te > div.AnonymousHome_home__col--rt_sQH > div.StockHotList_stock-hot__container_3fO.StockHotList_board_Yio", `document.querySelector(".StockHotList_board_Yio")`)
	if err != nil {
		log.Printf("GetHttpHtmlContent err : %v", err)
		return ""
	}

	log.Printf(htmlContent)

	info, err := Hdb.GetSpecialCycleData(htmlContent, "table tr")
	if err != nil {
		log.Printf("GetSpecialData err : %v", err)
		return ""
	}

	return info
}

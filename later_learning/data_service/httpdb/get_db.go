package httpdb

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"time"
)

type HttpDb struct{}

//获取网站上爬取的数据
func (h HttpDb) GetHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	
	defer chromedp.Cancel(chromeCtx)

	// 执行一个空task, 用提前创建Chrome实例
	chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 60*time.Second)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
	)
	if err != nil {
		log.Printf("Run err : %v\n", err)
		return "", err
	}
	//log.Println(htmlContent)

	return htmlContent, nil
}

//得到具体的数据
func (h HttpDb) GetSpecialData(htmlContent string, selector string) (string, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Println(err)
		return "", err
	}

	var str string
	dom.Find(selector).Each(func(i int, selection *goquery.Selection) {

		str = selection.Text()

	})
	return str, nil
}

//得到具体的属性数据
func (h HttpDb) GetSpecialAttrData(htmlContent string, selector string) (string, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Println(err)
		return "", err
	}

	var str string
	var ex bool
	dom.Find(selector).Each(func(i int, selection *goquery.Selection) {
		str, ex = selection.Attr("src")
		if !ex {
			log.Println("============= no exist ============== src")
		}
	})
	return str, nil
}

//得到具体的数据 有遍历
func (h HttpDb) GetSpecialCycleData(htmlContent string, selector string) (string, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Println(err)
		return "", err
	}

	var str = "<h3>雪球热股榜 - 沪深</h3><table border='2' width='500px'>"

	dom.Find(selector).Each(func(i int, selection *goquery.Selection) {
		str = fmt.Sprintf("%s<tr align='center'>", str)
		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			s := selection.Text()
			if selection.Index() != 2 {
				str = fmt.Sprintf("%s<td>%s</td>", str, s)
			} else {
				s, _ := selection.Find("i").Attr("class")
				lastIndex := strings.LastIndex(s, "  ")
				// 获取 / 后的字符串 ，这就是源文件名
				s = s[lastIndex+7:]
				//log.Println("=====",s)
				if s == "Down"{
					str = fmt.Sprintf("%s<td bgcolor='green'><b>↓</b></td>", str)
				}else{
					str = fmt.Sprintf("%s<td bgcolor='red'><b>↑</b></td>", str)
				}
			}
		})
		str = fmt.Sprintf("%s</tr>", str)
	})
	str = fmt.Sprintf("%s</table>", str)

	return str, nil
}

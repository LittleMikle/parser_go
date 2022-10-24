package main

import (
	"fmt"
	mydb "github.com/LittleMikle/parser_go/db"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	mydb.ConnectToDb()
	fmt.Println("connected to db")

	UrlParse()
}

func UrlMainParse() {
	// Request the HTML page https://www.igromania.ru/article/32249/Poigrali_v_Destroy_All_Humans!_2-Reprobed_i_delimsya_vpechatleniyami.html
	//https://www.igromania.ru/articles/
	time.Sleep(timesleep * time.Second)
	res, err := http.Get("https://www.igromania.ru/articles/")
	if err != nil {
		logrus.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {

		logrus.Fatalf("status code error: %d %s", res.StatusCode, res.Status)

	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logrus.Fatal(err)
	}

	// Find the review items Find(".aubli_data a")
	doc.Find(".aubli_data a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		Url, ok := s.Attr("href")
		if !ok {
			logrus.Info("error, articles not found")
		}
		ArticleParse(Url)
		fmt.Printf("ARTICLE URL %d: %s\n", i, Url)

	})
}

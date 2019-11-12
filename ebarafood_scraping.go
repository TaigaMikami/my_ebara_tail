package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strconv"
  //"github.com/k0kubun/pp"
)

var (
	urls []string
  menus []string
	image_urls []string
)

func ebaraFoodScraping() (urls, menus, imageUrls []string){
	ebaraUrl := "https://www.ebarafoods.com/recipe/cla_menu/49/?&limit=100"

	doc, err := goquery.NewDocument(ebaraUrl)
	if err != nil {
		fmt.Println(err)
	}

	// 詳細url
	doc.Find("ul.list-results li a").Each(func(i int, s *goquery.Selection) {
		href, _:= s.Attr("href")
		burl, _ := url.Parse(ebaraUrl)
		var fullUrl = toAbsUrl(burl, href)
		urls = append(urls, fullUrl)
	})

	// 順位:タイトル
	doc.Find("li dl").Each(func(i int, s *goquery.Selection) {
			menu := s.Find("dt").Text()
			rank := strconv.Itoa(i+1) + "位:"
			menus = append(menus, rank + menu)
	})

	// 画像URL
	doc.Find("ul.list-results > li > a > figure > img").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("src")
		burl, _ := url.Parse(ebaraUrl)
		var fullUrl= toAbsUrl(burl, href)
		imageUrls = append(imageUrls, fullUrl)
	})

	return urls, menus, imageUrls
}

func toAbsUrl(baseUrl *url.URL, webUrl string) string {
	relUrl, err := url.Parse(webUrl)
	if err != nil {
		return ""
	}
	absUrl := baseUrl.ResolveReference(relUrl)
	return absUrl.String()
}

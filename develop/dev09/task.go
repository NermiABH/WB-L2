package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

var (
	folder  = "sites"
	hrefSet = make(map[string]struct{})
)

func main() {
	args := os.Args
	if len(args) < 2 || args[0] != "wget" {
		log.Fatal("example: wget example.com example1.com ")
	}
	urls := args[1:]
	CreateFolder(folder)
	for _, link := range urls {
		link = strings.TrimRight(link, "/")
		urlStruct, err := url.ParseRequestURI(link)
		if err != nil {
			log.Fatal(err)
		}
		reg, err := regexp.Compile("https?://([a-z0-9]+[.])*" + urlStruct.Host)
		if err != nil {
			log.Fatal(err)
		}
		CreateFolder(folder + "/" + urlStruct.Host)
		col := colly.NewCollector(colly.URLFilters(reg))
		col.OnHTML("a[href]", func(el *colly.HTMLElement) {
			href := el.Request.AbsoluteURL(el.Attr("href"))
			if _, ok := hrefSet[href]; !ok {
				hrefSet[href] = struct{}{}
				col.Visit(href)
			}
		})
		col.OnResponse(func(r *colly.Response) {
			colPath := r.Request.URL.Path
			full := folder + "/" + urlStruct.Hostname() + colPath
			_, ok := hrefSet[full]
			if ok {
				return
			}
			hrefSet[full] = struct{}{}
			if path.Ext(full) == "" {
				CreateFolder(full)
			} else {
				CreateFolder(full[:strings.LastIndexByte(full, '/')])
			}
			if path.Ext(colPath) == "" {
				if full[len(full)-1] != '/' {
					full += "/"
				}
				full += "index.html"
				if _, err := os.Create(full); err != nil {
					fmt.Println(err)
				}
			}
			if err := r.Save(full); err != nil {
				log.Fatal(ok)
			}
			if err := col.Visit(urlStruct.String()); err != nil {
				log.Fatal(err)
			}
			col.Wait()
		})
	}
}

func CreateFolder(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}

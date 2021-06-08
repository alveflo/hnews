package main

import (
	"io/ioutil"
	"log"
	"encoding/xml"
	"html/template"
	"net/http"
)

type Rss2 struct {
	XMLName		xml.Name	`xml:"rss"`
	Version		string		`xml:"version,attr"`
	// Required
	Title		string		`xml:"channel>title"`
	Link		string		`xml:"channel>link"`
	Description	string		`xml:"channel>description"`
	// Optional
	PubDate		string		`xml:"channel>pubDate"`
	ItemList	[]Item		`xml:"channel>item"`
}

type Item struct {
	// Required
	Title		string		`xml:"title"`
	Link		string		`xml:"link"`
	Description	template.HTML	`xml:"description"`
	// Optional
	Content		template.HTML	`xml:"encoded"`
	PubDate		string		`xml:"pubDate"`
	Comments	string		`xml:"comments"`
}

func main() {
	resp, err := http.Get("https://news.ycombinator.com/rss")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	rss := Rss2{}

	err = xml.Unmarshal(body, &rss)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < len(rss.ItemList); i++ {
		log.Println(rss.ItemList[i].Title)
	}
}

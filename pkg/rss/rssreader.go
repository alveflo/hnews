package rss

import (
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Rss struct {
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
	Title		string			`xml:"title"`
	Link		string			`xml:"link"`
	Description	template.HTML	`xml:"description"`
	// Optional
	Content		template.HTML	`xml:"encoded"`
	PubDate		string			`xml:"pubDate"`
	Comments	string			`xml:"comments"`
}

type RssReader struct {}

func (r *RssReader) Get(url string) (Rss, error) {
	resp, err := http.Get(url)

	if err != nil {
		return Rss{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Rss{}, err
	}



	rss := Rss{}

	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return Rss{}, err
	}

	return rss, nil
}

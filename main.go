package main

import (
	"io/ioutil"
	"log"
	"encoding/xml"
	"html/template"
	"net/http"
	"runtime"
	"os/exec"
	"fmt"

	"github.com/manifoldco/promptui"
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
	Title		string			`xml:"title"`
	Link		string			`xml:"link"`
	Description	template.HTML	`xml:"description"`
	// Optional
	Content		template.HTML	`xml:"encoded"`
	PubDate		string			`xml:"pubDate"`
	Comments	string			`xml:"comments"`
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

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

	templates := &promptui.SelectTemplates{
		Label: "{{ . }}?",
		Active: "> {{ .Title | green }}",
		Inactive: "{{ .Title }}",
		Selected: "{{ .Title | green  }}",
	}

	prompt := promptui.Select{
		Label: "News",
		Items: rss.ItemList,
		Templates: templates,
		Size: 30,
	}

	index, _, err := prompt.Run()

	if err == nil {
		openbrowser(rss.ItemList[index].Link)
	}
}

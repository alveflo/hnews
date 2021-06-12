package main

import (
	"log"
	"runtime"
	"os/exec"
	"fmt"

	"github.com/alveflo/hnreader/pkg/rss"
	"github.com/manifoldco/promptui"
)

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
	reader := rss.RssReader{}
	rss, err := reader.Get("https://news.ycombinator.com/rss")
	if err != nil {
		log.Println(err)
		return
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

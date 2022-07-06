package main

import (
	"fmt"
	"golang.design/x/clipboard"
	"net/url"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		return
	}

	textByte := clipboard.Read(clipboard.FmtText)
	text := string(textByte)

	parse, err := url.Parse(text)
	if err != nil {
		fmt.Println("Not a url")
		return
	}

	text = fmt.Sprintf("[%s](%s)", parse.Host, text)
	fmt.Println(text)
	clipboard.Write(clipboard.FmtText, []byte(text))
}

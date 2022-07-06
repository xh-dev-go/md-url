package main

import (
	"fmt"
	"golang.design/x/clipboard"
	"strings"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		return
	}

	textByte := clipboard.Read(clipboard.FmtText)
	text := string(textByte)

	if strings.HasPrefix(text, "http") {
		if strings.ContainsAny(strings.ToLower(text), "https://stackoverflow.com") {
			text = fmt.Sprintf("[Stack Overflow](%s)", text)
		} else if strings.ContainsAny(strings.ToLower(text), "https://github.com") {
			text = fmt.Sprintf("[Git Hub](%s)", text)
		} else {
			text = fmt.Sprintf("[URL](%s)", text)
		}
	}

	clipboard.Write(clipboard.FmtText, []byte(text))

}

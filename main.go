package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const expectedLength = 2

var stringsToBeSkipped = []string{
	// People I already follow
	"@kas",
	"@freemor",
	"@nblade",
	"@iolfree",
	"@mdosch",
	"Barkrowler",
	// Bots
	"http://webmeup-crawler.com/",
	"http://ahrefs.com/robot/",
	"AhrefsBot",
	"Mastodon",
	"BLEXBot",
	"MastoPeek",
	"PetalBot",
	"YandexBot",
	"bingbot",
	"SeznamBot",
	"SemrushBot",
	"http://www.apple.com/go/applebot",
	"Googlebot",
}

func readUserAgent(jsonPartOfLine string) string{
		var result map[string]interface{}

		_ = json.Unmarshal([]byte(jsonPartOfLine), &result)

		request := result["request"].(map[string]interface{})
		headers := request["headers"].(map[string]interface{})
		return fmt.Sprintf("%v", headers["User-Agent"])
}

func main() {


	data, err := ioutil.ReadFile("/var/log/hjertnes.social.log")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	for _, line := range strings.Split(string(data), "\n") {
		parts := strings.Split(line, "handled request")
		if len(parts) < expectedLength {
			continue
		}

		userAgent := readUserAgent(parts[1])

		cont := false

		for _, skip := range stringsToBeSkipped {
			if strings.Contains(userAgent, skip) {
				cont = true
				break
			}
		}

		if cont {
			continue
		}

		if strings.Contains(userAgent, "http://") || strings.Contains(userAgent, "https://") {
			fmt.Println(userAgent)
		}
	}
}

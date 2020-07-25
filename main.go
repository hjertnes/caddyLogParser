package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const expectedLength = 2

var stringsToBeSkipped = []string{
	// Bots
	"Barkrowler",
	"http://webmeup-crawler.com/",
	"http://ahrefs.com/robot/",
	"AhrefsBot",
	"Mastodon",
	"MJ12bot",
	"BLEXBot",
	"MastoPeek",
	"PetalBot",
	"YandexBot",
	"bingbot",
	"SeznamBot",
	"SemrushBot",
	"http://www.apple.com/go/applebot",
	"Googlebot",
	"fediverse.network crawler",
	"Mail.RU_Bot",
	"Let's Encrypt validation server",
	"DotBot",
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

	printed := make([]string, 0)

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

		for _, skip := range printed {
			if skip == userAgent{
				cont = true
				break
			}
		}

		if cont {
			continue
		}

		if strings.Contains(userAgent, "http://") || strings.Contains(userAgent, "https://") {
			fmt.Println(userAgent)
			printed = append(printed, userAgent)
		}
	}
}

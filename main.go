package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const expectedLength = 2

func main() {
	stringsToBeSkipped := []string{
		"@kas",
		"@freemor",
		"@nblade",
		"@iolfree",
		"@mdosch",
		"Barkrowler",
		"MastoPeek",
		"PetalBot",
		"YandexBot",
		"bingbot",
		"SeznamBot",
		"SemrushBot",
		"http://www.apple.com/go/applebot",
		"Googlebot",
	}

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

		jsonPartOfLine := parts[1]

		var result map[string]interface{}

		_ = json.Unmarshal([]byte(jsonPartOfLine), &result)

		request := result["request"].(map[string]interface{})
		headers := request["headers"].(map[string]interface{})
		userAgent := fmt.Sprintf("%v", headers["User-Agent"])
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

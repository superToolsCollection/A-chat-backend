package logic

import (
	"A-chat-backend/global"
	"strings"
)

/**
* @Author: super
* @Date: 2020-09-23 08:38
* @Description:
**/

func FilterSensitive(content string) string {
	for _, word := range global.SensitiveWords {
		content = strings.ReplaceAll(content, word, "**")
	}

	return content
}

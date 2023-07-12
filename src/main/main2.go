package main

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	//// 第一步: 新建语言包
	bundle := i18n.NewBundle(language.English)
	//
	//// 第二步：使用一到多个语言标签来创建 localizer
	//loc := i18n.NewLocalizer(bundle, language.English.String())
	//
	//// 第三步：定义消息
	//messages := &i18n.Message{
	//	ID:          "Emails",
	//	Description: "The number of unread emails a user has",
	//	One:         "{{.Name}} has {{.Count}} email.",
	//	Other:       "{{.Name}} has {{.Count}} emails.",
	//}
	//
	//// 第四步：翻译消息
	//messagesCount := 2
	//translation := loc.MustLocalize(&i18n.LocalizeConfig{
	//	DefaultMessage: messages,
	//	TemplateData: map[string]interface{}{
	//		"Name":  "Theo",
	//		"Count": messagesCount,
	//	},
	//	PluralCount: messagesCount,
	//})
	//
	//fmt.Println(translation)

	// 从文件解析
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	bundle.MustLoadMessageFile("en.json")
	bundle.MustLoadMessageFile("el.json")

	loc := i18n.NewLocalizer(bundle, "en")
	messagesCount := 10

	translation := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "message",
		TemplateData: map[string]interface{}{
			"name":  "Alex",
			"Count": messagesCount,
			"Aaa":   "Bbb",
		},
		PluralCount: messagesCount,
	})

	fmt.Println(translation)

	//translation2 := loc.MustLocalize(&i18n.LocalizeConfig{
	//	MessageID: "hsms18",
	//	TemplateData: map[string]interface{}{
	//		"reason":   "Alex",
	//		"current":  "current",
	//		"boundary": "Bbb",
	//	},
	//})
	//fmt.Println(translation2)
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	// 创建一个Bundle对象，用于加载本地化文件
	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// 加载本地化文件
	bundle.LoadMessageFile("ru.json")

	// 创建一个Localizer对象，用于将本地化字符串转换为特定语言和地区的字符串
	localizer := i18n.NewLocalizer(bundle, "ru")

	//使用Localizer对象获取本地化字符串
	//pointAlarm := localizer.MustLocalize(&i18n.LocalizeConfig{
	//	DefaultMessage: &i18n.Message{
	//		ID:    "pointAlarm.告警",
	//		Other: "aa",
	//	}, MessageID: "pointAlarm.告警",
	//})

	pointAlarm := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "aaa",
	})
	fmt.Println(pointAlarm)
}

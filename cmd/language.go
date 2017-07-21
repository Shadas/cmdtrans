package cmd

import (
	"cmdtrans/config"
)

var languageDic map[string]string
var langCodeSortSeq []string

func ChooseLanguage() (sourceLanguage string, targetLanguage string, err error) {
	sourceLanguage = config.ParamsObj.SourceLang
	targetLanguage = config.ParamsObj.TargetLang
	if !checkLangCode(sourceLanguage) {
		err = ErrSourceLangCodeNotFound
		return
	}
	if !checkLangCode(targetLanguage) {
		err = ErrTargetLangCodeNotFound
		return
	}
	return
}

func checkLangCode(code string) bool {
	for k := range languageDic {
		if k == code {
			return true
		}
	}
	return false
}

func init() {
	languageDic = make(map[string]string)
	languageDic["auto"] = "自动检测"
	languageDic["zh"] = "中文"
	languageDic["en"] = "英语"
	languageDic["yue"] = "粤语"
	languageDic["wyw"] = "文言文"
	languageDic["jp"] = "日语"
	languageDic["kor"] = "韩语"
	languageDic["fra"] = "法语"
	languageDic["spa"] = "西班牙语"
	languageDic["th"] = "泰语"
	languageDic["ara"] = "阿拉伯语"
	languageDic["ru"] = "俄语"
	languageDic["pt"] = "葡萄牙语"
	languageDic["de"] = "德语"
	languageDic["it"] = "意大利语"
	languageDic["el"] = "希腊语"
	languageDic["nl"] = "荷兰语"
	languageDic["pl"] = "波兰语"
	languageDic["bul"] = "保加利亚语"
	languageDic["est"] = "爱沙尼亚语"
	languageDic["dan"] = "丹麦语"
	languageDic["fin"] = "芬兰语"
	languageDic["cs"] = "捷克语"
	languageDic["rom"] = "罗马尼亚语"
	languageDic["slo"] = "斯洛文尼亚语"
	languageDic["swe"] = "瑞典语"
	languageDic["hu"] = "匈牙利语"
	languageDic["cht"] = "繁体中文"
	languageDic["vie"] = "越南语"

	langCodeSortSeq = []string{
		"auto",
		"zh",
		"en",
		"yue",
		"wyw",
		"jp",
		"kor",
		"fra",
		"spa",
		"th",
		"ara",
		"ru",
		"pt",
		"de",
		"it",
		"el",
		"nl",
		"pl",
		"bul",
		"est",
		"dan",
		"fin",
		"cs",
		"rom",
		"slo",
		"swe",
		"hu",
		"cht",
		"vie",
	}

}

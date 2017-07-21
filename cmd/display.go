package cmd

import (
	"cmdtrans/domain"
	"fmt"
	"strings"
)

func Display(transret domain.TransResults) {
	str := ""
	for _, ret := range transret {
		str += (ret.Dst + " ")
	}
	str = strings.TrimSpace(str)
	fmt.Println(str)
}

// display the supported language codes
func ShowLanguageCodes() {
	str := ""
	for _, code := range langCodeSortSeq {
		newstr := code + " | " + languageDic[code] + "\n"
		str += newstr
	}
	fmt.Println(str)
}

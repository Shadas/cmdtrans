package cmd

import (
	"cmdtrans/config"
	"cmdtrans/domain"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const transAddr = "http://api.fanyi.baidu.com/api/trans/vip/translate"

func DealTrans(args []string) {
	str := ""
	for _, arg := range args {
		narg := arg + "\n"
		str += narg
	}
	str = strings.TrimRight(str, "\n")
	sourceLanguage, targetLanguage, err := ChooseLanguage()
	if err != nil {
		DealError(err)
		return
	}
	ret, err := Translate(str, sourceLanguage, targetLanguage)
	if err != nil {
		DealError(err)
		return
	}
	Display(ret)
	return
}

func Translate(query, from, to string) (domain.TransResults, error) {
	ret := domain.TransResults{}
	salt := getSalt()
	signature := getSignature(query, salt)
	req, err := http.NewRequest("GET", transAddr, nil)
	if err != nil {
		fmt.Println(err.Error())
		return ret, err
	}
	q := req.URL.Query()
	q.Add("q", query)
	q.Add("from", from)
	q.Add("to", to)
	q.Add("appid", config.AppId)
	q.Add("salt", fmt.Sprintf("%v", salt))
	q.Add("sign", signature)
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return ret, err
	}
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ret, err
	}
	tb := domain.TransBody{}
	if err := json.Unmarshal(buffer, &tb); err != nil {
		return ret, err
	}
	ret = tb.TransResult
	return ret, nil
}

func getSignature(query string, salt int) string {
	str := fmt.Sprintf("%v%v%v%v", config.AppId, query, salt, config.SecretKey)
	return encodeMd5(str)
}

func getSalt() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(100)
}

func encodeMd5(source string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(source))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

package config

const (
	AppId     = "" //Your app_id in baidu.
	SecretKey = "" //Your secret key.
)

var ParamsObj *Params

type Params struct {
	SourceLang string
	TargetLang string
}

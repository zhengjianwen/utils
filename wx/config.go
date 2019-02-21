package wx

import "time"

var AppID  = "wx73edafb2e05eaaff"
var Appsecret  = "57ebf04c934544dda4d0d681ca890598"

var G *Config

type Config struct {
	Debug 		bool
	Appid 		string
	AppScret 	string
	Token 		string
}

func Init()  {
	G = &Config{
		Appid:"wx73edafb2e05eaaff",
		AppScret:"57ebf04c934544dda4d0d681ca890598",
		}
	G.Token = GetToken()
}

func SyncToken()  {
	t := time.NewTicker(time.Minute * 100)
	for range t.C {
		G.Token = GetToken()
	}

}
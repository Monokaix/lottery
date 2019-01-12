package conf

type RDsConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	IsRunning bool
} 
var RDsCacheList = []RDsConfig{
	{
		Host:      "127.0.0.1",
		Port:      6379,
		User:      "",
		Pwd:       "",
		IsRunning: true,
	},
}

var RDsCache = RDsCacheList[0]
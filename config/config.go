package config

type BaseConfig struct {
	RedisAddr, RedisPassword, LogPath string
	OtherRpcAddr                      map[string]string
}

var DevelopConfig = BaseConfig{
	RedisAddr:     "127.0.0.1:6379",
	RedisPassword: "",
	LogPath:       "/data/logs/db.sync/",
	OtherRpcAddr:  map[string]string{"127.0.0.1": "7100"},
}

var TokyoConfig = BaseConfig{
	RedisAddr:     "liveme-octopus.han4cl.ng.0001.apne1.cache.amazonaws.com:6379",
	RedisPassword: "",
	LogPath:       "/data/logs/db.sync/",
	OtherRpcAddr:  map[string]string{"10.62.100.146": "7100", "10.66.104.156": "7100"},
}

var FrankfurtConfig = BaseConfig{
	RedisAddr:     "liveme-octopus.qcv1x9.ng.0001.euc1.cache.amazonaws.com:6379",
	RedisPassword: "",
	LogPath:       "/data/logs/db.sync/",
	OtherRpcAddr:  map[string]string{"10.68.100.169": "7100", "10.66.104.156": "7100"},
}

var UsEastConfig = BaseConfig{
	RedisAddr:     "liveme-octopus.ux4gcx.ng.0001.use1.cache.amazonaws.com:6379",
	RedisPassword: "",
	LogPath:       "/data/logs/db.sync/",
	OtherRpcAddr:  map[string]string{"10.68.100.169": "7100", "10.62.100.146": "7100"},
}

var EnvConfigMap = map[string]BaseConfig{
	"dev":       DevelopConfig,
	"tokyo":     TokyoConfig,
	"frankfurt": FrankfurtConfig,
	"us_east":   UsEastConfig,
}

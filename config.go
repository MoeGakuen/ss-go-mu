package main

var (
	config      = new(Config)
	cfgFilePath = "config.toml"
)

type Config struct {
	WebApi WebApiCfg `toml:"web_api"`
	Redis  RedisCfg  `toml:"redis"`
}

type WebApiCfg struct {
	Url    string `toml:"base_url"`
	Token  string `toml:"token"`
	NodeId int    `toml:"node_id"`
}

type RedisCfg struct {
	Host string `toml:"host"`
	Pass string `toml:"pass"`
	Db   int    `toml:"db"`
}

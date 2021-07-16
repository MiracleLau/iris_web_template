package config

type SiteConfig struct {
	SiteAddress string      `yaml:"site_address"`
	DbString    string      `yaml:"db_string"`
	LogPath     string      `yaml:"log_path"`
	SessionName string      `yaml:"session_name"`
	Redis       RedisConfig `yaml:"redis"`
}

type RedisConfig struct {
	Network  string `yaml:"network"`
	Addr     string `yaml:"addr"`
	Port     string `yaml:"port"`
	Prefix   string `yaml:"prefix"`
	Password string `yaml:"password"`
}

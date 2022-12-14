package config

type System struct {
	Env                 string `mapstructure:"env" json:"env" yaml:"env"`
	Addr                int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType              string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	OssType             string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`
	UseMultipoint       bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"`
	UseRedis            bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`
	IplimitCount        int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	IplimitTime         int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	UserDefaultPassword string `mapstructure:"user-default-password" json:"user-default-password" yaml:"user-default-password"`
}

package configs

import (
	"embed"
	"gopkg.in/yaml.v3"
	"os"
)

var Config embed.FS
var Configs Conf

type Conf struct {
	PHPSESSID        string `yaml:"PHPSESSID"`
	Port             string `yaml:"Port"`
	UserAgent        string `yaml:"UserAgent"`
	PageItems        int    `yaml:"PageItems"`
	ImageProxyServer string `yaml:"ImageProxyServer"`
}

func (conf *Conf) ReadConfig() {
	f, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(conf)
	if err != nil {
		panic(err)
	}
}

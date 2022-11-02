/**
* @file: config.go ==> model/config
* @package: config
* @author: jingxiu
* @since: 2022/11/2
* @desc: 解析配置文件
 */

package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type (
	Config struct {
		Gateway Gateway `yaml:"Gateway"`
		DB      DB      `yaml:"DB"`
		Rpc     Rpc     `yaml:"Rpc"`
	}
	// Gateway 网关启动地址
	Gateway struct {
		Listen string `yaml:"Listen"`
	}
	// DB  数据库链接配置
	DB struct {
		Type   string `yaml:"Type"`   // 链接类型
		Source string `yaml:"Source"` // 链接dns地址
	}
	// Rpc 链接配置
	Rpc struct {
		App struct {
			Host string `yaml:"Host"`
		} `yaml:"App"`
	}
)

var C *Config

func ConfigInit(path string) error {
	var err error
	C, err = ReadYamlConfig(path)
	return err
}
func ReadYamlConfig(path string) (*Config, error) {
	conf := &Config{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		err := yaml.NewDecoder(f).Decode(conf)
		if err != nil {
			return nil, err
		}
	}
	return conf, nil
}

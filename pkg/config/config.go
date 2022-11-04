package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../..")
)

var Config GConfig

type GConfig struct {
	Api struct {
		GinPort  int    `yaml:"apiPort"`
		ListenIP string `yaml:"listenIP"`
	}
	Rpc struct {
		ListenIP string `yaml:"listenIP"`
		RPCPort  int    `yaml:"rpcPort"`
	}
	Redis struct {
		Host   string `yaml:"host"`
		Passwd string `yaml:"passwd"`
		DB     int    `yaml:"db"`
	}
	Mysql struct {
		Mysql           string `yaml:"mysql"`
		MaxIdleConns    int    `yaml:"maxIdleConns"`
		MaxOpenConns    int    `yaml:"maxOpenConns"`
		ConnMaxLifetime int64  `yaml:"connMaxLifetime"`
	}
}

func NewGConfig() GConfig {
	return Config
}

func init() {
	bytes, err := ioutil.ReadFile(filepath.Join(Root, "config", "config.yaml"))
	if err != nil {
		panic(err.Error())
	}
	if err = yaml.Unmarshal(bytes, &Config); err != nil {
		panic(err.Error())
	}
}

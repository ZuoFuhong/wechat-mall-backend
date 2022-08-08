package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"sync/atomic"
)

var ServerConfigPath = defaultConfigPath

const (
	defaultConfigPath = "./app.yaml"
)

// serverConfigPath 获取服务启动的配置文件
func serverConfigPath() string {
	if ServerConfigPath == defaultConfigPath {
		flag.StringVar(&ServerConfigPath, "conf", defaultConfigPath, "server config path")
		flag.Parse()
	}
	return ServerConfigPath
}

type Config struct {
	Server struct {
		Name   string
		Addr   string
		Port   int
		Domain string
	}

	Db struct {
		Dsn             string `yaml:"dsn"`
		MaxIdleConns    int    `yaml:"max_idle_conns"`
		MaxOpenConns    int    `yaml:"max_open_conns"`
		ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
		LogLevel        int    `yaml:"log_level"`
		SlowThreshold   int    `yaml:"slow_threshold"`
	}

	Wxapp struct {
		Appid     string `yaml:"appid"`
		Appsecret string `yaml:"appsecret"`
	}

	Oss struct {
		AccessKeyId     string `yaml:"access_key_id"`
		AccessKeySecret string `yaml:"access_key_secret"`
		Endpoint        string `yaml:"endpoint"`
		BucketName      string `yaml:"bucket_name"`
	}
}

var globalConfig atomic.Value

func init() {
	globalConfig.Store(defaultConfig())
}

func defaultConfig() *Config {
	cfg := &Config{}
	return cfg
}

// GlobalConfig 获取全局配置对象
func GlobalConfig() *Config {
	return globalConfig.Load().(*Config)
}

// SetGlobalConfig 设置全局配置对象
func SetGlobalConfig(cfg *Config) {
	globalConfig.Store(cfg)
}

// LoadConfig 从配置文件加载配置, 并填充好默认值
func LoadConfig() (*Config, error) {
	configPath := serverConfigPath()
	cfg, err := parseConfigFromFile(configPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func parseConfigFromFile(configPath string) (*Config, error) {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	cfg := defaultConfig()
	if err := yaml.Unmarshal(buf, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

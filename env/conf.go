package env

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Conf struct {
	Mysql Mysql
	Redis Redis
	Wxapp Wxapp
}

type Mysql struct {
	Addr     string
	Username string
	Password string
}

type Redis struct {
	Addr   string
	Passwd string
	Db     int
}

type Wxapp struct {
	Appid     string
	Appsecret string
}

func LoadConf() *Conf {
	env := getMode()
	viper.SetConfigName("application")
	viper.AddConfigPath("conf/" + env)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	C := &Conf{}
	err = viper.Unmarshal(C)
	if err != nil {
		panic(env)
	}
	return C
}

func getMode() string {
	env := os.Getenv("RUN_MODE")
	if env == "" {
		env = "dev"
	}
	return env
}

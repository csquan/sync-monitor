package config

import (
	"fmt"
	"github.com/jialanli/windward"
	"github.com/spf13/viper"
	"os"
)

type stdout struct {
	Enable bool `mapstructure:"enable"`
	Level  int  `mapstructure:"level"`
}

type file struct {
	Enable bool   `mapstructure:"enable"`
	Path   string `mapstructure:"path"`
	Level  int    `mapstructure:"level"`
}

type kafka struct {
	Enable  bool     `mapstructure:"enable"`
	Level   int      `mapstructure:"level"`
	Brokers []string `mapstructure:"kafka_servers"`
	Topic   string   `mapstructure:"topic"`
}

type Config struct {
	Db struct {
		Name          string `yaml:"name"`
		Password      string `yaml:"password"`
		Ip            string `yaml:"ip"`
		Port          string `yaml:"port"`
		BTC_Database  string `yaml:"btc_database"`
		ETH_Database  string `yaml:"eth_database"`
		BSC_Database  string `yaml:"bsc_database"`
		HUI_Database  string `yaml:"hui_database"`
		TRON_Database string `yaml:"tron_database"`
	}
	Log struct {
		Stdout stdout `mapstructure:"stdout"`
		File   file   `mapstructure:"file"`
		Kafka  kafka  `mapstructure:"kafka"`
	}
	MonitorTime struct {
		BTC              int8   `mapstructure:"btc"`
		BSC              int8   `mapstructure:"bsc"`
		ETH              int8   `mapstructure:"eth"`
		HUI              int8   `mapstructure:"hui"`
		TRON             int8   `mapstructure:"tron"`
		QueryIntervalInt uint64 `mapstructure:"query_interval"`
	}
}

var Conf Config

func Readconfig(filename string) (*Config, error) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Errorf("++++++err ++++++++++: %v", err)
		return nil, fmt.Errorf("err : %v", err)
	}
	//加载配置文件
	file := path + "/" + filename
	w := windward.GetWindward()
	w.InitConf([]string{file}) //初始化自定义的配置文件

	//获取数据库连接名密码等数据
	var config Config //定义结构体【注意：这里需要有两层结构，因为w.ReadConfig读取的是data以及data中的数据】

	err = w.ReadConfig(file, &config)
	if err != nil {
		fmt.Sprintln("初始化配置文件失败")
		return nil, err
	}
	return &config, nil
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

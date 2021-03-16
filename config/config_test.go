package config

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"os"
	"strings"
	"testing"
)

// ServerConfig Web服务的配置
type ServerConfig struct {
	Port string `yaml:"Port"`
}

// GetPort 返回服务的端口
func (c *ServerConfig) GetPort() string {
	if strings.HasPrefix(c.Port, ":") {
		return c.Port
	}
	return ":" + c.Port
}

// RedisConfig Redis配置文件
type RedisConfig struct {
	Addr     string `yaml:"Addr"`
	Password string `yaml:"Password"`
	DB       int32  `yaml:"DB"`
}

// NebularConfig 图数据库配置文件
type NebularConfig struct {
	Address  string `yaml:"Address"`
	Port     int32  `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Space    string `yaml:"Space"`
}

type MySQLConfig struct {
	DSNTemplate  string `yaml:"DSNTemplate"`
	Username     string `yaml:"Username"`
	Password     string `yaml:"Password"`
	DBName       string `yaml:"DBName"`
	ConsulName   string `yaml:"ConsulName"`
	Address      string `yaml:"Address"`
	Timeout      string `yaml:"Timeout"`
	ReadTimeout  string `yaml:"ReadTimeout"`
	WriteTimeout string `yaml:"WriteTimeout"`
	MaxIdle      int    `yaml:"MaxIdle"`
	MaxOpen      int    `yaml:"MaxOpen"`
	IsMaster     bool   `yaml:"IsMaster"`
}

func (s *MySQLConfig) GenDSN() string {
	if s.Address != "" {
		DSN := fmt.Sprintf(s.DSNTemplate, s.Username, s.Password, s.Address, s.DBName, s.Timeout, s.ReadTimeout, s.WriteTimeout)
		return DSN
	}
	return ""
}

type OSSConfig struct {
	EndPoint        string `yaml:"Endpoint"`
	AccessKeyID     string `yaml:"AccessKeyID"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
}

type ESConfig struct {
	Address      []string `yaml:"Address"`
	Username     string   `yaml:"Username"`
	Password     string   `yaml:"Password"`
	MaxIdleConns int32    `yaml:"MaxIdleConns"`
	Timeout      int32    `yaml:"Timeout"`
}

type RocketMQConfig struct {
	NameServeAddress []string `yaml:"NameServe"`
	MaxRetryS        int
}

type Config struct {
	configDir    string
	ServerConfig *ServerConfig   `yaml:"Server"`
	Redis        *RedisConfig    `yaml:"Redis"`
	Mysql        []*MySQLConfig  `yaml:"Mysql"`
	Nebular      *NebularConfig  `yaml:"Nebular"`
	Oss          *OSSConfig      `yaml:"OSSConfig"`
	ES           *ESConfig       `yaml:"ESConfig"`
	MQConfig     *RocketMQConfig `yaml:"MQConfig"`
}


func TestLoadConfigs(t *testing.T) {
	AppConfig := Config{}

	LoadConfigFile("./test_conf/test1.yml", &AppConfig)
	fmt.Println(Render(AppConfig))
}

func TestLoadEnvVar(t *testing.T) {
	os.Setenv("TEST_ENV1", "1")
	os.Setenv("TEST_ENV2", "22")
	//os.Setenv("TEST_ENV3", "333")

	type EnvConf struct {
		TEST_ENV1 string
		TEST_ENV2 string
		TEST_ENV3 string
	}
	envConf := EnvConf{}

	LoadEnvVar(&envConf)
	fmt.Println(Render(envConf))

}

func Render(obj interface{}) string {
	s, err := jsoniter.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(s)
}
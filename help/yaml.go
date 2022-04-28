package help

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type mysql struct {
	User string `yaml:"user"`
	Host string `yaml:"host"`
	Password string `yaml:"password"`
	Port string `yaml:"port"`
	Name string `yaml:"name"`
}

type cache struct {
	Enable bool `yaml:"enable"`
	List []string `yaml:"list,flow"`
}

type email struct {
	ServerHost string `yaml:"serverhost"`
	ServerPort int `yaml:"serverport"`
	FromEmail string `yaml:"fromemail"`
	FromPasswd string `yaml:"frompasswd"`
}

type redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Password string `yaml:"password"`
	DB int `yaml:"db"`
}

type es struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
}

type Yaml struct {
	Mysql mysql
	Cache cache
	Email email
	Redis redis
	Es es
}

var Conf *Yaml = &Yaml{}

func InitYaml() error {

	yamlFile, err := ioutil.ReadFile("./config.yaml")

	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
		return err
	}

	err = yaml.Unmarshal(yamlFile, Conf)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return err
}

func (conf *Yaml) GetConfig(name string) interface{}{
	switch name {
	case "Mysql":
		return conf.Mysql
	case "Cache":
		return conf.Cache
	case "Email":
		return conf.Email
	case "Redis":
		return conf.Redis
	case "Es":
		return conf.Es
	default:
		return nil
	}
}
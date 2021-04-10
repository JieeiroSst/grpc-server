package config

import (
	"fmt"

	"github.com/JIeeiroSst/go-app/repositories/mysql"
	"github.com/kelseyhightower/envconfig"
)

type WebConfig struct {
	MysqlConfig mysql.Config 	`envconfig:"WEB_MYSQL"`
	URL string 					`envconfig:"WEB_URL"`
}

var Config WebConfig

func init(){
	envconfig.Process("",&Config)
	fmt.Println(Config)
}
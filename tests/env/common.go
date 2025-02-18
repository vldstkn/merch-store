package env

import (
	"fmt"
	"merch_store/internal/config"
	"strconv"
)

type Env struct {
	Dsn            string
	ApiAddress     string
	AccountAddress string
	Jwt            string
}

func NewEnv() *Env {
	conf := config.LoadConfig("../../configs", "test")
	return &Env{
		ApiAddress:     "http://" + conf.Public.Host + ":" + strconv.Itoa(conf.Public.Port) + "/api",
		Dsn:            fmt.Sprintf("port=%d host=%s user=user dbname=test password=123456 sslmode=disable", conf.Public.Database.Port, conf.Public.Database.Host),
		Jwt:            conf.Auth.Jwt,
		AccountAddress: conf.Addresses.Account,
	}
}

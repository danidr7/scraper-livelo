package config

import (
	"os"
)

type Env struct {
	MongoURL string
	LiveloPartnersURL string
}

func InitEnvs() *Env {
	return &Env {
		MongoURL: os.Getenv("MONGO_URL"),
		LiveloPartnersURL: os.Getenv("LIVELO_URL") + "/ccstore/v1/files/thirdparty/config_partners_compre_e_pontue.json",
	}
}
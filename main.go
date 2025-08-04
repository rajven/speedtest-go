package main

import (
	_ "time/tzdata"

	"github.com/librespeed/speedtest/config"
	"github.com/librespeed/speedtest/database"
	"github.com/librespeed/speedtest/results"
	"github.com/librespeed/speedtest/web"

	_ "github.com/breml/rootcerts"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf := config.Load(config.GetConfigFile())
	web.SetServerLocation(&conf)
	results.Initialize(&conf)
	database.SetDBInfo(&conf)
	log.Fatal(web.ListenAndServe(&conf))
}

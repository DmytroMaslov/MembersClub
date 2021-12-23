package app

import (
	"net/http"

	"github.com/DmytroMaslov/memberclub/src/config"
	"github.com/DmytroMaslov/memberclub/src/controllers"
	"github.com/DmytroMaslov/memberclub/src/log"
)

func StartApp() {
	conf := config.Load()
	if conf.Port == "" {
		log.Get().Error("$PORT must be set")
		panic("$PORT must be set")
	}
	port := ":" + conf.Port
	log.Get().Info("Service run on port:%v", port)
	http.HandleFunc("/addMember", controllers.AddUser)
	http.HandleFunc("/getAllMember", controllers.GetAllMember)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

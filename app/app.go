package app

import (
	"net/http"
	"os"

	"github.com/DmytroMaslov/memberclub/controllers"
	"github.com/DmytroMaslov/memberclub/log"
)

var port = os.Getenv("PORT")

func StartApp() {
	if port == "" {
		log.Get().Error("$PORT must be set")
		panic("$PORT must be set")
	}
	port = ":" + port
	log.Get().Info("Service run on port:%v", port)
	http.HandleFunc("/addMember", controllers.AddUser)
	http.HandleFunc("/getAllMember", controllers.GetAllMember)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

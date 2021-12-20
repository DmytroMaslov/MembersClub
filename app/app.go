package app

import (
	"net/http"

	"github.com/DmytroMaslov/memberclub/controllers"
)

func StartApp() {
	http.HandleFunc("/addMember", controllers.AddUser)
	http.HandleFunc("/getAllMember", controllers.GetAllMember)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

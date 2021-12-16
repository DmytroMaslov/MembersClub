package app

import (
	"memberclub/src/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/addMember", controllers.AddUser)
	http.HandleFunc("/getAllMember", controllers.GetAllMember)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

/*
endpoints:
view all members: return all members
add new member
name + email - ok 200
name + dublicate email - error mesage


*/

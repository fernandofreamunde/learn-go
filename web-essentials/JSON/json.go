package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age 	  int 	 `json:"age"`
}

//  OK this is not really working for me, it gives me empty json/struct 
// so encoding and decoding is not working as expected, will have to investigate later...
// Turns out I need to capitalize the vars in the struct for some reason
func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "\n%s %s is %d years old!\n\n", user.FirstName, user.LastName, user.Age)
	})

	http.HandleFunc("/encode", func (w http.ResponseWriter, r *http.Request) {
		peter := User {
			FirstName: "Peter",
			LastName:  "Parker",
			Age:       15,
		}

		fmt.Println(peter);

		json.NewEncoder(w).Encode(peter)
	})

	http.ListenAndServe(":81", nil)
}

package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type User struct {
	firstName string `json:"first_name"`
	lastName  string `json:"last_name"`
	age 	  int 	 `json:"age"`
}

//  OK this is not really working for me, it gives me empty json/struct 
// so encoding and decoding is not working as expected, will have to investigate later...
func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Println(user);
		fmt.Fprintf(w, "\n%s %s is %d years old!\n\n", user.firstName, user.lastName, user.age)
	})

	http.HandleFunc("/encode", func (w http.ResponseWriter, r *http.Request) {
		peter := User {
			firstName: "Peter",
			lastName:  "Parker",
			age:       15,
		}

		fmt.Println(peter);

		json.NewEncoder(w).Encode(peter)
	})

	http.ListenAndServe(":81", nil)
}

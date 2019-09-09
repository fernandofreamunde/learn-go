package main

import(
	"fmt"
	"log"
	"net/http"
	"time"
)

// Boilerplate for middlewares 
// func createNewMiddleware() Middleware {

//     // Create a new Middleware
//     middleware := func(next http.HandlerFunc) http.HandlerFunc {

//         // Define the http.HandlerFunc which is called by the server eventually
//         handler := func(w http.ResponseWriter, r *http.Request) {

//             // ... do middleware things

//             // Call the next middleware/handler in chain
//             next(w, r)
//         }

//         // Return newly created handler
//         return handler
//     }

//     // Return newly created middleware
//     return middleware
// }

type Middleware func(http.HandlerFunc) http.HandlerFunc

// logs all requests with its path and time to process
func Logging() Middleware {
	//create new  middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		//define the  handler
		return func(w http.ResponseWriter, r *http.Request) {
			// do middleware things
			start := time.Now()
			defer func() {
				log.Println(r.URL, time.Since(start))
			} ()

			//call the next middleware/handler
			f(w, r)
		}
	}
}

// now actually using the boilerplate
// ensures that url can only be requested with a specific method
// else returns a 400 Bad request
func Method(m string) Middleware {

    // Create a new Middleware
    middleware := func(next http.HandlerFunc) http.HandlerFunc {

        // Define the http.HandlerFunc which is called by the server eventually
        handler := func(w http.ResponseWriter, r *http.Request) {

			// ... do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

            // Call the next middleware/handler in chain
            next(w, r)
        }

        // Return newly created handler
        return handler
    }

    // Return newly created middleware
    return middleware
}

// applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":81", nil)
}

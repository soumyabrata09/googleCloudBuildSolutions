package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "./resources/views/index.html")
}

func main() {
	http.HandleFunc("/api/test", handlerFunc)
	http.ListenAndServe(":8000", nil)
}

//package main
//
//import (
//	"github.com/gorilla/mux"
//	"go-contacts/app"
//	"os"
//	"fmt"
//	"net/http"
//	"go-contacts/controllers"
//)
//
//func indexHandler(w http.ResponseWriter, r *http.Request){
//	fmt.Println("Hello....")
//}
//func main() {
//
//	router := mux.NewRouter()
//
//	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
//	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
//	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
//	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts
//	router.HandleFunc(path:"/api/me/test" , indexHandler)
//	router.Use(app.JwtAuthentication) //attach JWT auth middleware
//
//	//router.NotFoundHandler = app.NotFoundHandler
//
//	port := os.Getenv("APPLICATION_PORT")
//	if port == "" {
//		port = "8000" //localhost
//	}
//
//	fmt.Println(port)
//
//	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
//	if err != nil {
//		fmt.Print(err)
//	}
//}

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"./app"
// 	controllers "./controllers"
// 	"github.com/gorilla/mux" // this api is a type of web router framework for go lang
// )

// func main() {
// 	router := mux.NewRouter() //registering the new router
// 	//registering the api web routes
// 	// router.HandleFunc("goapi/user/new", controllers.createNewAccount).Methods("POST")
// 	// router.HandleFunc("goapi/user/login", controllers.letsAuthenticate).Methods("POST")
// 	// getTheControllerAccCreation := controllers.CreateAccount
// 	// getTheControllerAuthenticateAndLogin := controllers.letsAuthenticate
// 	router.HandleFunc("/api/user/new", controllers.CreateNewAccount).Methods("POST")
// 	router.HandleFunc("/api/user/login", controllers.LetsAuthenticate).Methods("POST")
// 	router.Use(app.JwtAuthentication) // making using the JwtAuthentication which initially checks  for JWT authenticatation and routes to the desired page if satisfied
// 	//getting the port from the .env file
// 	port := os.Getenv("APPLICATION_PORT")
// 	fmt.Printf(port) // for testing purposes
// 	app_err := http.ListenAndServe(":"+port, router)
// 	if app_err != nil {
// 		fmt.Print(app_err)
// 	}
// }

// //TESTING is going on
// // package main

// // import (
// // 	"fmt"
// // 	"log"
// // 	"net/http"
// // 	"os"
// // )

// // // a simple hello world program just to test the application
// // // on successful testing we will direct our concern towards main.go

// // func indexHandler(writer http.ResponseWriter, request *http.Request) {
// // 	urlPath := request.URL.Path
// // 	if urlPath != "/goapi/test" {
// // 		http.NotFound(writer, request)
// // 		return
// // 	}
// // 	//if everything goes well then we'll render the output of tha page here
// // 	//we will send a json response in future
// // 	fmt.Fprint(writer, "Hey everyone,This is hosted on goggle cloud primarily")

// // }
// // func main() {
// // 	http.HandleFunc("/goapi/test", indexHandler)
// // 	//loading the port from the .env file
// // 	port := os.Getenv("APPLICATION_PORT_TEST")
// // 	if port == "" {
// // 		// forcefully assigning the port to 8001 as our testing port
// // 		port = "8001"
// // 		log.Printf("Listening to the default port %s", port)

// // 	}

// // 	log.Printf("Listening port %s", port)
// // 	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

// //}

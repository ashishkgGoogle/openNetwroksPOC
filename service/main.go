package main

import "validator/validator"


func main() {

	validator.NewTekuriValidator("q")
}

// func main() {
// 	router := mux.NewRouter()

// 	// Define the API endpoint
// 	router.HandleFunc("/hello/{name}", sayHello).Methods("GET")

// 	// Start the server
// 	fmt.Println("Server listening on port 8000")
// 	log.Fatal(http.ListenAndServe(":8000", router))
// }

// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	// Get the name from the URL parameters
// 	params := mux.Vars(r)
// 	name := params["name"]

// 	// Create the greeting message
// 	message := fmt.Sprintf("Hello there Lord , %s!", name)

// 	// Write the message to the response
// 	fmt.Fprintln(w, message)
// }

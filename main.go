package main

import (
    "fmt"
    "net/http"
)

func handleGreet(w http.ResponseWriter, r *http.Request) {
    // Retrieve the name parameter from the request
    name := r.URL.Query().Get("name")
    
    // Greet the user
    greeting := fmt.Sprintf("Hello, %s! I hope you're having a great day :D", name)
    
    // Write the greeting as the response
    fmt.Fprintf(w, "%s", greeting)
}

func main() {
    // Serve the index.html file
    http.Handle("/", http.FileServer(http.Dir(".")))
    
    // Handle requests to the /greet endpoint
    http.HandleFunc("/greet", handleGreet)
    
    // Start the server on port 8080
    fmt.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", nil)
}

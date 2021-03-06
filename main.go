/**
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// [START container_hello_app]
package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"fmt"
)

func main() {
	// register hello function to handle all requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	// use PORT environment variable, or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))

}


func hello(w http.ResponseWriter, r *http.Request){

	log.Printf("Serving request: %s", r.URL.Path)
	inputs , _ := r.URL.Query()["input"]

	input, _ := strconv.Atoi(inputs[0])
	last_prime := 0
	i := 2
	for input != 1{
		for input % i == 0{
			last_prime = i
			input = input / i
		}
		i = i  + 1
	}


	
	fmt.Fprintf(w,"last_prime is: " + strconv.Itoa(last_prime)+"\n")

}


// [END container_hello_app]

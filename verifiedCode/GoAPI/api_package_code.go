/*
// Copyright 2018 Sendhil Vel. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

base_website.go
Date 		: 19/06/2018
Comment 	: This is seed file for creating any go website.
Version 	: 1.0.9
by Sendhil Vel
*/

package main

/*
	importing the packages
*/
import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	myawsfiles "github.com/Sendhil-Vel/Go_AWS_Files_Package"

	"github.com/subosito/gotenv"
)

/*
MyServer - this is a struct defined for server
*/
type MyServer struct {
	r *mux.Router
}

/*
Root - This is the function which is called when base url is called
*/
func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server Running")
}

/*
startServer - This function will starts the server.
				We define the routes in the function.
*/
func startServer(port int, url string) {

	/*
		Checks the port and url variables values
	*/
	if port <= 0 || len(url) == 0 {
		panic("invalid port or url")
	}

	/*
		Defines and prints the variable fullURL
	*/
	fullURL := fmt.Sprintf("%s:%d", url, port)
	fmt.Printf("starting server on %s\n", fullURL)

	/*
		Defines a Router
	*/
	rm := mux.NewRouter()

	/*
		initial splash screen
	*/

	/*
		Defines routes in the API
	*/
	/*
		url : /
	*/
	rm.HandleFunc("/", Root).Methods("GET")
	/*
		url : /uploadfile
		this route call the function in the package which actually upload the files.
	*/
	rm.HandleFunc("/uploadfile", myawsfiles.UploadFileS3).Methods("POST")

	/*
		Starts the API
		Set the port
	*/
	http.ListenAndServe(fullURL, &MyServer{rm})

}

/*
ServeHTTP - This set the parameter and configures the server parameter
*/
func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.r.ServeHTTP(rw, req)
}

func main() {
	/*
		Loads the env variables
	*/
	gotenv.Load()

	/*
		Get the port no from .env file.
		Convert string to int
		In case some error comes then process is stopped
	*/
	port, err := strconv.Atoi(os.Getenv("WEBSITE_PORT"))
	if err != nil {
		fmt.Println("port value is invalid")
		return
	}

	/*
		Get the status flag from .env file.
		Convert string to bool
		In case some error comes then setting to default value true
	*/
	statusflag, err := strconv.ParseBool(os.Getenv("statusflag"))
	if err != nil {
		fmt.Println("invalid status, setting to true")
		statusflag = true
	}

	/*
		Get all env variables
	*/
	url := os.Getenv("WEBSITE_URL")
	AWSACCESSKEY := os.Getenv("AWSACCESSKEY")
	AWSSECRETKEY := os.Getenv("AWSSECRETKEY")
	AWSREGION := os.Getenv("AWSREGION")
	FOLDERPATH := os.Getenv("FOLDERPATH")
	UPLOADBUCKET := os.Getenv("UPLOADBUCKET")
	AWSBASEENDPOINT := os.Getenv("AWSBASEENDPOINT")

	/*
		Calling the function in the package go_awsmyfiles_package.
		This function will verify the values.
		This will provide log if status flag is set to true.
		Provide information these credentials are valid and are we able to connect to AWS.
	*/
	myawsfiles.SetupPackage(statusflag, AWSACCESSKEY, AWSSECRETKEY, AWSREGION, FOLDERPATH, UPLOADBUCKET, AWSBASEENDPOINT)

	/*
		calls the function and starts the server.
	*/
	startServer(port, url)
}

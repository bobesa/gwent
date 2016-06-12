package main

import (
	"./gwent"
	"fmt"
	"net/http"
	//"log"
	"os"
)

func main() {
	fmt.Println("Gwent Server v" + gwent.Version)
	fmt.Println("Use " + gwent.PathQuit + " to shutdown the server safely")

	//Handle http
	http.HandleFunc(gwent.PathQuit, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Shutdown!")
		os.Exit(0)
	})
	http.HandleFunc(gwent.PathAPI, gwent.ProcessRestAPI)
	http.ListenAndServe(":8080", nil)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

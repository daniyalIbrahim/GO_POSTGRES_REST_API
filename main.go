package main

import (
	"fastReturns/routers"
	"log"
	"net/http"
	"os"
)

func main() {
	initilize()
	r := routers.Router()
	log.Fatal(http.ListenAndServe(":4654", r))

}
func initilize() {
	os.Setenv("User", "test")
	os.Setenv("Password", "test123")
	os.Setenv("Addr", "127.0.0.1:5432")
	os.Setenv("Database", "retourdb")
}

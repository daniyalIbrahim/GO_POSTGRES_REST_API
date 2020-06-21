package middleware

import (
	"encoding/json"
	db "fastReturns/database"
	"fastReturns/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

/**
CUSTOMER ENDPOINTS METHODS
*/
func DeleteCustomer(writer http.ResponseWriter, request *http.Request) {
	queries := mux.Vars(request)
	val, _ := queries["id"]
	intID, _ := strconv.Atoi(val)
	writer.Header().Set("Content-Type", "application/json")

	deleted := &models.Customer{
		Id: uint64(intID),
	}
	pgdb := db.Connect()
	i, _ := deleted.DeleteCustomer(pgdb)
	closeErr := pgdb.Close()
	if closeErr != nil {
		log.Printf("Error while closing database connection, Reason: %v\n", closeErr)
		os.Exit(100)
	}
	log.Printf("Connection closed successfully.\n")
	if i == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(`{"error": "not found"}`))
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(`{"success": "deleted"}`))
}

func PostCustomer(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	err := request.ParseForm()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	lon, _ := strconv.ParseFloat(request.FormValue("lon"), 64)
	lat, _ := strconv.ParseFloat(request.FormValue("lat"), 64)
	Area, _ := strconv.Atoi(request.FormValue("customerZip"))
	delivered, _ := strconv.ParseBool(request.FormValue("customerReceived"))
	newCustomer := &models.Customer{
		Name:             request.FormValue("customerName"),
		Address:          request.FormValue("customerAddress"),
		Postcode:         uint64(Area),
		Longitude:        lon,
		Latitude:         lat,
		Created_At:       time.Now(),
		Updated_At:       time.Now(),
		Order_Delievered: delivered,
	}

	_ = json.NewDecoder(request.Body).Decode(&newCustomer)
	pgdb := db.Connect()
	newCustomer.SaveAndReturnCustomer(pgdb)
	closeErr := pgdb.Close()
	if closeErr != nil {
		log.Printf("Error while closing database connection, Reason: %v\n", closeErr)
		os.Exit(100)
	}
	log.Printf("Connection closed successfully.\n")
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(`{"success": "created"}`))
}
func GetOneCustomer(writer http.ResponseWriter, request *http.Request) {
	queries := mux.Vars(request)
	val, _ := queries["id"]
	intID, _ := strconv.Atoi(val)
	writer.Header().Set("Content-Type", "application/json")
	temp := uint64(intID)
	if temp != 0 {
		getOne := &models.Customer{
			Id: uint64(intID),
		}
		pgdb := db.Connect()
		data, _ := getOne.GetCustomerById(pgdb, temp)
		if data != nil {
			b, err := json.Marshal(data)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				writer.Write([]byte(`{"error": "error marshalling data"}`))
				return
			}
			closeErr := pgdb.Close()
			if closeErr != nil {
				log.Printf("Error while closing database connection, Reason: %v\n", closeErr)
				os.Exit(100)
			}
			log.Printf("Connection closed successfully.\n")
			writer.WriteHeader(http.StatusOK)
			writer.Write(b)
			return
		}
	}
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte(`{"error": "not found"}`))
}

func UpdateCustomer(writer http.ResponseWriter, request *http.Request) {
	queries := mux.Vars(request)
	val, _ := queries["id"]
	intID, _ := strconv.Atoi(val)
	writer.Header().Set("Content-Type", "application/json")
	pgdb := db.Connect()
	lon, _ := strconv.ParseFloat(request.FormValue("lon"), 64)
	lat, _ := strconv.ParseFloat(request.FormValue("lat"), 64)
	Area, _ := strconv.Atoi(request.FormValue("customerZip"))
	Delievered, _ := strconv.ParseBool(request.FormValue("customerReceived"))
	updateCustomer := &models.Customer{
		Address:          request.FormValue("customerAddress"),
		Postcode:         uint64(Area),
		Longitude:        lon,
		Latitude:         lat,
		Updated_At:       time.Now(),
		Order_Delievered: Delievered,
	}
	temp := uint64(intID)
	updateCustomer.CostumerUpdateStatus(pgdb, temp)
	closeErr := pgdb.Close()
	if closeErr != nil {
		log.Printf("Error while closing database connection, Reason: %v\n", closeErr)
		os.Exit(100)
	}
	log.Printf("Connection closed successfully.\n")

}

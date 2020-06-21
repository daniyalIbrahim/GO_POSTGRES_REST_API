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

/*
	SUPPLIER ENDPOINT METHODS
*/
func DeleteSupplier(writer http.ResponseWriter, request *http.Request) {
	queries := mux.Vars(request)
	val, _ := queries["id"]
	intID, _ := strconv.Atoi(val)
	writer.Header().Set("Content-Type", "application/json")

	deleted := &models.Supplier{
		Id: uint64(intID),
	}
	pgdb := db.Connect()
	i, _ := deleted.DeleteSupplier(pgdb)
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

func PostSupplier(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	err := request.ParseForm()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	lon, _ := strconv.ParseFloat(request.FormValue("lon"), 64)
	lat, _ := strconv.ParseFloat(request.FormValue("lat"), 64)
	Postcode, _ := strconv.Atoi(request.FormValue("warehouseZip"))
	OrderArrived, _ := strconv.ParseBool(request.FormValue("orderarrived"))
	newSupplier := &models.Supplier{
		Name:          request.FormValue("supplierName"),
		Postcode:      uint64(Postcode),
		Longitude:     lon,
		Latitude:      lat,
		Created_At:    time.Now(),
		Updated_At:    time.Now(),
		Order_Arrived: OrderArrived,
	}
	_ = json.NewDecoder(request.Body).Decode(&newSupplier)
	pgdb := db.Connect()
	newSupplier.SaveAndReturnSuppliers(pgdb)
	closeErr := pgdb.Close()
	if closeErr != nil {
		log.Printf("Error while closing database connection, Reason: %v\n", closeErr)
		os.Exit(100)
	}
	log.Printf("Connection closed successfully.\n")
	json.NewEncoder(writer).Encode(newSupplier)
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(`{"success": "created"}`))
}

func GetAllSupplier(writer http.ResponseWriter, request *http.Request) {

}

func GetOneSupplier(writer http.ResponseWriter, request *http.Request) {
	queries := mux.Vars(request)
	val, _ := queries["id"]
	intID, _ := strconv.Atoi(val)
	writer.Header().Set("Content-Type", "application/json")
	temp := uint64(intID)
	if temp != 0 {
		getOne := &models.Supplier{
			Id: uint64(intID),
		}
		pgdb := db.Connect()
		data, _ := getOne.GetSupplierById(pgdb, temp)
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

func UpdateSupplier(writer http.ResponseWriter, request *http.Request) {
	queries := mux.Vars(request)
	val, _ := queries["id"]
	intID, _ := strconv.Atoi(val)
	writer.Header().Set("Content-Type", "application/json")
	pgdb := db.Connect()
	lon, _ := strconv.ParseFloat(request.FormValue("lon"), 64)
	lat, _ := strconv.ParseFloat(request.FormValue("lat"), 64)
	Area, _ := strconv.Atoi(request.FormValue("warehouseZip"))
	Arrived, _ := strconv.ParseBool(request.FormValue("orderArrived"))
	updateCustomer := &models.Supplier{
		Postcode:      uint64(Area),
		Longitude:     lon,
		Latitude:      lat,
		Updated_At:    time.Now(),
		Order_Arrived: Arrived,
	}
	temp := uint64(intID)
	updateCustomer.SupplierUpdateDetails(pgdb, temp)
	closeErr := pgdb.Close()
	if closeErr != nil {
		log.Printf("Error while closing database connection, Reason: %v\n", closeErr)
		os.Exit(100)
	}
	log.Printf("Connection closed successfully.\n")
}

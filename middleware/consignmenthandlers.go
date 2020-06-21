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
*	CONSIGNMENT ENDPOINT METHODS
 */
func UpdateConsignment(writer http.ResponseWriter, request *http.Request) {
	queries := mux.Vars(request)
	val, _ := queries["id"]
	intID, _ := strconv.Atoi(val)
	writer.Header().Set("Content-Type", "application/json")
	pgdb := db.Connect()
	lon, _ := strconv.ParseFloat(request.FormValue("lon"), 64)
	lat, _ := strconv.ParseFloat(request.FormValue("lat"), 64)
	Returned, _ := strconv.ParseBool(request.FormValue("isReturned"))
	updateCon := &models.Consignment{
		Id:              uint64(intID),
		Updated_At:      time.Now(),
		Final_Address:   request.FormValue("destination"),
		Current_Address: request.FormValue("currentAddress"),
		Longitude:       lon,
		Latitude:        lat,
		Is_Returned:     Returned,
	}
	temp := uint64(intID)
	_ = updateCon.ConsignmentUpdateLocation(pgdb, temp)
	closeErr := pgdb.Close()
	if closeErr != nil {
		log.Printf("Error while closing database connection, Reason: %v\n", closeErr)
		os.Exit(100)
	}
	log.Printf("Connection closed successfully.\n")
}

func DeleteConsignment(writer http.ResponseWriter, request *http.Request) {
	queries := mux.Vars(request)
	val, _ := queries["id"]
	intID, _ := strconv.Atoi(val)
	writer.Header().Set("Content-Type", "application/json")

	deleted := &models.Consignment{
		Id: uint64(intID),
	}
	pgdb := db.Connect()
	i, _ := deleted.DeleteConsignment(pgdb)
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

func PostConsignment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	err := request.ParseForm()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	lon, _ := strconv.ParseFloat(request.FormValue("lon"), 64)
	lat, _ := strconv.ParseFloat(request.FormValue("lat"), 64)
	customerid, _ := strconv.Atoi(request.FormValue("customerId"))
	supplierid, _ := strconv.Atoi(request.FormValue("supplierId"))
	Returned, _ := strconv.ParseBool(request.FormValue("isReturned"))
	newConsignment := &models.Consignment{
		Barcode:         request.FormValue("barcode"),
		Desc:            request.FormValue("description"),
		Supplier_ID:     uint64(supplierid),
		Customer_ID:     uint64(customerid),
		Final_Address:   request.FormValue("destination"),
		Current_Address: request.FormValue("currentAddress"),
		Longitude:       lon,
		Latitude:        lat,
		Created_At:      time.Now(),
		Updated_At:      time.Now(),
		Is_Returned:     Returned,
	}
	_ = json.NewDecoder(request.Body).Decode(&newConsignment)
	pgdb := db.Connect()
	newConsignment.SaveAndReturnConsignments(pgdb)
	closeErr := pgdb.Close()
	if closeErr != nil {
		log.Printf("Error while closing database connection, Reason: %v\n", closeErr)
		os.Exit(100)
	}
	log.Printf("Connection closed successfully.\n")
	json.NewEncoder(writer).Encode(newConsignment)
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(`{"success": "created"}`))
}

func AllConsignments(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	consignments := []*models.Consignment{}
	consignment := &models.Consignment{}
	pgdb := db.Connect()
	temp, err := consignment.GetAllConsignments(pgdb)
	if err != nil {
		log.Printf("Error while retrieving items, Reason: %v\n", err)
	}
	data := append(consignments, temp...)
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
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte(`{"error": "not found"}`))
}

func GetOneConsignment(writer http.ResponseWriter, request *http.Request) {
	queries := mux.Vars(request)
	val, _ := queries["id"]
	intID, _ := strconv.Atoi(val)
	writer.Header().Set("Content-Type", "application/json")
	temp := uint64(intID)

	if temp != 0 {
		getOne := &models.Consignment{
			Id: uint64(intID),
		}
		pgdb := db.Connect()
		data, _ := getOne.GetConsignmentById(pgdb, temp)
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

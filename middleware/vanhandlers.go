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
	VAN ENDPOINTS METHODS
*/
func DeleteVan(writer http.ResponseWriter, request *http.Request) {
	queries := mux.Vars(request)
	val, _ := queries["van_id"]
	intID, _ := strconv.Atoi(val)
	writer.Header().Set("Content-Type", "application/json")
	deleted := &models.Van{
		Id: uint64(intID),
	}
	pgdb := db.Connect()
	i, _ := deleted.DeleteVan(pgdb)
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

func PostVan(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	err := request.ParseForm()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	lon, _ := strconv.ParseFloat(request.FormValue("lon"), 64)
	lat, _ := strconv.ParseFloat(request.FormValue("lat"), 64)
	Areassigned, _ := strconv.Atoi(request.FormValue("zipVan"))
	supplierid, _ := strconv.Atoi(request.FormValue("supplierId"))
	delivered, _ := strconv.ParseBool(request.FormValue("orderDelivered"))
	newvan := &models.Van{
		Supplier_ID:      uint64(supplierid),
		Van_Number:       request.FormValue("vanNumber"),
		Area_Assigned:    uint64(Areassigned),
		Longitude:        lon,
		Latitude:         lat,
		CreatedAt:        time.Now(),
		Updated_At:       time.Now(),
		Parcel_Delivered: delivered,
	}
	_ = json.NewDecoder(request.Body).Decode(&newvan)
	pgdb := db.Connect()
	newvan.SaveVanRecord(pgdb)
	closeErr := pgdb.Close()
	if closeErr != nil {
		log.Printf("Error while closing database connection, Reason: %v\n", closeErr)
		os.Exit(100)
	}
	log.Printf("Connection closed successfully.\n")
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(`{"success": "created"}`))
}
func GetOneVan(writer http.ResponseWriter, request *http.Request) {
	queries := mux.Vars(request)
	val, _ := queries["id"]
	intID, _ := strconv.Atoi(val)
	writer.Header().Set("Content-Type", "application/json")
	temp := uint64(intID)
	if temp != 0 {
		getOne := &models.Van{
			Id: uint64(intID),
		}
		pgdb := db.Connect()
		data, _ := getOne.GetVanById(pgdb, temp)
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

func UpdateVan(writer http.ResponseWriter, request *http.Request) {
	queries := mux.Vars(request)
	val, _ := queries["id"]
	intID, _ := strconv.Atoi(val)
	writer.Header().Set("Content-Type", "application/json")
	pgdb := db.Connect()
	lon, _ := strconv.ParseFloat(request.FormValue("lon"), 64)
	lat, _ := strconv.ParseFloat(request.FormValue("lat"), 64)
	Areassigned, _ := strconv.Atoi(request.FormValue("zipVan"))
	Delievered, _ := strconv.ParseBool(request.FormValue("orderDelivered"))
	updateVan := &models.Van{
		Id:               uint64(intID),
		Area_Assigned:    uint64(Areassigned),
		Longitude:        lon,
		Latitude:         lat,
		Updated_At:       time.Now(),
		Parcel_Delivered: Delievered,
	}
	temp := uint64(intID)
	_ = updateVan.VanUpdateDeliveryDetails(pgdb, temp)
	closeErr := pgdb.Close()
	if closeErr != nil {
		log.Printf("Error while closing database connection, Reason: %v\n", closeErr)
		os.Exit(100)
	}
	log.Printf("Connection closed successfully.\n")

}

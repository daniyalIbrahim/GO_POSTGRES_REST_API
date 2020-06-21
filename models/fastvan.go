package models

import "C"
import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"time"
)

type Van struct {
	RefPointer       int       `sql:"-"													json:"-"`
	TableName        struct{}  `sql:"van"											json:"-"`
	Id               uint64    `sql:"van_id,pk,type:bigint"									json:"vanId"`
	Supplier_ID      uint64    `sql:"supplier_id,notnull,on_delete:CASCADE"					json:"supplierId"`
	Van_Number       string    `sql:"van_number,unique,type:text"								json:"vanNumber"`
	Area_Assigned    uint64    `sql:"area_of_van,type:bigint" 										json:"zipVan"`
	Longitude        float64   `sql:"longitude,type:double precision"							json:"lon"`
	Latitude         float64   `sql:"latitude,type:double precision"							json:"lat"`
	CreatedAt        time.Time ` pg:"default:now()"		json:"-"`
	Updated_At       time.Time `sql:"updated_at,type:timestamptz"							json:"-"`
	Parcel_Delivered bool      `sql:"parcel_delivered,type:boolean"       					json:"orderDelivered"`
}

func CreateVanTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Van{}, opts)
	if createErr != nil {
		log.Printf("Error while creating table van, Reason: %v\n", createErr)
		return createErr
	}
	if opts.IfNotExists == true {
		log.Printf("Table van was created already.\n")
		return nil
	}
	log.Printf("Table van was created sucessfully.\n")
	return nil
}

func (v *Van) SaveVanRecord(db *pg.DB) error {
	insertErr := db.Insert(v)
	if insertErr != nil {
		log.Printf("Error while inserting new item in Vans table, Reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("Van %s inserted sucessfully.\n", v.Van_Number)
	return nil
}

func (v *Van) SaveMultipleVans(db *pg.DB, items ...*Van) error {
	_, insertErr := db.Model(items).Insert()
	if insertErr != nil {
		log.Printf("Error while inserting bulk items in Van table, Reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("Bulk Van %s inserted sucessfully.\n", v.Van_Number)
	return nil
}

func (v *Van) DeleteVan(db *pg.DB) (int, error) {
	var i int = 0
	_, DeleteErr := db.Model(v).WherePK().Delete()
	if DeleteErr != nil {
		log.Printf("Error while deleting item, Reason:  %v\n", DeleteErr)
		return i, DeleteErr
	}
	log.Printf("Delete Sucessful \n")
	i += 1
	return i, nil
}
func (v *Van) GetVanById(db *pg.DB, cid uint64) (*Van, error) {
	getErr := db.Model(v).Where("van_id = ?", cid).Select()
	if getErr != nil {
		log.Printf("Error while selecting one item in van table, Reason: %v\n", getErr)
		return v, getErr
	}
	log.Printf("Get by id from van table sucessful %v\n")
	return v, nil
}

func (v *Van) VanUpdateDeliveryDetails(db *pg.DB, vid uint64) error {
	_, updateErr := db.Model(v).Set("area_of_van=?", v.Area_Assigned).Set("parcel_delivered=?", v.Parcel_Delivered).Set("longitude=?", v.Longitude).Set("latitude=?", v.Latitude).Set("updated_at=?", v.Updated_At).Where("van_id=?", vid).Update()
	if updateErr != nil {
		log.Printf("Error while updating Order Delivery Status, Reason:  %v\n", updateErr)
		return updateErr
	}
	log.Printf("Update Sucessful for Consignment ID %d\n", v.Id)
	return nil
}

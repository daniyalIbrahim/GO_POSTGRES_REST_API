package models

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"time"
)

type Supplier struct {
	RefPointer    int       `sql:"-"															json:"-"`
	TableName     struct{}  `sql:"suppliers"													json:"-"`
	Id            uint64    `sql:"supplier_id,pk,type:bigint"									json:"supplierId"`
	Name          string    `sql:"supplier_name,unique,type:text"										json:"supplierName"`
	Postcode      uint64    `sql:"warehouse_zip,type:bigint" 												json:"warehouseZip"`
	Longitude     float64   `sql:"longitude,type:double precision"							json:"lon"`
	Latitude      float64   `sql:"latitude,type:double precision"							json:"lat"`
	Created_At    time.Time `sql:"created_at,type:timestamptz" 	pg:"default:now()"				json:"-"`
	Updated_At    time.Time `sql:"updated_at,type:timestamptz"									json:"-"`
	Order_Arrived bool      `sql:"order_arrived,type:boolean"       							json:"orderarrived"`
}

func CreateSupplierTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Supplier{}, opts)
	if createErr != nil {
		log.Printf("Error while creating table Suppliers, Reason: %v\n", createErr)
		return createErr
	}
	if opts.IfNotExists == true {
		log.Printf("Table Suppliers was already created.\n")
		return nil
	}
	log.Printf("Table Suppliers was created sucessfully.\n")
	return nil
}

func (S *Supplier) SaveAndReturnSuppliers(db *pg.DB) (*Supplier, error) {
	InsertResult, insertErr := db.Model(S).Returning("*").Insert()
	if insertErr != nil {
		log.Printf("Error while inserting new item in Suppliers table, Reason: %v\n", insertErr)
		return nil, insertErr
	}
	log.Printf("Supplier %s inserted sucessfully.\n", S.Name)
	log.Printf("Number of rows affected are: %v\n", InsertResult.RowsAffected())
	return S, nil
}

func (S *Supplier) SaveMultipleSuppliers(db *pg.DB, items ...*Supplier) error {
	_, insertErr := db.Model(items).Insert()
	if insertErr != nil {
		log.Printf("Error while inserting bulk items in Supplier table, Reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("Bulk Supplier %s inserted sucessfully.\n", S.Name)
	return nil
}

func (S *Supplier) DeleteSupplier(db *pg.DB) (int, error) {
	var i int = 0
	_, DeleteErr := db.Model(S).WherePK().Delete()
	if DeleteErr != nil {
		log.Printf("Error while deleting item, Reason:  %v\n", DeleteErr)
		return i, DeleteErr
	}
	log.Printf("Delete Sucessful \n")
	i += 1
	return i, nil
}
func (s *Supplier) GetSupplierById(db *pg.DB, sid uint64) (*Supplier, error) {
	getErr := db.Model(s).Where("supplier_id = ?", sid).Select()
	if getErr != nil {
		log.Printf("Error while selecting one item from Supplier table, Reason: %v\n", getErr)
		return s, getErr
	}
	log.Printf("Get by id from supplier table sucessful %v\n", *s)
	return s, nil
}

func (s *Supplier) SupplierUpdateDetails(db *pg.DB, sid uint64) error {
	_, updateErr := db.Model(s).Set("order_arrived=?", s.Order_Arrived).Set("longitude=?", s.Longitude).Set("latitude=?", s.Latitude).Set("updated_at=?", s.Updated_At).Where("supplier_id=?", sid).Update()
	if updateErr != nil {
		log.Printf("Error while updating Supplier Details, Reason:  %v\n", updateErr)
		return updateErr
	}
	log.Printf("Update Sucessful for Supplier ID %d\n", s.Id)
	return nil
}

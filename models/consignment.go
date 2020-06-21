package models

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"time"
)

type Consignment struct {
	RefPointer      int       `sql:"-" 														json:"-"`
	TableName       struct{}  `sql:"consignments" 											json:"-"`
	Id              uint64    `sql:"consignment_id,pk,type:bigint" 						json:"consignmentId"`
	Barcode         string    `sql:"barcode,type:text" 										json:"barcode"`
	Desc            string    `sql:"description,type:text" 									json:"description"`
	Supplier_ID     uint64    `sql:"supplier_id,notnull,on_delete:CASCADE"					json:"supplierId"`
	Customer_ID     uint64    `sql:"customer_id,notnull,on_delete:CASCADE"					json:"customerId"`
	Final_Address   string    `sql:"destination_address,type:text" 							json:"destination"`
	Current_Address string    `sql:"current_address,type:text"  							json:"currentAddress"`
	Longitude       float64   `sql:"longitude,type:double precision"							json:"lon"`
	Latitude        float64   `sql:"latitude,type:double precision"							json:"lat"`
	Created_At      time.Time `sql:"created_at,type:timestamptz" pg:"default:now()"			json:"-"`
	Updated_At      time.Time `sql:"updated_at,type:timestamptz" 							json:"-"`
	Is_Returned     bool      `sql:"is_returned,type:boolean" 								json:"isReturned"`
}

func CreateConsignmentTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Consignment{}, opts)
	if createErr != nil {
		log.Printf("Error while creating table consignments, Reason: %v\n", createErr)
		return createErr
	}
	if opts.IfNotExists == true {
		log.Printf("Table consignments was already created.\n")
		return nil
	}
	log.Printf("Table consignments was created sucessfully.\n")
	return nil
}
func (c *Consignment) GetConsignmentById(db *pg.DB, cid uint64) (*Consignment, error) {
	getErr := db.Model(c).Column("consignment_id",
		"barcode", "description", "supplier_id", "customer_id", "destination_address", "current_address",
		"longitude", "latitude", "updated_at", "is_returned").Where("consignment_id = ?", cid).Select()
	if getErr != nil {
		log.Printf("Error while selecting one item in consignments table, Reason: %v\n", getErr)
		return c, getErr
	}
	log.Printf("Get by id from consignments table sucessful \n")
	return c, nil
}

func (c *Consignment) GetAllConsignments(db *pg.DB) ([]*Consignment, error) {
	var set []*Consignment
	getErr := db.Model(&set).Offset(0).Limit(5).Select()
	if getErr != nil {
		log.Printf("Error while selecting one item from consignments table, Reason: %v\n", getErr)
		return nil, getErr
	}
	log.Printf("Get consignments was sucessful: %v \n", set)
	return set, nil
}

func (c *Consignment) SaveAndReturnConsignments(db *pg.DB) (*Consignment, error) {
	InsertResult, insertErr := db.Model(c).Returning("*").Insert()
	if insertErr != nil {
		log.Printf("Error while inserting new item in consignments table, Reason: %v\n", insertErr)
		return nil, insertErr
	}
	log.Printf("Consignment %s inserted sucessfully.\n", c.Desc)
	log.Printf("Number of rows affected are: %v\n", InsertResult.RowsAffected())
	return c, nil
}

func (c *Consignment) SaveMultipleConsignments(db *pg.DB, items ...*Consignment) error {
	_, insertErr := db.Model(items).Insert()
	if insertErr != nil {
		log.Printf("Error while inserting bulk items in consignment table, Reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("Bulk Consignment %s inserted sucessfully.\n", c.Desc)
	return nil
}

func (c *Consignment) ConsignmentUpdateLocation(db *pg.DB, cid uint64) error {
	_, updateErr := db.Model(c).Set("current_address =?", c.Current_Address).Set("destination_address=?", c.Final_Address).Set("longitude=?", c.Longitude).Set("latitude=?", c.Latitude).Set("is_returned=?", c.Is_Returned).Set("updated_at=?", c.Updated_At).Where("consignment_id=?", cid).Update()
	if updateErr != nil {
		log.Printf("Error while updating Order Delivery Status, Reason:  %v\n", updateErr)
		return updateErr
	}
	log.Printf("Update Sucessful for Consignment ID %d\n", c.Id)
	return nil
}

func (C *Consignment) DeleteConsignment(db *pg.DB) (int, error) {
	var i int = 0
	_, DeleteErr := db.Model(C).WherePK().Delete()
	if DeleteErr != nil {
		log.Printf("Error while deleting item, Reason:  %v\n", DeleteErr)
		return i, DeleteErr
	}
	log.Printf("Delete Sucessful \n")
	i += 1
	return i, nil
}

package models

import "C"
import (
	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
	"log"
	"time"
)

type Customer struct {
	RefPointer       int       `sql:"-"											json:"-"`
	TableName        struct{}  `sql:"customers"									json:"-"`
	Id               uint64    `sql:"customer_id,pk,type:bigint"				json:"customerId"`
	Name             string    `sql:"customer_name,unique,type:text"						json:"customerName"`
	Address          string    `sql:"customer_address,type:text"							json:"customerAddress"`
	Postcode         uint64    `sql:"customer_zip,type:bigint"							json:"customerZip"`
	Longitude        float64   `sql:"longitude,type:double precision"							json:"lon"`
	Latitude         float64   `sql:"latitude,type:double precision"							json:"lat"`
	Created_At       time.Time `sql:"created_at,type:timestamptz"	pg:"default:now()"			json:"-"`
	Updated_At       time.Time `sql:"updated_at,type:timestamptz"				json:"-"`
	Order_Delievered bool      `sql:"order_delievered,type:boolean"				json:"customerReceived"`
}

func CreateCustomerTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Customer{}, opts)
	if createErr != nil {
		log.Printf("Error while creating table customers, Reason: %v\n", createErr)
		return createErr
	}
	if opts.IfNotExists == true {
		log.Printf("Table customers was already created.\n")
		return nil
	}
	log.Printf("Table customers was created sucessfully.\n")
	return nil
}

func (c *Customer) SaveCustomer(db *pg.DB) error {
	insertErr := db.Insert(c)
	if insertErr != nil {
		log.Printf("Error while inserting new item in customers table, Reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("Customer %s inserted sucessfully.\n", c.Name)
	return nil
}

func (c *Customer) SaveAndReturnCustomer(db *pg.DB) (*Customer, error) {
	InsertResult, insertErr := db.Model(c).Returning("*").Insert()
	if insertErr != nil {
		log.Printf("Error while inserting new item in customers table, Reason: %v\n", insertErr)
		return nil, insertErr
	}
	log.Printf("Customer %s inserted sucessfully.\n", c.Name)
	log.Printf("Number of rows affected are: %v\n", InsertResult.RowsAffected())
	return c, nil
}

func (c *Customer) SaveMultipleCustomers(db *pg.DB, items ...*Customer) error {
	_, insertErr := db.Model(items).Insert()
	if insertErr != nil {
		log.Printf("Error while inserting bulk items in customers table, Reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("Bulk Customer %s inserted sucessfully.\n", c.Name)
	return nil
}

func (c *Customer) CustomerRecievedOrder(db *pg.DB) error {
	_, updateErr := db.Model(c).Set("order_delievered = ?order_delievered").Where("id=?id").Update()
	if updateErr != nil {
		log.Printf("Error while updating Order Delivery Status, Reason:  %v\n", updateErr)
		return updateErr
	}
	log.Printf("Update Sucessful for Costumer ID %d\n", c.Id)
	return nil
}

func (c *Customer) DeleteCustomer(db *pg.DB) (int, error) {
	var i int = 0
	_, DeleteErr := db.Model(c).WherePK().Delete()
	if DeleteErr != nil {
		log.Printf("Error while deleting item, Reason:  %v\n", DeleteErr)
		return i, DeleteErr
	}
	log.Printf("Delete Sucessful \n")
	i += 1
	return i, nil
}

func (c *Customer) GetCustomerById(db *pg.DB, cid uint64) (*Customer, error) {
	getErr := db.Model(c).Where("customer_id = ?", cid).Select()
	if getErr != nil {
		log.Printf("Error while selecting one item in customer table, Reason: %v\n", getErr)
		return c, getErr
	}
	log.Printf("Get by id from customer table sucessful, Reason: %v\n", *c)
	return c, nil
}
func (c *Customer) CostumerUpdateStatus(db *pg.DB, cid uint64) error {
	_, updateErr := db.Model(c).Set("customer_address =?", c.Address).Set("customer_zip=?", c.Postcode).Set("longitude=?", c.Longitude).Set("latitude=?", c.Latitude).Set("order_delievered=?", c.Order_Delievered).Set("updated_at=?", c.Updated_At).Where("customer_id=?", cid).Update()
	if updateErr != nil {
		log.Printf("Error while updating Order Delivery Status, Reason:  %v\n", updateErr)
		return updateErr
	}
	log.Printf("Update Sucessful for Customer ID %d\n", c.Id)
	return nil
}

package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rushikeshkandekar/database"
	"github.com/rushikeshkandekar/model"
	"log"
	"net/http"
)

func GetAllCustomers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "application/json")

	db := database.ConnectDB()
	defer db.Close()

	var customers []model.Customer

	err := db.Model(&customers).Select()

	Handleerror(err)

	json.NewEncoder(writer).Encode(customers)

}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	log.Printf("creating new customer .......................")

	w.Header().Set("Content-Type", "application/json")

	//get connect
	db := database.ConnectDB()
	defer db.Close()

	customer := &model.Customer{
		Id: uuid.New().String(),
	}

	//decoding request
	_ = json.NewDecoder(r.Body).Decode(&customer)

	//inserting into database
	_, err := db.Model(customer).Insert()
	Handleerror(err)

	//returning product
	json.NewEncoder(w).Encode(customer)

	log.Println("Created successfully customer into the database............", customer.Id, customer.Firstname, customer.Lastname)
	defer log.Println("closing the connection of the postgres........")

}
func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.ConnectDB()
	defer db.Close()

	params := mux.Vars(r)
	customerId := params["id"]
	customer := &model.Customer{Id: customerId}
	if err := db.Model(customer).Where("id = ?", customer.Id).Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error in finding the customer by id")
		return
	}
	//returning the customer
	json.NewEncoder(w).Encode(customer)

	var customerList []model.Customer

	for _, customer := range customerList {
		if customer.Id == params["id"] {
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

}

func Handleerror(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package controller

import (
	"encoding/json"
	"github.com/Saurabhkanawade/golang-practice/database"
	"github.com/Saurabhkanawade/golang-practice/helper"
	_ "github.com/Saurabhkanawade/golang-practice/helper"
	"github.com/Saurabhkanawade/golang-practice/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func init() {

	//customer := model.Customer{CustomerId: uuid.New(), Firstname: "saurabh", Lastname: "Kanawade"}
	//fmt.Println(customer)
}

// methods of the rest api

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	log.Printf("creating new customer .......................")

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	//get connect
	db := database.Connect()
	defer db.Close()

	//creating product instance
	//customer := &model.Customer{
	//	CustomerID: strconv.Itoa(rand.Intn(1000)),
	//}

	customer := &model.Customer{
		CustomerId: uuid.New().String(),
	}

	//decoding request
	_ = json.NewDecoder(r.Body).Decode(&customer)

	//inserting into database
	_, err := db.Model(customer).Insert()
	helper.CheckErrorNill(err)

	//returning product
	json.NewEncoder(w).Encode(customer)

	log.Println("Created successfully customer into the database............", customer.CustomerId, customer.Firstname, customer.Lastname)
	defer log.Println("closing the connection of the postgres........")

}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	log.Printf("fetching customer from the database .......................")

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	//get connection with db
	db := database.Connect()
	defer db.Close()

	//creating the instance of customer
	var customer []model.Customer

	//get customer from database
	err := db.Model(&customer).Select()
	helper.CheckErrorNill(err)

	//returning products
	json.NewEncoder(w).Encode(customer)

	log.Println("fetched success with the customers ...........", customer)
	defer log.Println("closing the connection of the postgres........")

}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	log.Printf("fetching customer by id from the database .......................")

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	// make connection to db
	db := database.Connect()
	defer db.Close()

	//get ID from pathvariable
	params := mux.Vars(r)
	customerId := params["id"]
	log.Println("the customer id is .........", customerId)

	customer := &model.Customer{CustomerId: customerId}
	if err := db.Model(customer).WherePK().Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//returning the customer
	json.NewEncoder(w).Encode(customer)

	log.Println("fetched success with the customer by id .... where id : ", customerId)
	defer log.Println("closing the connection of the postgres........")

}

func UpdateCustomerById(w http.ResponseWriter, r *http.Request) {
	log.Printf("updating customer by id  .......................")
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	//connect to db
	db := database.Connect()
	db.Close()

	//getId
	params := mux.Vars(r)
	customerId := params["id"]

	//creating product instance
	customer := &model.Customer{
		CustomerId: customerId,
	}

	//sending payload
	_ = json.NewDecoder(r.Body).Decode(&customer)

	//updating record
	_, err := db.Model(customer).WherePK().Set("firstname=?", "lastname=?").Update()
	helper.CheckErrorNill(err)

	log.Println("update success of the customer ........")
	defer log.Println("closing the connection of the postgres........")
}

func DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
	log.Printf("Deleting the customer by the id ...........")
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	//connect to db
	db := database.Connect()
	db.Close()

	log.Println("Deleted successfully customer ............")
	defer log.Println("closing the connection of the postgres........")

}

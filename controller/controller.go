package controller

import (
	"encoding/json"
	"github.com/Saurabhkanawade/golang-practice/database"
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

	w.Header().Set("Content-Type", "application/json")

	//get connect
	db := database.Connect()
	defer db.Close()

	//creating product instance
	//customer := &model.Customer{
	//	CustomerID: strconv.Itoa(rand.Intn(1000)),
	//}

	customer := &model.Customer{
		Id: uuid.New().String(),
	}

	//decoding request
	_ = json.NewDecoder(r.Body).Decode(&customer)

	//inserting into database
	_, err := db.Model(customer).Insert()
	checkNilError(err)

	//returning product
	json.NewEncoder(w).Encode(customer)

	log.Println("Created successfully customer into the database............", customer.Id, customer.Firstname, customer.Lastname)
	defer log.Println("closing the connection of the postgres........")

}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	log.Printf("fetching customer from the database .......................")

	w.Header().Set("Content-Type", "application/json")

	//get connection with db
	db := database.Connect()
	defer db.Close()

	//creating the instance of customer
	var customer []model.Customer

	//get customer from database
	err := db.Model(&customer).Select()
	checkNilError(err)

	//returning products
	json.NewEncoder(w).Encode(customer)

	log.Println("fetched success with the customers ...........", customer)
	defer log.Println("closing the connection of the postgres........")

}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	log.Printf("fetching customer by id from the database .......................")

	w.Header().Set("Content-Type", "application/json")

	// make connection to db
	db := database.Connect()
	defer db.Close()

	//get ID from pathvariable
	params := mux.Vars(r)
	customerId := params["id"]

	log.Println("the customer id is .........", customerId)

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

	log.Println("fetched success with the customer by id .... where id : ")
	defer log.Println("closing the connection of the postgres........")

}

func UpdateCustomerById(w http.ResponseWriter, r *http.Request) {
	log.Printf("updating customer by id  .......................")
	w.Header().Set("Content-Type", "application/json")

	//connect to db
	db := database.Connect()
	defer db.Close()

	//getId
	params := mux.Vars(r)
	customerId := params["id"]

	//creating product instance
	customer := &model.Customer{
		Id: customerId,
	}

	//sending payload
	_ = json.NewDecoder(r.Body).Decode(&customer)

	//updating record
	_, err := db.Model(customer).Where("Id=?", customer.Id).Update()
	checkNilError(err)

	log.Println("updated the customer .......", customer.Firstname, customer.Lastname)

	json.NewEncoder(w).Encode(customer)

	log.Println("update success of the customer ........")
	defer log.Println("closing the connection of the postgres........")
}

func DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
	log.Printf("Deleting the customer by the id ...........")
	w.Header().Set("Content-Type", "application/json")

	//conneting the db
	db := database.Connect()
	defer db.Close()

	//getId
	params := mux.Vars(r)
	custId := params["id"]

	//instance of the customer
	var customer = model.Customer{
		Id: custId,
	}

	result, err := db.Model(&customer).WherePK().Delete()
	checkNilError(err)

	//returning the res
	json.NewEncoder(w).Encode("Deleted the customer successfully........." + custId)

	log.Println("Deleted successfully customer ............", result)
	defer log.Println("closing the connection of the postgres........")
}
func checkNilError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

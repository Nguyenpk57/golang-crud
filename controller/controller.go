package controller

import (
	"encoding/json"
	"fmt"
	"golang-crud/config"
	model "golang-crud/model"
	"log"
	"net/http"
)

func AllEmployee(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	var response model.Response
	var arrEmployee []model.Employee

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, city FROM employee")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&employee.Id, &employee.Name, &employee.City)
		if err != nil {
			log.Fatal(err.Error())
			response.Status = 404
			response.Message = "fail"
			response.Data = arrEmployee

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(response)
			return
		} else {
			arrEmployee = append(arrEmployee, employee)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrEmployee

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func InsertEmployee(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	name := r.FormValue("name")
	city := r.FormValue("city")

	_, err = db.Exec("INSERT INTO employee(name, city) VALUES(?, ?)", name, city)

	if err != nil {
		log.Print(err)
		response.Status = 404
		response.Message = "Insert data fail"
		fmt.Print("Insert data to database fail")

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
		return
	} else {
		response.Status = 200
		response.Message = "Insert data successfully"
		fmt.Print("Insert data to database")

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")
	name := r.FormValue("name")
	city := r.FormValue("city")

	if name != "" && city == "" {
		_, err = db.Exec("UPDATE employee SET name=? WHERE id=?", name, id)
	} else if city != "" && name == "" {
		_, err = db.Exec("UPDATE employee SET city=? WHERE id=?", city, id)
	} else {
		_, err = db.Exec("UPDATE employee SET name=?, city=? WHERE id=?", name, city, id)
	}

	if err != nil {
		log.Print(err)
	}

	response.Status = 200
	response.Message = "Update data successfully"
	fmt.Print("Update data successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")

	_, err = db.Exec("DELETE FROM employee WHERE id=?", id)

	if err != nil {
		log.Print(err)
		return
	}

	response.Status = 200
	response.Message = "Delete data successfully"
	fmt.Print("Delete data successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

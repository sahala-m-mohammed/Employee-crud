package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Input struct to represent the JSON input
type Input struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Designation string `json:"designation"`
}

// Input struct to represent the JSON input
type EmployeeDB struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Designation string `json:"designation"`
}

// Response struct to represent the JSON response
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    EmployeeDB
}

var employees = []EmployeeDB{
	{ID: 1, Name: "Alex", Designation: "Manager"},
	{ID: 2, Name: "John", Designation: "Developer"},
}

func main() {

	http.HandleFunc("/getEmployee", getEmployee)
	http.HandleFunc("/getEmployees", getEmployees)
	http.HandleFunc("/addEmployee", addEmployee)
	// http.HandleFunc("/deleteEmployee", deleteEmployee)
	// http.HandleFunc("/updateEmployee", updateEmployee)

	// Start the HTTP server
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
func getEmployee(w http.ResponseWriter, r *http.Request) {

	// Decode JSON input from request body
	var input Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if all required fields are present
	if input.ID == 0 || input.Name == "" || input.Designation == "" {
		// If any required field is missing, return an error response
		response := Response{
			Status:  "error",
			Message: "One or more required fields are missing",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// If all required fields are present, return success response
	response := Response{
		Status:  "success",
		Message: "All inputs are present",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	var input []Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, data := range input {
		employee := EmployeeDB{
			ID:          data.ID,
			Name:        data.Name,
			Designation: data.Designation,
		}
		employees = append(employees, employee)
	}
	response := Response{
		Status:  "success",
		Message: "Added successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// func deleteEmployee(w http.ResponseWriter, r *http.Request) {
// 	var input Input
// 	err := json.NewDecoder(r.Body).Decode(&input)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	for i, employee := range employees {
// 		if employee.ID == input.ID {
// 			employees = append(employees[:i], employees[i+1:]...)
// 			break
// 		}
// 	}
// 	response := Response{
// 		Status:  "success",
// 		Message: "Deleted successfully",

// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(response)
// }

// func updateEmployee(w http.ResponseWriter, r *http.Request) {
// 	var input Input
// 	err := json.NewDecoder(r.Body).Decode(&input)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	for i, n := range employees {
// 		if employees[i].ID == input.ID {
// 			employees[i].Name = input.Name
// 			employees[i].Designation = input.Designation

// 			if n.ID == input.ID {
// 				n.Name = input.Name
// 				n.Designation = input.Designation
// 				break
// 			}
// 		}
// 	}
// 	response := Response{
// 		Status:  "success",
// 		Message: "Updated successfully",
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(response)
// }

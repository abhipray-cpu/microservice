package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses -file
type Course struct {
	CourseId    string `json:"courseid"`
	Coursename  string `json:"coursename"`
	CoursePrice int    `json:"price"`
	// passing a refrence remember that we don't want to copy the data
	// therefor will be passing a refrence and we are defining types here
	// therefore this is a pointer type
	Author *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// creating a fake DB
var courses []Course

//middleware,helper -file

func (c Course) IsEmpty() bool {
	return c.Coursename == ""
}

//controllers -file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Dhat!! teri mayi ka chodo!!</h1>"))
}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	//setting the headers
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
	// here we are creating a new encoder which required a writer
	// and then passing in the courses to get encoded to json format
}
func getOne(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab the ID from the request
	params := mux.Vars(r)
	fmt.Println(params)

	// loop through the courses and find the matching ID

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creat one course!")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Send some data mf!!")
	}

	// what about data - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) // here we have to pass a refrence
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Send some data mf!!")
		return
	}
	//generating a uniqeID and then convert it into a string
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
	return
}
func updateOneCourse(w http.ResponseWriter, r *http.Request){
	
}
func main() {
	fmt.Println("building first API in golang")
	// saving frequentyle injects lexer which might create some issues

}

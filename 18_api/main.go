package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// models

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake database

var courses []Course

// middleware

func (c *Course) IsEmpty() bool  {
	return c.CourseName == ""
}

// controllers

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	// get is from request
	params := mux.Vars(r)

	for _, course := range courses {
		if(course.CourseId == params["id"]) {
			json.NewEncoder(w).Encode(course)
			return
		}

		json.NewEncoder(w).Encode("No course found with te given id.")
		return
	}
}

func createOneCourse(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	// check if body is empty
	if(r.Body == nil) {
		json.NewEncoder(w).Encode("No data present to add")
	}

	// if json data is sent
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if(course.IsEmpty()) {
		json.NewEncoder(w).Encode("No data inside Json")
		return
	}
	
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if(course.CourseId == params["id"]) {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if(course.CourseId == params["id"]) {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}

func main() {

	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	// routing

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":4000", r))

}
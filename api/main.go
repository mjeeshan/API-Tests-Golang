package main

import (
	"log"
	"time"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"math/rand"

)

type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice uint16 `json:"courseprice"`
	Author *Author
}

type Author struct {
	Fullname string `json:"fullname"`
	Github string `json:"github"`
}

// fake db
var courses []Course

// middleware
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main()  {
	fmt.Println("API - For Testing")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "1", CourseName: "Golang",
	CoursePrice: 1500, Author: &Author{Fullname : "Mohd Jeeshan", Github: "github.com/jeeshan12"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Cypress",
	CoursePrice: 1500, Author: &Author{Fullname : "Mohd Jeeshan", Github: "github.com/jeeshan12"}})

	r.HandleFunc("/", serveAPIHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))

}

// controllers



// serve home route
func serveAPIHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Creating API for Backend Tests</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses");
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}


func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get  Course");
	w.Header().Set("Content-Type", "application/json")
	parmas := mux.Vars(r)
	fmt.Println(parmas)
	for _, course := range courses {
		if course.CourseId == parmas["id"] {
			json.NewEncoder(w).Encode(course);
			return;
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("No Course found with given Id")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send  request body")
		return
	}

	var course Course
	_ =json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Request body is empty")
		return
	}

	for _, course := range courses {
		if course.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Can not use same course name again")
			return
		}
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a Course")
	w.Header().Set("Content-Type", "application/json")
	parmas := mux.Vars(r)
	fmt.Println(parmas)
	for index, course := range courses {
			if course.CourseId == parmas["id"] {
				courses = append(courses[:index], courses[index+1:]...)
				var course Course
				_ =json.NewDecoder(r.Body).Decode(&course)
				course.CourseId = parmas["id"]
				courses = append(courses, course)
				json.NewEncoder(w).Encode(course);
				return;
			}
		}
		json.NewEncoder(w).Encode("No Course found with given Id to udpate")
		return
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delte a  Course")
	w.Header().Set("Content-Type", "application/json")
	parmas := mux.Vars(r)
	fmt.Println(parmas)
	for index, course := range courses {
		if course.CourseId == parmas["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course is Deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given Id to delete")
	return
}
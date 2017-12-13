package main

import "net/http"

// Coordinate .
type Coordinate struct {
	Latitude  float64
	Longitude float64
}

type LoginPage struct {
	LoginStatus  bool
	LoginProblem string
}

type PrivateIndex struct {
	Username          string
	UserID            int
}

type UploadPage struct {
	UploadStatus bool
}

type Record struct {
	ID int
	Name string
	ISAnimal bool
	Kingdom string
	Phylum string
	Class string
	Order string
	Family string
	Genus string
	Species string
	Food string
	Stage string
	Season string
	Note string
	Habitat string
	Photo map[int]string // index, photo path
}

type Gallery struct {
	Records map[int]Record
}

func loadPrivateIndexData(w http.ResponseWriter, r *http.Request) PrivateIndex {
	l := PrivateIndex{}
	cu, _ := r.Cookie("username")
	l.Username = cu.Value
	return l
}
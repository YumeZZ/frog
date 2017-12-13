package main

// Coordinate .
type Coordinate struct {
	Latitude  float64
	Longitude float64
}

type LoginPage struct {
	LoginStatus  bool
	LoginProblem string
}

type UploadPage struct {
	UploadStatus bool
}

type Organism struct {
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
	Records map[int]Organism
}
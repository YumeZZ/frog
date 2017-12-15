package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func register(w http.ResponseWriter, r *http.Request) {
	logined := verifyLoginStatus(r)
	if r.Method == "POST" && logined == false {
		r.ParseForm()
		em := r.FormValue("email")
		un := r.FormValue("username")
		pw := r.FormValue("password")
		emailFormatOK := filterEmail(em)
		usernameFormatOK := filterUsername(un)
		passwordFormatOK := filterPassword(pw)
		catchFalse(emailFormatOK, "register email format err")
		catchFalse(usernameFormatOK, "register userinfo format err")
		catchFalse(passwordFormatOK, "register password format err")
		storeAccount(em, un, pw)
		success := storeSession(w, r, un, pw)
		if success == true {
			http.Redirect(w, r, "/", 303)
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	loginStatus := false
	canSkipLogin := verifyLoginStatus(r)
	if canSkipLogin == false {
		if r.Method == "POST" {
			r.ParseForm()
			loginField := r.FormValue("account")
			pw := r.FormValue("password")
			isMail := classifyLoginField(loginField)
			//fmt.Println(loginField, pw)
			if isMail == true {

			} else {
				ok := verifyPasswordByUsername(loginField, pw)
				if ok == true {
					storeSession(w, r, loginField, pw)
					loginStatus = true
				}
			}
		}
	}
	p := LoginPage{LoginStatus: loginStatus}
	b, _ := json.Marshal(p)
	w.Write(b)
}

func logout(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", 303)
}

func forgot(w http.ResponseWriter, r *http.Request) {

}

func searchRecordsByKeyword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	searchtype := r.FormValue("searchtype")
	keyword := r.FormValue("keyword")
	switch searchtype {
	case "organismname":
		records := searchRecordsByOrganismName(keyword)
		b, _ := json.Marshal(records)
		w.Write(b)
	case "category":
		records := searchRecordsByCategory(keyword)
		b, _ := json.Marshal(records)
		w.Write(b)
	case "locationname":
		records := searchRecordsByLocationName(keyword)
		b, _ := json.Marshal(records)
		w.Write(b)
	case "gps":
		longitude := r.FormValue("longitude")
		latitude := r.FormValue("latitude")
		records := searchRecordsByGPS(longitude, latitude)
		b, _ := json.Marshal(records)
		w.Write(b)
	case "season":
		records := searchRecordsBySeason(keyword)
		b, _ := json.Marshal(records)
		w.Write(b)
	case "daterange":
		datefrom := r.FormValue("datefrom")
		dateto := r.FormValue("dateto")
		records := searchRecordsByDateRange(datefrom, dateto)
		fmt.Println(records)
		b, _ := json.Marshal(records)
		w.Write(b)
	}
}

func searchRecordByRecordID(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "POST" {
			r.ParseForm()
			record := getRecordByRecordID(r.Form.Get("recordid"))
			b, _ := json.Marshal(record)
			w.Write(b)
		}
	}
}

func uploadRecord(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "POST" {
			successUpload := false
			cu, _ := r.Cookie("username")
			unFormatOK := filterUsername(cu.Value)
			if unFormatOK == true {
				UID, getUIDOK := getUIDByUsername(cu.Value)
				catchFalse(getUIDOK, "get uid by username err")
				successUpload = storeRecord(UID, r)
			}
			p := UploadPage{UploadStatus: successUpload}
			b, _ := json.Marshal(p)
			w.Write(b)
		}
	}
}

func uploadRecordPhotosByRecordID(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "POST" {
		}
	}
}

func updateRecordByRecordID(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "POST" {
			successUpdate := false
			/*
				cu, _ := r.Cookie("username")
				unFormatOK := filterUsername(cu.Value)
				if unFormatOK == true {
					UID, getUIDOK := getUIDByUsername(cu.Value)
					catchFalse(getUIDOK, "get uid by username err")
					storeDescribeOK, recordID := storeRecord(UID, r)
					photoQuantity := len(r.MultipartForm.File["photos"])
					if storeDescribeOK == true && photoQuantity == 0 {
						successUpdate = true
					} else if storeDescribeOK == true && photoQuantity != 0 {
						storePhotoOK := storeRecordPhoto(r, UID, recordID)
						if storePhotoOK == true {
							successUpdate = true
						}
					}
				}
			*/
			successUpdate = alterRecordByRecordID(r)

			p := UploadPage{UploadStatus: successUpdate}
			b, _ := json.Marshal(p)
			w.Write(b)
		}
	}
}

//舊的刪除 上傳新的 綁新的
func updateRecordPhotosByRecordID() {

}

func deleteRecordByRecordID(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "POST" {
			r.ParseForm()
			removeRecordByRecordID(r.Form.Get("recordid"))
			// need return json tall ok or not, and ajax reload
		}
	}
}

func deleteRecordPhotosByPhotoID() {

}

func parseCoordinateString(val string) float64 {
	chunks := strings.Split(val, ",")
	hours, _ := strconv.ParseFloat(strings.TrimSpace(chunks[0]), 64)
	minutes, _ := strconv.ParseFloat(strings.TrimSpace(chunks[1]), 64)
	seconds, _ := strconv.ParseFloat(strings.TrimSpace(chunks[2]), 64)
	return hours + (minutes / 60) + (seconds / 3600)
}

func parseCoordinate(latitudeValue, latitudePosition, longitudeValue, longitudePosition string) (string, string) {
	lati := parseCoordinateString(latitudeValue)
	long := parseCoordinateString(longitudeValue)

	if latitudePosition == "S" {
		lati *= -1
	}

	if longitudePosition == "W" {
		long *= -1
	}
	la := strconv.FormatFloat(lati, 'f', 6, 64)
	lo := strconv.FormatFloat(long, 'f', 6, 64)
	return la, lo
}

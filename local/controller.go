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
			if isMail == true {

			} else {
				ok := verifyPasswordByUsername(loginField, pw)
				if ok == true {
					storeSession(w, r, loginField, pw)
					loginStatus = true
				} else {

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

func uploadOrganism(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "POST" {
			successUpload := false
			cu, _ := r.Cookie("username")
			unFormatOK := filterUsername(cu.Value)
			if unFormatOK == true {
				UID, getUIDOK := getUIDByUsername(cu.Value)
				catchFalse(getUIDOK, "get uid by username err")
				storeDescribeOK, ecologyID := storeEcology(UID, r)
				photoQuantity := len(r.MultipartForm.File["photos"])
				if storeDescribeOK == true && photoQuantity == 0 {
					successUpload = true
				} else if storeDescribeOK == true && photoQuantity != 0 {
					storePhotoOK := storeEcologyPhoto(r, UID, ecologyID)
					if storePhotoOK == true {
						successUpload = true
					}
				}
			}
			p := UploadPage{UploadStatus: successUpload}
			b, _ := json.Marshal(p)
			w.Write(b)
		}
	}
}

func searchOrganism(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "POST" {
			r.ParseForm()
			organismName := r.FormValue("organismname")
			fmt.Println(organismName)
		}
	}
}

func searchCategory(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		if r.Method == "POST" {
			r.ParseForm()
			// ajax return json change category.html
			// w.write
		}
	}
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

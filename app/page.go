package main

import (
	"fmt"
	"net/http"
)

func initDynamic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		initIndex(w, r)
	} else if r.URL.Path[0:6] == "/tags/" {
		initGalleryPage(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func initIndex(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func initLoginPage(w http.ResponseWriter, r *http.Request) {
	sueecssLogin := verifyLoginStatus(r)
	if sueecssLogin == false {
		renderTemplate(w, "login.html", nil)
	} else {
		http.Redirect(w, r, "/", 303)
	}
}

func initForgotPage(w http.ResponseWriter, r *http.Request) {

}

func initRegisterPage(w http.ResponseWriter, r *http.Request) {
	logined := verifyLoginStatus(r)
	if logined == false {
		renderTemplate(w, "register.html", nil)
	} else {
		http.Redirect(w, r, "/", 303)
	}
}

func initConsolePage(w http.ResponseWriter, r *http.Request) {
}

func initUploadPage(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		renderTemplate(w, "upload.html", nil)
	} else {
		http.Redirect(w, r, "/login", 303)
	}
}

func initCategoryPage(w http.ResponseWriter, r *http.Request) {
	loginStatus := verifyLoginStatus(r)
	if loginStatus == true {
		renderTemplate(w, "category.html", nil)
	}
}

func initGalleryPage(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Path[6 : len(r.URL.Path)-1]
	//fmt.Println(keyword)

	//要以相簿為單位丟json吧 這可難辦了ＱＱ
	Gallery := searchGalleryByOrganismName(keyword)
	fmt.Println(Gallery)
	renderTemplate(w, "gallery.html", Gallery)
}

func initTestPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "test.html", nil)
}

package main

import (
	"fmt"
	"bytes"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xiam/exif"
)


var (
	randomCharacterTable = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

func newRandomString(length int) string {
	randString := ""
	var buffer bytes.Buffer
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		num := r.Intn(35)
		buffer.WriteString(randomCharacterTable[num])
	}
	randString = buffer.String()
	return randString
}

func getPasswordByUsername(un string) (string, bool) {
	pw := ""
	exist := false
	err := db.QueryRow("SELECT password FROM userinfo WHERE username = ?", un).Scan(&pw)
	checkInfo(err, "select password from userinfo err")
	if err == nil {
		exist = true
	}
	return pw, exist
}

func getUIDByUsername(un string) (string, bool) {
	UID := ""
	exist := false
	err := db.QueryRow("SELECT id FROM userinfo WHERE username = ?", un).Scan(&UID)
	checkInfo(err, "select uid from userinfo err")
	if err == nil {
		exist = true
	}
	return UID, exist
}

func searchUsernameByUsername(un string) (string, bool) {
	var u string
	exist := false
	err := db.QueryRow("SELECT username FROM userinfo WHERE username = ?", un).Scan(&u)
	checkErr(err, "can not get username")
	if err == nil {
		exist = true
	}
	return u, exist
}

func searchOrganismNameByOrganismName(organismName string) (string, bool) {
	var o string
	exist := false
	err := db.QueryRow("SELECT organismname FROM ecology WHERE organismname = ?", organismName).Scan(&o)
	checkErr(err, "can not get username")
	if err == nil {
		exist = true
	}
	return o, exist
}

func searchGalleryByOrganismName(organismName string) Gallery {
	fmt.Println(organismName)
	recordIDs, gallery := []int{}, Gallery{}
	gallery.Records = make(map[int]Record)
	idrows, queryErr := db.Query("SELECT id FROM ecology WHERE organismname=?", organismName)
	checkErr(queryErr, "query organis id from comment with mysql error")
	defer idrows.Close()
	for idrows.Next() {
		var tmp int
		scanErr := idrows.Scan(&tmp)
		checkErr(scanErr, "scan organis id from comment with mysql error")
		recordIDs = append(recordIDs, tmp)
	}

	// ecology table 有空改 record table, ecologyid 有空改 recordid
	for index, id := range recordIDs {
		name, food, stage, season, note := "", "", "", "", ""
		db.QueryRow("SELECT organismname FROM ecology WHERE id = ?", id).Scan(&name)
		db.QueryRow("SELECT food FROM ecology WHERE id = ?", id).Scan(&food)
		db.QueryRow("SELECT stage FROM ecology WHERE id = ?", id).Scan(&stage)
		db.QueryRow("SELECT season FROM ecology WHERE id = ?", id).Scan(&season)
		db.QueryRow("SELECT note FROM ecology WHERE id = ?", id).Scan(&note)

		r:=Record{
			ID: id,
			Name: name,
			Food: food,
			Stage: stage,
			Season: season,	
			Note: note}
		r.Photo = make(map[int]string)
		idrows, queryErr := db.Query("SELECT path FROM photo WHERE ecologyid = ?", id)
		checkErr(queryErr, "query photo path from comment with mysql error")
		defer idrows.Close()
		i := 0
		for idrows.Next() {
			var tmp string
			scanErr := idrows.Scan(&tmp)
			checkErr(scanErr, "scan photo path from comment with mysql error")
			r.Photo[i] = tmp
			i++
		}
		gallery.Records[index] = r
	}
	return gallery
}

func storeRecord(UID string, r *http.Request) (bool, string) {
	successStore, ecologyID := false, ""
	r.ParseMultipartForm(32 << 20)

	organismName := r.Form.Get("organismname")

	result, storeRecordErr := db.Exec("INSERT INTO ecology SET userid=?, organismname=?, createtime=CURRENT_TIMESTAMP", UID, organismName)
	checkErr(storeRecordErr, "store record err")
	if storeRecordErr == nil {
		id, getEcologyIDErr := result.LastInsertId()
		checkErr(getEcologyIDErr, "get record id err")
		ecologyID = strconv.FormatInt(id, 10)
		if getEcologyIDErr == nil {
			for key, values := range r.Form {
				if key != "organismname" {
					for _, value := range values {
						updateCommand := "UPDATE ecology SET " + key + "=?" + " WHERE id=?"
						// when value quantity == 1, can do this
						_, updateErr :=db.Exec(updateCommand, value, ecologyID)
						if updateErr != nil {
							fmt.Println(updateErr)
						}
					}
				}
			}
			successStore = true
		}
	}
	return successStore, ecologyID
}

func storeRecordPhoto(r *http.Request, UID string, ecologyID string) bool {
	successStore := false
	m := r.MultipartForm
	photos := m.File["photos"]
	for index, photo := range photos {
		source, openPhotoFileErr := photo.Open()
		defer source.Close()
		checkErr(openPhotoFileErr, "open photo file err")

		randString := newRandomString(50)
		photoExt := filepath.Ext(photo.Filename)
		photoPath := randString + photoExt

		destination, createPhotoFileErr := os.Create(frogConfig.StoragePath + "photo/" + photoPath)
		defer destination.Close()
		checkErr(createPhotoFileErr, "create photo file err")

		_, copyErr := io.Copy(destination, source)
		checkErr(copyErr, "copy photo file err")

		data, decodeErr := exif.Read(frogConfig.StoragePath + "photo/" + photoPath)
		checkWarn(decodeErr, "decode photo err")

		if copyErr == nil {
			if decodeErr != nil {
				_, storeRecordPhotoErr := db.Exec("INSERT INTO photo SET userid=?, ecologyid=?, initorder=?, path=?, name=?, createtime=CURRENT_TIMESTAMP", UID, ecologyID, index, photoPath, photo.Filename)
				checkErr(storeRecordPhotoErr, "store record photo err")
				if storeRecordPhotoErr == nil {
					successStore = true
				}
			} else {
				latitudePosition := data.Tags["North or South Latitude"]
				longitudePosition := data.Tags["East or West Longitude"]
				latitudeValue := data.Tags["Latitude"]
				longitudeValue := data.Tags["Longitude"]

				latitude := ""
				longitude := ""
				latitude, longitude = parseCoordinate(latitudeValue, latitudePosition, longitudeValue, longitudePosition)

				_, storeRecordPhotoErr := db.Exec("INSERT INTO photo SET userid=?, ecologyid=?, initorder=?, path=?, name=?, longitude=?, latitude=?, createtime=CURRENT_TIMESTAMP", UID, ecologyID, index, photoPath, photo.Filename, longitude, latitude)
				checkErr(storeRecordPhotoErr, "store record photo err")
				if storeRecordPhotoErr == nil {
					successStore = true
				}

			}
		}
	}
	return successStore
}

func getRecordByRecordID(recordID string) Record {
	r:=Record{}
	return r
}


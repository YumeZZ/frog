package main

import (
	"bytes"
	"fmt"
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

func searchRecordsByOrganismName(organismName string) Records {
	fmt.Println(organismName)
	recordIDs, records := []int{}, Records{}
	records.Records = make(map[int]Record)
	idrows, queryErr := db.Query("SELECT id FROM record WHERE organismname=?", organismName)
	checkErr(queryErr, "query organis id from comment with mysql error")
	defer idrows.Close()
	for idrows.Next() {
		var tmp int
		scanErr := idrows.Scan(&tmp)
		checkErr(scanErr, "scan organis id from comment with mysql error")
		recordIDs = append(recordIDs, tmp)
	}

	for index, id := range recordIDs {
		name, food, stage, season, note := "", "", "", "", ""
		db.QueryRow("SELECT organismname FROM record WHERE id = ?", id).Scan(&name)
		db.QueryRow("SELECT food FROM record WHERE id = ?", id).Scan(&food)
		db.QueryRow("SELECT stage FROM record WHERE id = ?", id).Scan(&stage)
		db.QueryRow("SELECT season FROM record WHERE id = ?", id).Scan(&season)
		db.QueryRow("SELECT note FROM record WHERE id = ?", id).Scan(&note)

		r := Record{
			ID:     id,
			Name:   name,
			Food:   food,
			Stage:  stage,
			Season: season,
			Note:   note}
		r.PhotoSrc = make(map[int]string)
		pathRows, queryErr := db.Query("SELECT path FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo path from comment with mysql error")
		defer pathRows.Close()
		i := 0
		for pathRows.Next() {
			var tmp string
			scanErr := pathRows.Scan(&tmp)
			checkErr(scanErr, "scan photo path from comment with mysql error")

			r.PhotoSrc[i] = tmp
			i++
		}

		r.PhotoLatitude = make(map[int]string)
		latitudeRows, queryErr := db.Query("SELECT latitude FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo latitude from comment with mysql error")
		defer latitudeRows.Close()
		j := 0
		for latitudeRows.Next() {
			var tmp string
			scanErr := latitudeRows.Scan(&tmp)
			checkErr(scanErr, "scan photo latitude from comment with mysql error")
			r.PhotoLatitude[i] = tmp
			j++
		}

		r.PhotoLongitude = make(map[int]string)
		longitudeRows, queryErr := db.Query("SELECT longitude FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo longitude from comment with mysql error")
		defer longitudeRows.Close()
		k := 0
		for longitudeRows.Next() {
			var tmp string
			scanErr := longitudeRows.Scan(&tmp)
			checkErr(scanErr, "scan photo longitude from comment with mysql error")
			r.PhotoLongitude[i] = tmp
			k++
		}
		records.Records[index] = r
	}
	return records
}

func searchRecordsByCategory(category string) Records {
	fmt.Println(category)
	recordIDs, records := []int{}, Records{}
	records.Records = make(map[int]Record)
	englishCategory := false
	englishCategory = isEnglish(category)
	if englishCategory == true {
		idrows, queryErr := db.Query("SELECT id FROM record WHERE categorytagenglish=?", category)
		checkErr(queryErr, "query organis id from comment with mysql error")
		defer idrows.Close()
		for idrows.Next() {
			var tmp int
			scanErr := idrows.Scan(&tmp)
			checkErr(scanErr, "scan organis id from comment with mysql error")
			recordIDs = append(recordIDs, tmp)
		}

		for index, id := range recordIDs {
			name, food, stage, season, note := "", "", "", "", ""
			db.QueryRow("SELECT organismname FROM record WHERE id = ?", id).Scan(&name)
			db.QueryRow("SELECT food FROM record WHERE id = ?", id).Scan(&food)
			db.QueryRow("SELECT stage FROM record WHERE id = ?", id).Scan(&stage)
			db.QueryRow("SELECT season FROM record WHERE id = ?", id).Scan(&season)
			db.QueryRow("SELECT note FROM record WHERE id = ?", id).Scan(&note)

			r := Record{
				ID:     id,
				Name:   name,
				Food:   food,
				Stage:  stage,
				Season: season,
				Note:   note}
			r.PhotoSrc = make(map[int]string)
			idrows, queryErr := db.Query("SELECT path FROM photo WHERE recordid = ?", id)
			checkErr(queryErr, "query photo path from comment with mysql error")
			defer idrows.Close()
			i := 0
			for idrows.Next() {
				var tmp string
				scanErr := idrows.Scan(&tmp)
				checkErr(scanErr, "scan photo path from comment with mysql error")
				r.PhotoSrc[i] = tmp
				i++
			}
			records.Records[index] = r
		}
	} else {
		idrows, queryErr := db.Query("SELECT id FROM record WHERE categorytagchinese=?", category)
		checkErr(queryErr, "query organis id from comment with mysql error")
		defer idrows.Close()
		for idrows.Next() {
			var tmp int
			scanErr := idrows.Scan(&tmp)
			checkErr(scanErr, "scan organis id from comment with mysql error")
			recordIDs = append(recordIDs, tmp)
		}

		for index, id := range recordIDs {
			name, food, stage, season, note := "", "", "", "", ""
			db.QueryRow("SELECT organismname FROM record WHERE id = ?", id).Scan(&name)
			db.QueryRow("SELECT food FROM record WHERE id = ?", id).Scan(&food)
			db.QueryRow("SELECT stage FROM record WHERE id = ?", id).Scan(&stage)
			db.QueryRow("SELECT season FROM record WHERE id = ?", id).Scan(&season)
			db.QueryRow("SELECT note FROM record WHERE id = ?", id).Scan(&note)

			r := Record{
				ID:     id,
				Name:   name,
				Food:   food,
				Stage:  stage,
				Season: season,
				Note:   note}
			r.PhotoSrc = make(map[int]string)
			idrows, queryErr := db.Query("SELECT path FROM photo WHERE recordid = ?", id)
			checkErr(queryErr, "query photo path from comment with mysql error")
			defer idrows.Close()
			i := 0
			for idrows.Next() {
				var tmp string
				scanErr := idrows.Scan(&tmp)
				checkErr(scanErr, "scan photo path from comment with mysql error")
				r.PhotoSrc[i] = tmp
				i++
			}
			records.Records[index] = r
		}
	}
	return records
}

func searchRecordsByLocationName(locationName string) Records {
	fmt.Println(locationName)
	recordIDs, records := []int{}, Records{}
	records.Records = make(map[int]Record)
	idrows, queryErr := db.Query("SELECT id FROM record WHERE locationname=?", locationName)
	checkErr(queryErr, "query organis id from comment with mysql error")
	defer idrows.Close()
	for idrows.Next() {
		var tmp int
		scanErr := idrows.Scan(&tmp)
		checkErr(scanErr, "scan organis id from comment with mysql error")
		recordIDs = append(recordIDs, tmp)
	}

	for index, id := range recordIDs {
		name, food, stage, season, note := "", "", "", "", ""
		db.QueryRow("SELECT organismname FROM record WHERE id = ?", id).Scan(&name)
		db.QueryRow("SELECT food FROM record WHERE id = ?", id).Scan(&food)
		db.QueryRow("SELECT stage FROM record WHERE id = ?", id).Scan(&stage)
		db.QueryRow("SELECT season FROM record WHERE id = ?", id).Scan(&season)
		db.QueryRow("SELECT note FROM record WHERE id = ?", id).Scan(&note)

		r := Record{
			ID:     id,
			Name:   name,
			Food:   food,
			Stage:  stage,
			Season: season,
			Note:   note}
		r.PhotoSrc = make(map[int]string)
		idrows, queryErr := db.Query("SELECT path FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo path from comment with mysql error")
		defer idrows.Close()
		i := 0
		for idrows.Next() {
			var tmp string
			scanErr := idrows.Scan(&tmp)
			checkErr(scanErr, "scan photo path from comment with mysql error")
			r.PhotoSrc[i] = tmp
			i++
		}
		records.Records[index] = r
	}
	return records
}

func searchRecordsByGPS(longitude, latitude string) Records {
	records := Records{}
	return records
}

func searchRecordsBySeason(season string) Records {
	fmt.Println(season)
	recordIDs, records := []int{}, Records{}
	records.Records = make(map[int]Record)
	idrows, queryErr := db.Query("SELECT id FROM record WHERE season=?", season)
	checkErr(queryErr, "query organis id from comment with mysql error")
	defer idrows.Close()
	for idrows.Next() {
		var tmp int
		scanErr := idrows.Scan(&tmp)
		checkErr(scanErr, "scan organis id from comment with mysql error")
		recordIDs = append(recordIDs, tmp)
	}

	for index, id := range recordIDs {
		name, food, stage, season, note := "", "", "", "", ""
		db.QueryRow("SELECT organismname FROM record WHERE id = ?", id).Scan(&name)
		db.QueryRow("SELECT food FROM record WHERE id = ?", id).Scan(&food)
		db.QueryRow("SELECT stage FROM record WHERE id = ?", id).Scan(&stage)
		db.QueryRow("SELECT season FROM record WHERE id = ?", id).Scan(&season)
		db.QueryRow("SELECT note FROM record WHERE id = ?", id).Scan(&note)

		r := Record{
			ID:     id,
			Name:   name,
			Food:   food,
			Stage:  stage,
			Season: season,
			Note:   note}
		r.PhotoSrc = make(map[int]string)
		idrows, queryErr := db.Query("SELECT path FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo path from comment with mysql error")
		defer idrows.Close()
		i := 0
		for idrows.Next() {
			var tmp string
			scanErr := idrows.Scan(&tmp)
			checkErr(scanErr, "scan photo path from comment with mysql error")
			r.PhotoSrc[i] = tmp
			i++
		}
		records.Records[index] = r
	}
	return records
}

func searchRecordsByDateRange(dateFrom, dateTo string) Records {
	fmt.Println(dateFrom, dateTo)
	recordIDs, records := []int{}, Records{}
	records.Records = make(map[int]Record)

	idrows, queryErr := db.Query("SELECT id FROM record WHERE recorddate BETWEEN " + dateFrom + " AND " + dateTo + "")
	checkErr(queryErr, "query organis id from comment with mysql error")
	defer idrows.Close()
	for idrows.Next() {
		var tmp int
		scanErr := idrows.Scan(&tmp)
		checkErr(scanErr, "scan organis id from comment with mysql error")
		recordIDs = append(recordIDs, tmp)
	}

	for index, id := range recordIDs {
		name, food, stage, season, note := "", "", "", "", ""
		db.QueryRow("SELECT organismname FROM record WHERE id = ?", id).Scan(&name)
		db.QueryRow("SELECT food FROM record WHERE id = ?", id).Scan(&food)
		db.QueryRow("SELECT stage FROM record WHERE id = ?", id).Scan(&stage)
		db.QueryRow("SELECT season FROM record WHERE id = ?", id).Scan(&season)
		db.QueryRow("SELECT note FROM record WHERE id = ?", id).Scan(&note)

		r := Record{
			ID:     id,
			Name:   name,
			Food:   food,
			Stage:  stage,
			Season: season,
			Note:   note}
		r.PhotoSrc = make(map[int]string)
		idrows, queryErr := db.Query("SELECT path FROM photo WHERE recordid = ?", id)
		checkErr(queryErr, "query photo path from comment with mysql error")
		defer idrows.Close()
		i := 0
		for idrows.Next() {
			var tmp string
			scanErr := idrows.Scan(&tmp)
			checkErr(scanErr, "scan photo path from comment with mysql error")
			r.PhotoSrc[i] = tmp
			i++
		}
		records.Records[index] = r
	}
	return records
}

func checkOrganismNameExistByOrganismName(organismName string) bool {
	var o string
	exist := false
	err := db.QueryRow("SELECT organismname FROM record WHERE organismname = ?", organismName).Scan(&o)
	checkErr(err, "can not get username")
	if err == nil && o != "" {
		exist = true
	}
	return exist
}

func storeRecord(UID string, r *http.Request) (bool, string) {
	successStore, recordID := false, ""
	r.ParseMultipartForm(32 << 20)

	organismName := r.Form.Get("organismname")
	result, storeRecordErr := db.Exec("INSERT INTO record SET userid=?, organismname=?, createtime=CURRENT_TIMESTAMP", UID, organismName)
	checkErr(storeRecordErr, "store record err")
	if storeRecordErr == nil {
		id, getRecordIDErr := result.LastInsertId()
		checkErr(getRecordIDErr, "get record id err")
		recordID = strconv.FormatInt(id, 10)
		if getRecordIDErr == nil {
			for key, values := range r.Form {
				if key != "organismname" {
					for _, value := range values {
						//need fix
						updateCommand := "UPDATE record SET `" + key + "`=?" + " WHERE id=?"
						fmt.Println(updateCommand)
						// when value quantity == 1, can do this
						_, updateErr := db.Exec(updateCommand, value, recordID)
						if updateErr != nil {
							fmt.Println(updateErr)
						}
					}
				}
			}
			successStore = true
		}
	}
	return successStore, recordID
}

func storeRecordPhoto(r *http.Request, UID string, recordID string) bool {
	fmt.Println("storeRecordPhoto:")
	successStore := false
	m := r.MultipartForm
	photos := m.File["photos"]
	for index, photo := range photos {
		source, openPhotoFileErr := photo.Open()
		defer source.Close()
		checkErr(openPhotoFileErr, "open photo file err")
		fmt.Println("openPhotoFileErr:", openPhotoFileErr)

		fmt.Println("filename:", photo.Filename, " bytes:", photo.Size)

		randString := newRandomString(50)
		photoExt := filepath.Ext(photo.Filename)
		photoPath := randString + photoExt

		destination, createPhotoFileErr := os.Create(frogConfig.StoragePath + "photo/" + photoPath)
		defer destination.Close()
		checkErr(createPhotoFileErr, "create photo file err")

		_, copyErr := io.Copy(destination, source)
		checkErr(copyErr, "copy photo file err")
		fmt.Println("copyErr", copyErr)

		data, decodeErr := exif.Read(frogConfig.StoragePath + "photo/" + photoPath)
		checkWarn(decodeErr, "decode photo exif err")
		fmt.Println("decodeErr", decodeErr)

		if copyErr == nil {
			if decodeErr != nil {
				_, storeRecordPhotoErr := db.Exec("INSERT INTO photo SET userid=?, recordid=?, initorder=?, path=?, name=?, createtime=CURRENT_TIMESTAMP", UID, recordID, index, photoPath, photo.Filename)
				checkErr(storeRecordPhotoErr, "store record photo err")
				if storeRecordPhotoErr == nil {
					successStore = true
				}
			} else {
				for key, value := range data.Tags {
					fmt.Println(key, "=", value)
				}

				latitudePosition := data.Tags["North or South Latitude"]
				longitudePosition := data.Tags["East or West Longitude"]
				latitudeValue := data.Tags["Latitude"]   //緯度
				longitudeValue := data.Tags["Longitude"] //經度
				altitude := data.Tags["Altitude"]        //海拔
				dateAndTime := data.Tags["Date and Time"]
				//GPSDate := data.Tags["GPS Date"] 慢八小時
				charsDateAndTime := []rune(dateAndTime)
				charsDateAndTime[4], charsDateAndTime[7] = '-', '-'
				dateAndTime = string(charsDateAndTime)

				latitude := ""
				longitude := ""
				latitude, longitude = parseCoordinate(latitudeValue, latitudePosition, longitudeValue, longitudePosition)

				_, storeRecordPhotoErr := db.Exec("INSERT INTO photo SET userid=?, recordid=?, initorder=?, path=?, name=?, longitude=?, latitude=?, altitude=?, shootdatetime=?, createtime=CURRENT_TIMESTAMP", UID, recordID, index, photoPath, photo.Filename, longitude, latitude, altitude, dateAndTime)
				checkErr(storeRecordPhotoErr, "store record photo err")
				if storeRecordPhotoErr == nil {
					successStore = true
				}

			}
		}
	}
	return successStore
}

func addPhotosToRecordByRecordID(r *http.Request, UID string) bool {
	addSucccess := false
	m := r.MultipartForm
	recordID := r.Form.Get("recordid")
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
		checkWarn(decodeErr, "decode photo exif err")

		if copyErr == nil {
			if decodeErr != nil {
				_, storeRecordPhotoErr := db.Exec("INSERT INTO photo SET userid=?, recordid=?, initorder=?, path=?, name=?, createtime=CURRENT_TIMESTAMP", UID, recordID, index, photoPath, photo.Filename)
				checkErr(storeRecordPhotoErr, "store record photo err")
				if storeRecordPhotoErr == nil {
					addSucccess = true
				}
			} else {
				latitudePosition := data.Tags["North or South Latitude"]
				longitudePosition := data.Tags["East or West Longitude"]
				latitudeValue := data.Tags["Latitude"]
				longitudeValue := data.Tags["Longitude"]

				latitude := ""
				longitude := ""
				latitude, longitude = parseCoordinate(latitudeValue, latitudePosition, longitudeValue, longitudePosition)

				_, storeRecordPhotoErr := db.Exec("INSERT INTO photo SET userid=?, recordid=?, initorder=?, path=?, name=?, longitude=?, latitude=?, createtime=CURRENT_TIMESTAMP", UID, recordID, index, photoPath, photo.Filename, longitude, latitude)
				checkErr(storeRecordPhotoErr, "store record photo err")
				if storeRecordPhotoErr == nil {
					addSucccess = true
				}

			}
		}
	}
	return addSucccess
}

func getRecordByRecordID(recordID string) Record {
	name, food, stage, season, note := "", "", "", "", ""
	db.QueryRow("SELECT organismname FROM record WHERE id = ?", recordID).Scan(&name)
	db.QueryRow("SELECT food FROM record WHERE id = ?", recordID).Scan(&food)
	db.QueryRow("SELECT stage FROM record WHERE id = ?", recordID).Scan(&stage)
	db.QueryRow("SELECT season FROM record WHERE id = ?", recordID).Scan(&season)
	db.QueryRow("SELECT note FROM record WHERE id = ?", recordID).Scan(&note)
	id, _ := strconv.Atoi(recordID)
	r := Record{
		ID:     id,
		Name:   name,
		Food:   food,
		Stage:  stage,
		Season: season,
		Note:   note}
	r.PhotoSrc = make(map[int]string)
	idrows, queryErr := db.Query("SELECT path FROM photo WHERE recordid = ?", recordID)
	checkErr(queryErr, "query photo path from comment with mysql error")
	defer idrows.Close()
	i := 0
	for idrows.Next() {
		var tmp string
		scanErr := idrows.Scan(&tmp)
		checkErr(scanErr, "scan photo path from comment with mysql error")
		r.PhotoSrc[i] = tmp
		i++
	}
	return r
}

func alterRecordByRecordID(r *http.Request) bool {
	successAlter := true
	r.ParseMultipartForm(32 << 20)
	recordID := r.Form.Get("recordid")
	for key, values := range r.Form {
		for _, value := range values {
			updateCommand := "UPDATE record SET " + key + "=?" + " WHERE id=?"
			// when value quantity == 1, can do this
			_, updateErr := db.Exec(updateCommand, value, recordID)
			if updateErr != nil {
				successAlter = false
			}
		}
	}
	return successAlter
}

func alterRecordPhotoByRecordID(r *http.Request) bool {
	successAlter := true

	return successAlter
}

func removeRecordByRecordID(recordID string) bool {
	successDelete := false
	_, deleteRecordWithMysqlErr := db.Exec("DELETE FROM record WHERE id=?", recordID)
	checkErr(deleteRecordWithMysqlErr, "deleteRecordWithMysqlErr")
	if deleteRecordWithMysqlErr == nil {
		successDelete = true
	}
	return successDelete
}

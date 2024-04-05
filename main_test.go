package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"2024Backend/mydb"
	"2024Backend/redis"

	"github.com/stretchr/testify/assert"
)

func TestPostAd(t *testing.T) {
	mydb.CreateTableSQL()

	if err := mydb.InitializeDB(); err != nil {
		panic(err)
	}
	defer mydb.DB.Close()

	router := setupRouter()

	w := httptest.NewRecorder()

	reqBody := strings.NewReader(`{
		"title": "AD 55",
		"startAt": "2024-04-01",
		"endAt": "2024-05-30",
		"condition": {
		  "ageStart": 20,
		  "ageEnd": 30,
		  "gender": "M",
		  "country": ["TW","US"],
		  "platform": ["web"]
		}
	  }`)

	req, _ := http.NewRequest("POST", "/api/v1/ad", reqBody)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestGetAd(t *testing.T) {
	redis.InitializeRedis()
	if err := mydb.InitializeDB(); err != nil {
		panic(err)
	}
	defer mydb.DB.Close()

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/ad", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}

func TestGetAdParam(t *testing.T) {
	redis.InitializeRedis()
	if err := mydb.InitializeDB(); err != nil {
		panic(err)
	}
	defer mydb.DB.Close()

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/ad?offset=0&ageStart=20&ageEnd=30&country=TW&country=US&platform=web&limit=3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}

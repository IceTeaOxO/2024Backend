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

// 測試建立廣告：全部參數都有
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

// 測試建立廣告：condition缺少所有參數
func TestPostAdParam(t *testing.T) {
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
		}
	  }`)

	req, _ := http.NewRequest("POST", "/api/v1/ad", reqBody)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

// 測試取得廣告：沒有任何參數
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

// 測試取得廣告：使用參數
func TestGetAdParam(t *testing.T) {
	redis.InitializeRedis()
	if err := mydb.InitializeDB(); err != nil {
		panic(err)
	}
	defer mydb.DB.Close()

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/ad?offset=0&age=30&country=TW&country=US&platform=web&limit=3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}

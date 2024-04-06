package main

import (
	"2024Backend/mydb"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"2024Backend/models"

	"2024Backend/redis"
	"2024Backend/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CreateAd(c *gin.Context) {
	var newAd models.Ads
	if err := c.ShouldBindJSON(&newAd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 檢查條件參數是否為空，若為空則設置默認值
	if newAd.Condition.AgeStart == nil {
		defaultValue := 1
		newAd.Condition.AgeStart = &defaultValue
	}

	if newAd.Condition.AgeEnd == nil {
		defaultValue := 100
		newAd.Condition.AgeEnd = &defaultValue
	}

	if len(newAd.Condition.Gender) == 0 {
		newAd.Condition.Gender = ""
	}

	if len(newAd.Condition.Country) == 0 {
		newAd.Condition.Country = []string{}
	}

	if len(newAd.Condition.Platform) == 0 {
		newAd.Condition.Platform = []string{}
	}

	// 參數驗證
	validAgeStart := *newAd.Condition.AgeStart
	validAgeEnd := *newAd.Condition.AgeEnd
	// 驗證年齡值需在1~100之間
	if err := utils.ValidateAge(validAgeStart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := utils.ValidateAge(validAgeEnd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 起始年齡需小於等於結束年齡
	if err := utils.ValidateAgeRange(validAgeStart, validAgeEnd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 驗證country需在ISO_3166-1中，或者為空
	for _, country := range newAd.Condition.Country {
		if !utils.ValidateCountry(country) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country code: " + country})
			return
		}
	}

	stmt, err := mydb.DB.Prepare("INSERT INTO ad (title, startAt, endAt, ageStart, ageEnd, gender, country, platform) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	var ageStart, ageEnd, gender interface{}

	ageStart = newAd.Condition.AgeStart
	ageEnd = newAd.Condition.AgeEnd
	gender = newAd.Condition.Gender

	// 將StringArray轉換為String
	var countryStr string
	if len(newAd.Condition.Country) > 0 {
		countryStr = utils.ConvertStringArrayToString(newAd.Condition.Country)
	} else {
		countryStr = ""
	}

	var platformStr string
	if len(newAd.Condition.Platform) > 0 {
		platformStr = utils.ConvertStringArrayToString(newAd.Condition.Platform)
	} else {
		platformStr = ""
	}

	_, err = stmt.Exec(newAd.Title, newAd.StartAt, newAd.EndAt, ageStart, ageEnd, gender, countryStr, platformStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var lastInsertID int64
	err = mydb.DB.QueryRow("SELECT LAST_INSERT_ID()").Scan(&lastInsertID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newAd.ID = int(lastInsertID)

	c.JSON(http.StatusCreated, gin.H{"message": "Ad created successfully", "ad": newAd})
}

// GET端點
func GetAds(c *gin.Context) {
	// 解析查詢參數
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "5")
	ageStr := c.DefaultQuery("age", "1")
	gender := c.DefaultQuery("gender", "")
	countryArr, _ := c.GetQueryArray("country")
	platformArr, _ := c.GetQueryArray("platform")

	// 轉換成整數方便驗證
	offset, _ := strconv.Atoi(offsetStr)
	limit, _ := strconv.Atoi(limitStr)
	age, _ := strconv.Atoi(ageStr)

	var countryStr string
	if len(countryArr) > 0 {
		countryStr = utils.ConvertStringArrayToString(countryArr)
	} else {
		countryStr = ""
	}
	var platformStr string
	if len(countryArr) > 0 {
		platformStr = utils.ConvertStringArrayToString(platformArr)
	} else {
		platformStr = ""
	}

	// 參數驗證
	if err := utils.ValidateOffsetLimit(offset, limit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := utils.ValidateAge(age); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := utils.ValidateGender(gender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, country := range countryArr {
		if !utils.ValidateCountry(country) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country code: " + country})
			return
		}
	}
	for _, platform := range platformArr {
		if err := utils.ValidatePlatform(platform); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// 構建 Redis 緩存鍵
	cacheKey := "get:offset=" + offsetStr + "&limit=" + limitStr + "&age=" + ageStr + "&gender=" + gender + "&country=" + countryStr + "&platform=" + platformStr

	// 從 Redis 緩存中獲取結果
	cachedResult, err := redis.GetFromCache(cacheKey)
	if err == nil && cachedResult != "" {
		// 如果緩存存在，直接返回緩存結果
		c.JSON(http.StatusOK, cachedResult)
		return
	}

	// 如果緩存不存在，則執行資料庫查詢
	// 構建SQL查詢語句
	// 只查詢活動中的廣告
	query := "SELECT * FROM ad WHERE NOW() > StartAt AND NOW() < EndAt"

	if ageStr != "" {
		query += fmt.Sprintf(" AND %d BETWEEN AgeStart AND AgeEnd", age)
	}
	if gender != "" {
		query += fmt.Sprintf(" AND Gender = '%s'", gender)
	}
	if len(countryArr) > 0 {
		query += fmt.Sprintf(" AND Country LIKE '%%%s%%'", countryStr)
	}
	if len(platformArr) > 0 {
		query += fmt.Sprintf(" AND Platform LIKE '%%%s%%'", platformStr)
	}
	query += " ORDER BY EndAt ASC LIMIT ? OFFSET ?"
	// 執行SQL查詢
	rows, err := mydb.DB.Query(query, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var ads []models.Ad
	for rows.Next() {
		var ad models.Ad
		err := rows.Scan(&ad.ID, &ad.Title, &ad.StartAt, &ad.EndAt, &ad.AgeStart, &ad.AgeEnd, &ad.Gender, &ad.Country, &ad.Platform)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ads = append(ads, ad)
	}

	// 將查詢結果轉換為 JSON 字串
	adsJSON, err := json.Marshal(ads)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	adsStr := string(adsJSON)

	// 將結果存入 Redis 緩存
	err = redis.SetToCache(cacheKey, adsStr, 5*time.Minute) // 設定 5 分鐘過期時間
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ads)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/v1/ad", CreateAd)

	r.GET("/api/v1/ad", GetAds)

	return r
}

func main() {
	mydb.CreateTableSQL()
	// 初始化資料庫連接
	if err := mydb.InitializeDB(); err != nil {
		panic(err)
	}
	defer mydb.DB.Close()

	redis.InitializeRedis()

	r := setupRouter()

	r.Run(":8080")
}

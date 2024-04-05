package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func InitializeRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 伺服器地址
		Password: "",               // Redis 密碼，若無可不填
		DB:       0,                // 使用的 Redis 資料庫
	})

	// 測試 Redis 連接
	pong, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
}

func GetFromCache(key string) (string, error) {
	val, err := RDB.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func SetToCache(key string, value string, expiration time.Duration) error {
	err := RDB.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// func CreateAd(c *gin.Context) {
// 	var newAd Ads
// 	if err := c.ShouldBindJSON(&newAd); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// 將新廣告資訊轉換為 JSON 字串
// 	newAdJSON, err := json.Marshal(newAd)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	newAdStr := string(newAdJSON)

// 	// 在 Redis 中暫存 POST 請求
// 	err = setToCache("post:"+strconv.Itoa(newAd.ID), newAdStr, 5*time.Minute) // 設定 5 分鐘過期時間
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Insert into database
// 	stmt, err := db.Prepare("INSERT INTO ad (title, startAt, endAt, age, gender, country, platform) VALUES (?, ?, ?, ?, ?, ?, ?)")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer stmt.Close()

// 	// Check if fields are valid and assign values accordingly
// 	var age, gender, country, platform interface{}

// 	age = newAd.Condition.Age
// 	gender = newAd.Condition.Gender // gender is now a string, not a sql.NullString
// 	country = newAd.Condition.Country
// 	platform = newAd.Condition.Platform

// 	_, err = stmt.Exec(newAd.Title, newAd.StartAt, newAd.EndAt, age, gender, country, platform)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"message": "Ad created successfully", "ad": newAd})
// }

// func main() {
// 	// 初始化 Redis 連接
// 	initializeRedis()

// 	// 其餘部分保持不變
// }

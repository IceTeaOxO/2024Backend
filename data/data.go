package main

import (
	"2024Backend/models"
	"2024Backend/mydb"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var validCountries = []string{"TW", "US", "CA", "JP", "CN"} // 有效的國家代碼列表
var validPlatforms = []string{"web", "ios", "android"}      // 有效的平台列表

func main() {
	// 初始化資料庫連接
	if err := mydb.InitializeDB(); err != nil {
		panic(err)
	}
	defer mydb.DB.Close()

	// 生成活躍的廣告
	fmt.Println("Generating active ads...")
	for i := 1; i <= 1000; i++ { // 生成1000筆活躍廣告
		ad := generateActiveAd()
		insertAd(ad)
		time.Sleep(2 * time.Millisecond) // 休眠2毫秒
	}

	// 生成不活躍的廣告
	fmt.Println("Generating inactive ads...")
	for i := 1; i <= 10000; i++ { // 生成10000筆不活躍廣告
		ad := generateInactiveAd()
		insertAd(ad)
		time.Sleep(2 * time.Millisecond) // 休眠2毫秒
	}

	fmt.Println("Data generation completed.")
}

func generateActiveAd() models.Ads {
	var ad models.Ads

	// 生成標題
	ad.Title = fmt.Sprintf("Ad %d", rand.Intn(1000)+1)

	// 生成開始日期和結束日期
	startDateStr := "2024-04-01"
	endDateStr := "2024-05-30"
	startTime, _ := time.Parse("2006-01-02", startDateStr)
	endTime, _ := time.Parse("2006-01-02", endDateStr)
	ad.StartAt = startTime.Format("2006-01-02")
	ad.EndAt = endTime.Format("2006-01-02")

	// 生成年齡範圍，50%的概率為空
	if rand.Float32() < 0.5 {
		ageStart := rand.Intn(50) + 1 // 年齡範圍1~50歲
		ad.Condition.AgeStart = &ageStart
	}
	if rand.Float32() < 0.5 {
		ageEnd := rand.Intn(50) + 51 // 年齡範圍51~100歲
		ad.Condition.AgeEnd = &ageEnd
	}

	// 生成性別，25%的概率為空
	if rand.Float32() < 0.75 {
		ad.Condition.Gender = ""
		if rand.Float32() < 0.5 {
			ad.Condition.Gender = "M"
		} else {
			ad.Condition.Gender = "F"
		}
	}

	// 生成國家列表
	countryList := make([]string, 0)
	for i := 0; i < rand.Intn(3)+1; i++ { // 隨機生成1~3個國家代碼
		countryList = append(countryList, validCountries[rand.Intn(len(validCountries))])
	}
	ad.Condition.Country = countryList

	// 生成平台列表
	platformList := make([]string, 0)
	for i := 0; i < rand.Intn(3)+1; i++ { // 隨機生成1~3個平台
		platformList = append(platformList, validPlatforms[rand.Intn(len(validPlatforms))])
	}
	ad.Condition.Platform = platformList
	return ad
}

func generateInactiveAd() models.Ads {
	var ad models.Ads

	// 生成標題
	ad.Title = fmt.Sprintf("Ad %d", rand.Intn(1000)+1)

	// 生成開始日期和結束日期
	startDateStr := "2024-01-01"
	endDateStr := "2024-03-30"
	startTime, _ := time.Parse("2006-01-02", startDateStr)
	endTime, _ := time.Parse("2006-01-02", endDateStr)
	ad.StartAt = startTime.Format("2006-01-02")
	ad.EndAt = endTime.Format("2006-01-02")

	// 生成年齡範圍，50%的概率為空
	if rand.Float32() < 0.5 {
		ageStart := rand.Intn(50) + 1 // 年齡範圍1~50歲
		ad.Condition.AgeStart = &ageStart
	}
	if rand.Float32() < 0.5 {
		ageEnd := rand.Intn(50) + 51 // 年齡範圍51~100歲
		ad.Condition.AgeEnd = &ageEnd
	}

	// 生成性別，25%的概率為空
	if rand.Float32() < 0.75 {
		ad.Condition.Gender = ""
		if rand.Float32() < 0.5 {
			ad.Condition.Gender = "M"
		} else {
			ad.Condition.Gender = "F"
		}
	}

	// 生成國家列表
	countryList := make([]string, 0)
	for i := 0; i < rand.Intn(3)+1; i++ { // 隨機生成1~3個國家代碼
		countryList = append(countryList, validCountries[rand.Intn(len(validCountries))])
	}
	ad.Condition.Country = countryList

	// 生成平台列表
	platformList := make([]string, 0)
	for i := 0; i < rand.Intn(3)+1; i++ { // 隨機生成1~3個平台
		platformList = append(platformList, validPlatforms[rand.Intn(len(validPlatforms))])
	}
	ad.Condition.Platform = platformList
	return ad
}

func insertAd(ad models.Ads) {
	stmt, err := mydb.DB.Prepare("INSERT INTO ad (title, startAt, endAt, ageStart, ageEnd, gender, country, platform) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var ageStart, ageEnd, gender interface{}

	ageStart = ad.Condition.AgeStart
	ageEnd = ad.Condition.AgeEnd
	gender = ad.Condition.Gender

	country := convertToStringList(ad.Condition.Country)
	platform := convertToStringList(ad.Condition.Platform)

	_, err = stmt.Exec(ad.Title, ad.StartAt, ad.EndAt, ageStart, ageEnd, gender, country, platform)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted ad '%s'\n", ad.Title)
}

func convertToStringList(list []string) string {
	return strings.Join(list, ",")
}

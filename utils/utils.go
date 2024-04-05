package utils

import (
	"errors"
	"strings"
)

// ConvertStringArrayToString 將 string 陣列轉換為逗號分隔的字符串
func ConvertStringArrayToString(arr []string) string {
	return strings.Join(arr, ",")
}

func ValidateAge(age int) error {
	if age < 1 || age > 100 {
		return errors.New("年齡必須在 1 到 100 之間")
	}
	return nil
}

func ValidateAgeRange(ageStart, ageEnd int) error {
	if ageStart > ageEnd {
		return errors.New("起始年龄不能大于结束年龄")
	}
	return nil
}

var validCountries = map[string]bool{
	"AF": true, "AX": true, "AL": true, "DZ": true, "AS": true, "AD": true, "AO": true, "AI": true,
	"AQ": true, "AG": true, "AR": true, "AM": true, "AW": true, "AU": true, "AT": true, "AZ": true,
	"BS": true, "BH": true, "BD": true, "BB": true, "BY": true, "BE": true, "BZ": true, "BJ": true,
	"BM": true, "BT": true, "BO": true, "BQ": true, "BA": true, "BW": true, "BV": true, "BR": true,
	"IO": true, "BN": true, "BG": true, "BF": true, "BI": true, "CV": true, "KH": true, "CM": true,
	"CA": true, "KY": true, "CF": true, "TD": true, "CL": true, "CN": true, "CX": true, "CC": true,
	"CO": true, "KM": true, "CG": true, "CD": true, "CK": true, "CR": true, "CI": true, "HR": true,
	"CU": true, "CW": true, "CY": true, "CZ": true, "DK": true, "DJ": true, "DM": true, "DO": true,
	"EC": true, "EG": true, "SV": true, "GQ": true, "ER": true, "EE": true, "SZ": true, "ET": true,
	"FK": true, "FO": true, "FJ": true, "FI": true, "FR": true, "GF": true, "PF": true, "TF": true,
	"GA": true, "GM": true, "GE": true, "DE": true, "GH": true, "GI": true, "GR": true, "GL": true,
	"GD": true, "GP": true, "GU": true, "GT": true, "GG": true, "GN": true, "GW": true, "GY": true,
	"HT": true, "HM": true, "VA": true, "HN": true, "HK": true, "HU": true, "IS": true, "IN": true,
	"ID": true, "IR": true, "IQ": true, "IE": true, "IM": true, "IL": true, "IT": true, "JM": true,
	"JP": true, "JE": true, "JO": true, "KZ": true, "KE": true, "KI": true, "KP": true, "KR": true,
	"KW": true, "KG": true, "LA": true, "LV": true, "LB": true, "LS": true, "LR": true, "LY": true,
	"LI": true, "LT": true, "LU": true, "MO": true, "MK": true, "MG": true, "MW": true, "MY": true,
	"MV": true, "ML": true, "MT": true, "MH": true, "MQ": true, "MR": true, "MU": true, "YT": true,
	"MX": true, "FM": true, "MD": true, "MC": true, "MN": true, "ME": true, "MS": true, "MA": true,
	"MZ": true, "MM": true, "NA": true, "NR": true, "NP": true, "NL": true, "NC": true, "NZ": true,
	"NI": true, "NE": true, "NG": true, "NU": true, "NF": true, "MP": true, "NO": true, "OM": true,
	"PK": true, "PW": true, "PS": true, "PA": true, "PG": true, "PY": true, "PE": true, "PH": true,
	"PN": true, "PL": true, "PT": true, "PR": true, "QA": true, "RE": true, "RO": true, "RU": true,
	"RW": true, "BL": true, "SH": true, "KN": true, "LC": true, "MF": true, "PM": true, "VC": true,
	"WS": true, "SM": true, "ST": true, "SA": true, "SN": true, "RS": true, "SC": true, "SL": true,
	"SG": true, "SX": true, "SK": true, "SI": true, "SB": true, "SO": true, "ZA": true, "GS": true,
	"SS": true, "ES": true, "LK": true, "SD": true, "SR": true, "SJ": true, "SE": true, "CH": true,
	"SY": true, "TW": true, "TJ": true, "TZ": true, "TH": true, "TL": true, "TG": true, "TK": true,
	"TO": true, "TT": true, "TN": true, "TR": true, "TM": true, "TC": true, "TV": true, "UG": true,
	"UA": true, "AE": true, "GB": true, "US": true, "UM": true, "UY": true, "UZ": true, "VU": true,
	"VE": true, "VN": true, "VG": true, "VI": true, "WF": true, "EH": true, "YE": true, "ZM": true,
	"ZW": true,
}

func ValidateCountry(country string) bool {
	// 將代碼轉換為大寫
	country = strings.ToUpper(country)

	// 檢查國家代碼是否在指定列表中
	_, ok := validCountries[country]
	return ok
}

// ValidateOffsetLimit 驗證 offset 和 limit 参数是否為非負整數
func ValidateOffsetLimit(offset, limit int) error {
	if offset < 0 && limit <= 0 {
		return errors.New("offset和limit需要大於等於0")
	}
	return nil
}

// ValidateGender 驗證性別參數
func ValidateGender(gender string) error {
	if gender != "M" && gender != "F" && gender != "" {
		return errors.New("性別需要是 M 或 F 或者空白")
	}
	return nil
}

// ValidatePlatform
func ValidatePlatform(platform string) error {
	if platform != "web" && platform != "ios" && platform != "android" {
		return errors.New("platform 需要是 web, ios, android")
	}
	return nil
}

package Utils

import (
	"fmt"
	"strconv"
	"time"
)

func Int64ToString(x int64) string {
	return strconv.FormatInt(x, 10)
}

func StringToInt64(s string) int64 {
	temp, _ := strconv.ParseInt(s, 10, 64)
	return temp
}

func Float64ToString(x float64) string {
	return strconv.FormatFloat(x, 'f', -1, 64)
}

func StringToFloat64(s string) float64 {
	temp, _ := strconv.ParseFloat(s, 64)
	return temp
}

func BoolToString(x bool) string {
	return strconv.FormatBool(x)
}

func StringToBool(s string) bool {
	temp, _ := strconv.ParseBool(s)
	return temp
}

func TimeToString(x time.Time) string {
	return x.Format("02/01/2006 15:04:05")
}

func StringToTime(s string) time.Time {
	temp, _ := time.Parse("2006-01-02T15:04:05.000", s)
	return temp
}

func ConvertInterface(value interface{}) string {
	switch v := value.(type) {
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", value)
	}
}

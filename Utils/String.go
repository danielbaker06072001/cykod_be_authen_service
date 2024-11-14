package Utils

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

func GenerateSearchKeyword(s string) string {
	if s != "" {
		s = RemoveAccent(s)
		s = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s, "")
		return strings.ToLower(strings.Replace(norm.NFD.String(s), " ", "", -1))
	}
	return ""
}

func GenerateUnsignName(s string) string {
	if s != "" {
		s = RemoveAccent(s)
		return strings.ToLower(strings.Replace(norm.NFD.String(s), " ", "", -1))
	}
	return ""
}

func UpperFirstChar(s string) string {
	//ex: channelConfigID -> ChannelConfigID
	return strings.ToUpper(s[0:1]) + s[1:]
}

func GetSortedFieldQuery(sortedField string, sortedDirection bool, table string) string {
	//ex: channelConfigID -> "Newss"."ChannelConfigID" | "Newss"."ChannelConfigID" desc
	var sortedFieldQuery string = table + ".\"" + UpperFirstChar(sortedField) + "\""
	if sortedDirection == true {
		sortedFieldQuery = sortedFieldQuery + " desc"
	}
	return sortedFieldQuery
}

// Mang cac ky tu goc co dau var
var SOURCE_CHARACTERS, LL_LENGTH = stringToRune("ÀÁÂÃÈÉÊÌÍÒÓÔÕÙÚÝàáâãèéêìíòóôõùúýĂăĐđĨĩŨũƠơƯưẠạẢảẤấẦầẨẩẪẫẬậẮắẰằẲẳẴẵẶặẸẹẺẻẼẽẾếỀềỂểỄễỆệỈỉỊịỌọỎỏỐốỒồỔổỖỗỘộỚớỜờỞởỠỡỢợỤụỦủỨứỪừỬửỮữỰự")

// Mang cac ky tu thay the khong dau var
var DESTINATION_CHARACTERS, _ = stringToRune(`AAAAEEEIIOOOOUUYaaaaeeeiioooouuyAaDdIiUuOoUuAaAaAaAaAaAaAaAaAaAaAaAaEeEeEeEeEeEeEeEeIiIiOoOoOoOoOoOoOoOoOoOoOoOoUuUuUuUuUuUuUu`)

func stringToRune(s string) ([]string, int) {
	ll := utf8.RuneCountInString(s)
	var texts = make([]string, ll+1)
	var index = 0
	for _, runeValue := range s {
		texts[index] = string(runeValue)
		index++
	}
	return texts, ll
}
func binarySearch(sortedArray []string, key string, low int, high int) int {
	var middle int = (low + high) / 2
	if high < low {
		return -1
	}
	if key == sortedArray[middle] {
		return middle
	} else if key < sortedArray[middle] {
		return binarySearch(sortedArray, key, low, middle-1)
	} else {
		return binarySearch(sortedArray, key, middle+1, high)
	}
}
func removeAccentChar(ch string) string {
	var index int = binarySearch(SOURCE_CHARACTERS, ch, 0, LL_LENGTH)
	if index >= 0 {

		ch = DESTINATION_CHARACTERS[index]
	}
	return ch
}
func RemoveAccent(s string) string {
	var buffer bytes.Buffer
	for _, runeValue := range s {
		buffer.WriteString(removeAccentChar(string(runeValue)))
	}
	return buffer.String()

}

func ContainsOnlyNumbers(input string) error {
	_, err := regexp.MatchString("^[0-9]+$", input)
	return err
}

func ConvertStringtoTime(input string) (time.Time, error) {
	var rTime time.Time
	var err error

	// Parse input string into time.Time object
	var parsedTime time.Time
	if len(input) == len("2006-01-02T15:04:05") {
		parsedTime, err = time.Parse("2006-01-02T15:04:05", input)
	} else {
		parsedTime, err = time.Parse("2006-01-02T15:04:05.000", input)
	}
	if err != nil {
		return rTime, err
	}

	// Format time.Time object into desired layout
	formattedTime := parsedTime.Format("2006-01-02 15:04:05.000")

	// Parse formatted time back to time.Time object
	parsedBackTime, err := time.Parse("2006-01-02 15:04:05.000", formattedTime)
	if err != nil {
		return rTime, err
	}
	rTime = parsedBackTime
	return rTime, err
}

// Chuyển các giá trị bất ky ở dạng interface{} về dạng string
func ConvertToBoolString(value interface{}) string {
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

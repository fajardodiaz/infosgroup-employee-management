package utils

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	randomDataTime "github.com/duktig-solutions/go-random-date-generator"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// // Random int between min and max
func RandonInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomGender() string {
	genders := []string{"M", "F", "O"}
	n := len(genders)
	return genders[rand.Intn(n)]
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomBirthDate() time.Time {
	randomDate, err := randomDataTime.GenerateDOB(18, 65)
	if err != nil {
		log.Fatal(err)
	}

	date, _ := time.Parse("2006-01-02", randomDate)

	return date
}

func RandomIngressDate() time.Time {

	randomDate, err := randomDataTime.GenerateDate("2018-01-01", "2023-01-01")
	if err != nil {
		log.Fatal(err)
	}

	date, _ := time.Parse("2006-01-02", randomDate)

	return date

}

func Add3MonthsToDate(date time.Time) time.Time {
	newDate := date.AddDate(0, 3, 0)

	return newDate
}

// // Random int between min and max
func RandonPhone(min, max int64) string {

	phone := min + rand.Int63n(max-min+1)

	newStringPhone := strconv.Itoa(int(phone))

	return newStringPhone
}

package date_time

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"time"
)

type Months int
type Dates int

const (
	January Months = iota << 1
	February
	March
	Apirl
	May
	June
	July
	August
	September
	October
	November
	December
	AfterDecember
)

const CURRENT_YEAR int = 2022
const INTRL_ERROR string = "internal error"

const (
	One Dates = iota << 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Eleven
	Twelve
	Thirteen
	Forteen
	Fifteen
	Sixteen
	Seventeen
	Eighteen
	Nineeen
	Twenty
	Twenty_One
	Twenty_Two
	Twenty_Three
	Twenty_Four
	Twenty_Five
	Twenty_Six
	Twenty_Seven
	Twenty_Eight
	Twenty_Nine
	Thirty
	Thirty_One
	Thirty_Two
)

// DATE PARSE
func Date(s string) (int, error) {

	if ok, err := regexp.MatchString("[0-9]+", s[8:9]); err == nil && !ok {

		log.Fatalln(" Error date is not a valid", s[8:9])
		return -1, err
	}

	if date, err := strconv.Atoi(s[8:10]); err == nil && date < int(One) {
		log.Fatalln(" Error month is not a valid", s[8:10])
		return -1, err
	}

	date, _ := strconv.Atoi(s[8:10])

	if date >= int(One) && date < int(Thirty_Two) {
		return date, nil
	}

	return -1, errors.New(INTRL_ERROR)
}

// MONTH PARSE
func Month(s string) (int, error) {

	if ok, err := regexp.MatchString("[0-9]+", s[5:7]); err == nil && !ok {
		log.Fatalln(" Error month is not a valid", s[5:7])
		return -1, err
	}

	if month, err := strconv.Atoi(s[5:7]); err == nil && month < int(January) {
		log.Fatalln(" Error month is not a valid", s[5:7])
		return -1, err
	}

	month, _ := strconv.Atoi(s[5:7])

	if month >= int(January) && month < int(AfterDecember) {
		return month, nil
	}

	return -1, errors.New(INTRL_ERROR)

}

// YEAR PARSE
func Year(s string) (int, error) {

	if ok, err := regexp.MatchString("[0-9]+", s[0:4]); err == nil && !ok {
		log.Fatalln(" Error year is not a valid", s[0:4])
		return -1, err
	}

	if year, err := strconv.Atoi(s[0:4]); err == nil && year < CURRENT_YEAR {
		log.Fatalln(" Error year had been passed", s[0:4])
		return -1, err
	}

	year, _ := strconv.Atoi(s[0:4])
	return year, nil
}

func GetToday(year int, month time.Month, date int) time.Time {

	return time.Date(year, month, date, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.UTC)
}

func Elasped(t time.Time, start time.Time) time.Duration {

	return time.Duration(t.Sub(<-time.After(time.Duration(start.Second()))))
}

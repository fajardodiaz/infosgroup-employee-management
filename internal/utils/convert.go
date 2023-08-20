package utils

import (
	"fmt"
	"time"
)

func ConvertToDate(date_to_convert string) time.Time {

	inputDate := date_to_convert
	desiredLayout := "2006-01-02T15:04:05.999999Z"

	// Step 1: Parse the input date string
	parsedTime, err := time.Parse("2006-01-02", inputDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}

	// Step 2: Format the parsed time into the desired layout
	formattedTime := parsedTime.Format(desiredLayout)

	my_date, err := time.Parse(desiredLayout, formattedTime)
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}

	return my_date
}

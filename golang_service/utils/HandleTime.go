package res

import (
	"fmt"
	"math/big"
	"time"
)

func EpochTimeToDate(timestamp *big.Int) time.Time {

	return time.Unix(int64(timestamp.Uint64()), 0)
}

func FormatTime(stringTime string) time.Time {
	date, error := time.Parse("2006-01-02", stringTime)
	if error != nil {
		fmt.Println(error)
	}
	return date
}

func SubtractTime(end time.Time, start time.Time) int {
	number := end.Sub(start)
	fmt.Println(int(number.Hours() / 24)) // number of days
	return int(number.Hours()/24) + 1

}

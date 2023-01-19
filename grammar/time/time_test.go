package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeMethod(t *testing.T) {
	var now = time.Now()
	fmt.Printf("current time:%v\n", now)
	var year = now.Year()
	var month = now.Month()
	var day = now.Day()
	var hour = now.Hour()
	var minute = now.Minute()
	var second = now.Second()
	fmt.Println(year, month, day, hour, minute, second)
}

func TestZone(t *testing.T) {
	secondEastOfUTC := int((time.Hour * 8).Seconds())
	location := time.FixedZone("Beijing Time", secondEastOfUTC)
	newYork, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("load Asia/Tokyo location failed", err)
		return
	}
	fmt.Println()

	localDate := time.Date(2023, 1, 19, 13, 38, 39, 0, location)
	localDate2 := time.Date(2023, 1, 19, 13, 38, 39, 0, newYork)
	fmt.Printf("%v\n", localDate)
	fmt.Printf("%v\n", localDate2)

}

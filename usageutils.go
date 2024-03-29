package utils

import (
	"fmt"
	"math"
	"strconv"
)

func HumanSize(size int) string {
	base := math.Log(float64(size)) / math.Log(1024)
	getSize := round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	suffix := "B"

	switch int(math.Floor(base)) {
	case 0:
		suffix = "B"
	case 1:
		suffix = "KB"
	case 2:
		suffix = "MB"
	case 3:
		suffix = "GB"
	case 4:
		suffix = "TB"
	}
	return fmt.Sprint(strconv.FormatFloat(getSize, 'f', -1, 64), suffix)

}

func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

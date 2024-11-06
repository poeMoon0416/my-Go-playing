package reverse

import (
	"strconv"
)

/*  Int returns reverse int. */
func Int(i int) int {
	revI, _ := strconv.Atoi(String(strconv.Itoa(i)))
	return revI
}

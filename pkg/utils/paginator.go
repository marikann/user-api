package utils

import (
	"strconv"
)

func PaginationWithDefaults(limit, offset string) (int64, int64) {

	//To generate a cumulative error message for both offset and limit values to be corrupt.
	limitInt, limErr := strconv.ParseInt(limit, 10, 64)
	if limit == "" || limErr != nil || limitInt <= 0 {
		limitInt = 25
	} else if limitInt > 100 {
		limitInt = 100
	}
	offsetInt, oErr := strconv.ParseInt(offset, 10, 64)
	if offset == "" || oErr != nil || offsetInt < 0 {
		offsetInt = 0
	}

	return limitInt, offsetInt
}

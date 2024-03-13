package restlib

import "strconv"

func AsInt64(value string) (int64, error) {
	int64Value, err := strconv.ParseInt(value, 10, 64)
	return int64Value, err
}

package models

import "time"

func GetOffset(page int, pageSize int) int {
	var offset int
	if page <= 1 {
		offset = 0
	} else {

		offset = (page - 1) * pageSize
	}

	return offset
}
func TimeToString(time time.Time )(string)  {
	return time.Format("2006-01-02 15:04:05")
}

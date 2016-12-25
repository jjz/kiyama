package models

const (
	DELETED_NORMAL=0
	DELETED_DELETED=1
)
func GetOffset(page int, pageSize int) int {
	var offset int
	if page <= 1 {
		offset = 0
	} else {

		offset = (page - 1) * pageSize
	}

	return offset
}

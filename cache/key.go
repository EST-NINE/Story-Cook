package cache

import (
	"fmt"
	"strconv"
)

func UserCountKey(id uint) string {
	return fmt.Sprintf("UserCount:%s", strconv.Itoa(int(id)))
}

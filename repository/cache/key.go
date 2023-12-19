package cache

import (
	"fmt"
	"strconv"
)

func UserCountKey(id uint) string {
	return fmt.Sprintf("UserID:%s", strconv.Itoa(int(id)))
}

func DeleteUserCountKeys() error {
	keys, err := RedisClient.Keys("UserID:*").Result()
	if err != nil {
		return err
	}

	// 删除所有匹配的键
	if len(keys) > 0 {
		err = RedisClient.Del(keys...).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

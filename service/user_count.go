package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"story-cook-be/dao"
	"story-cook-be/model"
	"story-cook-be/pkg/util"
	"time"
)

type UserCountSrv struct {
}

func (s *UserCountSrv) Update(ctx *gin.Context) error {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	countDao := dao.NewUseCountDao(ctx)
	date := time.Now().Format("2006-01-02")
	userCount, err := countDao.FindUserCountByDate(userInfo.Id, date)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		userCount := &model.UserCount{
			UID:  userInfo.Id,
			Date: date,
		}

		// 没有count记录就创建一条记录
		if err = countDao.CreateUserCount(userCount); err != nil {
			return err
		}
	}

	// 判断count次数
	if userCount.Count >= 5 {
		return errors.New("今日次数已用完,请回味一下今日份故事吧")
	}

	userCount.Count++
	fmt.Println(userCount.Count)
	if err := countDao.UpdateUserCount(userCount); err != nil {
		return err
	}
	fmt.Println(userCount.Count)
	return nil
}

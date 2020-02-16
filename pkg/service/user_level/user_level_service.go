package user_level

import (
	"github.com/arpb2/C-3PO/api/model"
	userlevelservice "github.com/arpb2/C-3PO/api/service/user_level"
)

func CreateService() userlevelservice.Service {
	return &userLevelService{}
}

var inMemory = map[uint]map[uint]*model.UserLevel{} // TODO: For the moment (for testing) is a super simple single in-memory holder without err handling as userId:levelId:code

type userLevelService struct{}

func (c *userLevelService) GetUserLevel(userId uint, levelId uint) (userLevel *model.UserLevel, err error) {
	return inMemory[userId][levelId], nil
}

func (c *userLevelService) CreateUserLevel(userId uint, code string) (userLevel *model.UserLevel, err error) {
	inMemory[userId] = map[uint]*model.UserLevel{}
	inMemory[userId][1] = &model.UserLevel{
		UserId:  userId,
		LevelId: 1,
		Code:    code,
	}
	return inMemory[userId][1], nil
}

func (c *userLevelService) ReplaceUserLevel(userLevel *model.UserLevel) error {
	inMemory[userLevel.UserId] = map[uint]*model.UserLevel{}
	inMemory[userLevel.UserId][1] = userLevel
	return nil
}

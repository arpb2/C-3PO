package user_level

import (
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	userlevelservice "github.com/arpb2/C-3PO/api/service/user_level"
)

func CreateService() userlevelservice.Service {
	return &userLevelService{}
}

var inMemory = map[uint]map[uint]model.UserLevel{} // TODO: For the moment (for testing) is a super simple single in-memory holder without err handling as userId:levelId:code

type userLevelService struct{}

func (c *userLevelService) GetUserLevel(userId uint, levelId uint) (userLevel model.UserLevel, err error) {
	val, ok := inMemory[userId][levelId]
	if !ok {
		return val, http.CreateNotFoundError()
	}
	return val, nil
}

func (c *userLevelService) StoreUserLevel(data model.UserLevel) (result model.UserLevel, err error) {
	inMemory[data.UserId] = map[uint]model.UserLevel{}
	inMemory[data.UserId][data.LevelId] = data
	return inMemory[data.UserId][data.LevelId], nil
}

package code_service

import (
	"github.com/arpb2/C-3PO/api/model"
	code_service "github.com/arpb2/C-3PO/api/service/code"
)

func CreateService() code_service.Service {
	return &codeService{}
}

var inMemory = map[uint]map[uint]*model.Code{} // TODO: For the moment (for testing) is a super simple single in-memory holder without err handling as userId:codeId:code

type codeService struct{}

func (c *codeService) GetCode(userId uint, codeId uint) (code *model.Code, err error) {
	return inMemory[userId][codeId], nil
}

func (c *codeService) CreateCode(userId uint, code string) (codeModel *model.Code, err error) {
	inMemory[userId] = map[uint]*model.Code{}
	inMemory[userId][1] = &model.Code{
		UserId: userId,
		Id:     1,
		Code:   code,
	}
	return inMemory[userId][1], nil
}

func (c *codeService) ReplaceCode(code *model.Code) error {
	inMemory[code.UserId] = map[uint]*model.Code{}
	inMemory[code.UserId][1] = code
	return nil
}

package code_service

import "github.com/arpb2/C-3PO/src/api/service"

func CreateService() service.CodeService {
	return &codeService{}
}

var inMemory = map[uint]map[uint]*string{} // TODO: For the moment (for testing) is a super simple single in-memory holder without err handling as userId:codeId:code

type codeService struct{}
func (c *codeService) GetCode(userId uint, codeId uint) (code *string, err error) {
	return inMemory[userId][codeId], nil
}

func (c *codeService) CreateCode(userId uint, code *string) (codeId uint, err error) {
	inMemory[userId] = map[uint]*string{}
	inMemory[userId][1] = code
	return 1, nil
}

func (c *codeService) ReplaceCode(userId uint, codeId uint, code *string) error {
	inMemory[userId] = map[uint]*string{}
	inMemory[userId][1] = code
	return nil
}

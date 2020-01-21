package code_service

import "github.com/arpb2/C-3PO/src/api/service"

func GetService() service.CodeService {
	return globalRef
}

var globalRef = &codeService{
	InMemory: map[uint]map[uint]*string{}, // TODO Same as below
}

type codeService struct {
	InMemory map[uint]map[uint]*string // TODO: For the moment (for testing) is a super simple single in-memory holder without err handling as userId:codeId:code
}
func (c *codeService) GetCode(userId uint, codeId uint) (code *string, err error) {
	return c.InMemory[userId][codeId], nil
}

func (c *codeService) CreateCode(userId uint, code *string) (codeId uint, err error) {
	c.InMemory[userId] = map[uint]*string{}
	c.InMemory[userId][1] = code
	return 1, nil
}

func (c *codeService) ReplaceCode(userId uint, codeId uint, code *string) error {
	c.InMemory[userId] = map[uint]*string{}
	c.InMemory[userId][1] = code
	return nil
}

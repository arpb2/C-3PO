package code_service

import "github.com/arpb2/C-3PO/src/api/service"

func GetService() service.CodeService {
	return globalRef
}

var globalRef = &codeService{
	InMemory: map[string]map[string]*string{}, // TODO Same as below
}

type codeService struct {
	InMemory map[string]map[string]*string // TODO: For the moment (for testing) is a super simple single in-memory holder without err handling as userId:codeId:code
}
func (c *codeService) GetCode(userId string, codeId string) (code *string, err error) {
	return c.InMemory[userId][codeId], nil
}

func (c *codeService) CreateCode(userId string, code *string) (codeId string, err error) {
	c.InMemory[userId] = map[string]*string{}
	c.InMemory[userId]["1"] = code
	return "1", nil
}

func (c *codeService) ReplaceCode(userId string, codeId string, code *string) error {
	c.InMemory[userId] = map[string]*string{}
	c.InMemory[userId]["1"] = code
	return nil
}

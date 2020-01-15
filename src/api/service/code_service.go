package service

type CodeService interface {

	GetCode(userId string, codeId string) (code *string, err error)

	CreateCode(userId string, code *string) (codeId string, err error)

	ReplaceCode(userId string, codeId string, code *string) error

}
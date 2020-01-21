package service

type CodeService interface {

	GetCode(userId uint, codeId uint) (code *string, err error)

	CreateCode(userId uint, code *string) (codeId uint, err error)

	ReplaceCode(userId uint, codeId uint, code *string) error

}
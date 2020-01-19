package service

type CodeService interface {

	Read(userId string, codeId string) (code *string, err error)

	Write(userId string, code *string) (codeId string, err error)

	Replace(userId string, codeId string, code *string) error

}
package code_service

import "github.com/arpb2/C-3PO/api/model"

type Service interface {
	GetCode(userId uint, codeId uint) (code *model.Code, err error)

	CreateCode(userId uint, code string) (model *model.Code, err error)

	ReplaceCode(code *model.Code) error
}

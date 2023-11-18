package api

import (
	"changeme/pkg/sqlite"
)

type OperatorApi struct {
	opDb *sqlite.OperatorDb
}

func NewOperatorApi() *OperatorApi {
	return &OperatorApi{
		opDb: sqlite.GetOperatorDb(),
	}
}

func (o *OperatorApi) List(req ListRequest) BaseResponse {
	defer recoverFromPanic()

	ops, err := o.opDb.List(req.Limit)
	if err != nil {
		return RespErr(CodeFail, err.Error())
	}
	return RespSuccess(ops)
}

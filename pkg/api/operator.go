package api

import (
	"github.com/sirupsen/logrus"

	"changeme/pkg/log"
	"changeme/pkg/models"
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

func (o *OperatorApi) List(req models.ListRequest) models.BaseResponse {
	defer recoverFromPanic()

	ops, err := o.opDb.List(req.Limit)
	if err != nil {
		return RespErr(models.CodeFail, err.Error())
	}
	log.Log.WithFields(logrus.Fields{
		"ops": ops,
	}).Info("list operator success")
	return RespSuccess(ops)
}

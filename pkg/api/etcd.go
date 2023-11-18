package api

import (
	"fmt"

	"changeme/pkg/log"
	"changeme/pkg/models"
	"changeme/pkg/service"
	"changeme/pkg/sqlite"
)

// return: [是否允许操作，不允许操作的原因]
func checkActionAllow(kType KeyType, action Action) (bool, string) {
	allowAction, exist := dTypeActionMap[kType]
	if !exist {
		return false, fmt.Sprintf("key type %s not exist", kType)
	}
	_, exist = allowAction[action]
	if !exist {
		return false, fmt.Sprintf("key type %s does not allow operation %s", kType, action)
	}
	return true, ""
}

type EtcdApi struct {
	wholeKeySvc *service.WholeKeyService

	opDb *sqlite.OperatorDb
}

func NewEtcdApi() *EtcdApi {
	return &EtcdApi{
		wholeKeySvc: service.GetWholeKeyService(),
		opDb:        sqlite.GetOperatorDb(),
	}
}

func (e *EtcdApi) DoAction(req BaseRequest) BaseResponse {
	allow, reason := checkActionAllow(req.KeyType, req.Action)
	if !allow {
		return RespErr(CodeActionNotAllowed, reason)
	}

	// add operator record
	errMsg := ""
	defer func() {
		var op *models.Operation
		if errMsg != "" {
			op = models.NewOperationFailed(req.KeyType.String(), req.Action.String(), errMsg)
		} else {
			op = models.NewOperationSuccess(req.KeyType.String(), req.Action.String())
		}
		_ = e.opDb.Insert(op)
	}()
	// do action
	switch e.combTypeAction(req.KeyType, req.Action) {
	case e.combTypeAction(DTypeWholeKey, ActionGet):
		val, err := e.wholeKeySvc.Get(req.Data)
		if err != nil {
			errMsg = err.Error()
			log.Log.Errorf("get whole key [%s] failed: [%s], return: [%v]", req.Data, err.Error(), RespErr(CodeFail, err.Error()))
			return RespErr(CodeFail, err.Error())
		}
		return RespSuccess(val)
	default:
		return RespErr(CodeFail, fmt.Sprintf("key type [%s] not support action [%s]", req.KeyType, req.Action))
	}
}

func (e *EtcdApi) combTypeAction(t KeyType, a Action) string {
	return fmt.Sprintf("%s_%s", t, a)
}

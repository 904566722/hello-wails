package api

import (
	"fmt"

	"changeme/pkg/log"
	mdl "changeme/pkg/models"
	"changeme/pkg/service"
	"changeme/pkg/sqlite"
)

// dTypeActionMap 定义 key 类型允许的操作
var dTypeActionMap = map[mdl.KeyType]map[mdl.Action]struct{}{
	// 允许通过完整的 key 进行：1.get 操作；2.put 操作；3.delete 操作；4.list 操作；5.list value 操作
	mdl.DTypeWholeKey: {
		mdl.ActionGet:       struct{}{}, // 允许通过完整的 key 进行 get 操作
		mdl.ActionPut:       struct{}{}, // 允许通过完整的 key 进行 put 操作
		mdl.ActionDelete:    struct{}{}, // 允许通过完整的 key 进行 delete 操作
		mdl.ActionList:      struct{}{}, // 允许通过完整的 key 进行 list 操作
		mdl.ActionListValue: struct{}{}, // 允许通过完整的 key 进行 list value 操作
	},
	// 允许通过前缀进行：1.删除所有key对应的值；2.查询所有key；3.查询所有value；4.查询所有key和value
	mdl.DTypePrefix: {
		mdl.ActionDelete:    struct{}{}, // 允许通过前缀进行 delete 操作 (需要询问是否删除全部的key)
		mdl.ActionListKey:   struct{}{}, // 允许通过前缀进行查询所有 key 操作
		mdl.ActionListValue: struct{}{}, // 允许通过前缀进行查询所有 value 操作
		mdl.ActionList:      struct{}{}, // 允许通过前缀进行查询所有 key 和 value 操作
	},
	// 允许通过关键字进行：1.删除所有key对应的值；2.查询所有key；3.查询所有value；4.查询所有key和value
	mdl.DTypeKeyword: {
		mdl.ActionDelete:    struct{}{}, // 允许通过关键字进行 delete 操作 (需要询问是否删除全部的key)
		mdl.ActionListKey:   struct{}{}, // 允许通过关键字进行查询所有 key 操作
		mdl.ActionListValue: struct{}{}, // 允许通过关键字进行查询所有 value 操作
		mdl.ActionList:      struct{}{}, // 允许通过关键字进行查询所有 key 和 value 操作
	},
}

// return: [是否允许操作，不允许操作的原因]
func checkActionAllow(kType mdl.KeyType, action mdl.Action) (bool, string) {
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
	wholeKeySvc  *service.WholeKeyService
	prefixKeySvc *service.PrefixKeyService
	keywordSvc   *service.KeywordService

	opDb *sqlite.OperatorDb
}

func NewEtcdApi() *EtcdApi {
	return &EtcdApi{
		wholeKeySvc:  service.GetWholeKeyService(),
		prefixKeySvc: service.GetPrefixKeyService(),
		keywordSvc:   service.GetKeywordService(),
		opDb:         sqlite.GetOperatorDb(),
	}
}

func (e *EtcdApi) DoAction(req mdl.BaseRequest) mdl.BaseResponse {
	allow, reason := checkActionAllow(req.KeyType, req.Action)
	if !allow {
		return RespErr(mdl.CodeActionNotAllowed, reason)
	}

	log.Log.Debugf("do action: [%s] [%s] [%s] [%s]", req.KeyType, req.Action, req.Key, req.Value)

	var (
		err      error
		finalVal interface{}
	)

	// add operator record
	defer func() {
		var op *mdl.Operation
		if err != nil {
			op = mdl.NewOperationFailed(req.KeyType, req.Action, req.Key, req.Value, err.Error())
		} else {
			op = mdl.NewOperationSuccess(req.KeyType, req.Action, req.Key, req.Value)
		}
		_ = e.opDb.Insert(op)
	}()
	// do action
	switch e.combTypeAction(req.KeyType, req.Action) {
	// 根据完整的 key 进行操作
	case e.combTypeAction(mdl.DTypeWholeKey, mdl.ActionGet):
		finalVal, err = e.wholeKeySvc.Get(req.Key)
		return e.ret(finalVal, err)
	case e.combTypeAction(mdl.DTypeWholeKey, mdl.ActionPut):
		err = e.wholeKeySvc.Put(req.Key, req.Value)
		return e.ret(nil, err)
	case e.combTypeAction(mdl.DTypeWholeKey, mdl.ActionDelete):
		err = e.wholeKeySvc.Del(req.Key)
		return e.ret(nil, err)

	// 根据前缀进行操作
	case e.combTypeAction(mdl.DTypePrefix, mdl.ActionDelete):
		var delCnt int64
		delCnt, err = e.prefixKeySvc.Del(req.Key)
		return e.ret(e.delMsg(req.KeyType, req.Key, delCnt), err)
	case e.combTypeAction(mdl.DTypePrefix, mdl.ActionListKey):
		finalVal, err = e.prefixKeySvc.ListKey(req.Key)
		return e.ret(finalVal, err)
	case e.combTypeAction(mdl.DTypePrefix, mdl.ActionListValue):
		finalVal, err = e.prefixKeySvc.ListValue(req.Key)
		return e.ret(finalVal, err)
	case e.combTypeAction(mdl.DTypePrefix, mdl.ActionList):
		return e.ret("todo:暂未完成该功能", nil)

	// 根据关键字进行操作
	case e.combTypeAction(mdl.DTypeKeyword, mdl.ActionDelete):
		var delCnt int64
		delCnt, err = e.keywordSvc.Del(req.Key)
		return e.ret(e.delMsg(req.KeyType, req.Key, delCnt), err)
	case e.combTypeAction(mdl.DTypeKeyword, mdl.ActionListKey):
		finalVal, err = e.keywordSvc.ListKey(req.Key)
		return e.ret(finalVal, err)
	case e.combTypeAction(mdl.DTypeKeyword, mdl.ActionListValue):
		finalVal, err = e.keywordSvc.ListValue(req.Key)
		return e.ret(finalVal, err)
	case e.combTypeAction(mdl.DTypeKeyword, mdl.ActionList):
		return e.ret("todo:暂未完成该功能", nil)
	default:
		return RespErr(mdl.CodeFail, fmt.Sprintf("key type [%s] not support action [%s]", req.KeyType, req.Action))
	}
}

// combTypeAction
func (e *EtcdApi) combTypeAction(t mdl.KeyType, a mdl.Action) string {
	return fmt.Sprintf("%s_%s", t, a)
}

func (e *EtcdApi) ret(val interface{}, err error) mdl.BaseResponse {
	if err != nil {
		log.Log.Errorf("do action failed: [%v]", err)
		return RespErr(mdl.CodeFail, err.Error())
	}
	return RespSuccess(val)
}

func (e *EtcdApi) delMsg(kType mdl.KeyType, key string, delCnt int64) string {
	return fmt.Sprintf("根据「%s」(%s)共删除 %d 条记录", kType.Chinese(), key, delCnt)
}

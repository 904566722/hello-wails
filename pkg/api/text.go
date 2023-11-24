package api

import (
	"fmt"

	"changeme/pkg/log"
	"changeme/pkg/models"
	"changeme/pkg/utils"
)

const (
	historyLen = 10
)

type TextApi struct {
	TempJson string

	history       map[string]string
	histNameSlice []string
}

func NewTextApi() *TextApi {
	return &TextApi{
		history: make(map[string]string, historyLen),
	}
}

func (t *TextApi) JsonDiff(req models.ReqJsonDiff) models.BaseResponse {
	if req.Old == "" && t.TempJson == "" {
		return RespErr(models.CodeFail, "old json is empty")
	} else if req.Old == "" {
		req.Old = t.TempJson
	}
	res, err := utils.JsonDiff(req.Old, req.New)
	if err != nil {
		log.ErrorWithFields(map[string]interface{}{
			"old": req.Old,
			"new": req.New,
			"err": err,
		}, "json diff failed")
		return RespErr(models.CodeFail, err.Error())
	}

	return RespSuccess(res)
}

func (t *TextApi) TempStoreJson(name, json string) models.BaseResponse {
	t.TempJson = json
	// 存储历史
	if name != "" {
		if len(t.history) > historyLen {
			delete(t.history, t.histNameSlice[0])
			t.histNameSlice[0] = ""
			t.histNameSlice = t.histNameSlice[1:]
		}

		t.history[name] = json
		t.histNameSlice = append(t.histNameSlice, name)
	}

	return RespSuccess(nil)
}

func (t *TextApi) GetTempJson() models.BaseResponse {
	return RespSuccess(t.TempJson)
}

func (t *TextApi) GenHistoryName(key string) models.BaseResponse {
	if _, exist := t.history[key]; exist {
		for i := 1; i < historyLen*100; i++ {
			nameNumber := t.appendNumber(key, i)
			if _, exist := t.history[nameNumber]; !exist {
				return RespSuccess(t.appendNumber(key, i))
			}
		}
		return RespErr(models.CodeFail, "history name is full")
	}
	return RespSuccess(key)
}

func (t *TextApi) GetHistory() models.BaseResponse {
	return RespSuccess(t.history)
}

func (t *TextApi) appendNumber(name string, n int) string {
	return fmt.Sprintf("%s(%d)", name, n)
}

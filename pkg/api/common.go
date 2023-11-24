package api

import (
	"errors"
	"fmt"
	"runtime/debug"

	"changeme/pkg/log"
	"changeme/pkg/models"
	"changeme/pkg/service"
	"changeme/pkg/utils"
)

var (
	ErrFormatValueFailed = errors.New("format value failed")
)

func RespSuccess(data interface{}) models.BaseResponse {
	// 是否格式化数据
	var (
		code     = models.CodeSuccess
		msg      = models.MsgSuccess
		finalVal = data
	)

	switch val := data.(type) {
	case string:
		if !utils.IsJsonFormat(val) {
			break
		}
		value, err := FormatValue(val)
		if err != nil {
			finalVal = err.Error()
		} else {
			finalVal = value
		}
	case []string:
		if len(val) == 0 || !utils.IsJsonFormat(val[0]) {
			break
		}
		values, err := FormatValues(val)
		if err != nil {
			finalVal = err.Error()
		} else {
			finalVal = values
		}
	case nil:
		finalVal = nil
	default:
		//code = models.CodeUnSupportFormatType
		//msg = fmt.Sprintf("unsupport format type: [%v]", reflect.TypeOf(val))
		//finalVal = msg
		finalVal = data
	}

	return models.BaseResponse{
		Code:    code,
		Message: msg,
		Data:    finalVal,
	}
}

func RespErr(code int, msg string) models.BaseResponse {
	return models.BaseResponse{
		Code:    code,
		Message: msg,
	}
}

func recoverFromPanic() {
	if err := recover(); err != nil {
		log.Errorf("panic: [%v]", err)
		debug.PrintStack()
	}
}

func FormatValue(val string) (string, error) {
	if service.GlobalConfig.JsonFormat {
		unCompact, err := utils.UnCompact(val)
		if err != nil {
			return "", fmt.Errorf("format value failed: %w", err)
		}
		return unCompact, nil
	}
	return val, nil
}

func FormatValues(vals []string) ([]string, error) {
	if !service.GlobalConfig.JsonFormat {
		return vals, nil
	}
	var ret []string
	for _, val := range vals {
		jsonVal, err := FormatValue(val)
		if err != nil {
			return nil, err
		}
		ret = append(ret, jsonVal)
	}
	return ret, nil
}

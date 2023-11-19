package api

import (
	"changeme/pkg/models"
	"changeme/pkg/service"
)

type GlobalConfigApi struct {
	gcSvc *service.GlobalConfigService
}

func NewGlobalConfigApi() *GlobalConfigApi {
	return &GlobalConfigApi{
		gcSvc: service.GetGlobalConfigService(),
	}
}

func (gc *GlobalConfigApi) Get() models.BaseResponse {
	conf, err := gc.gcSvc.GetConfig()
	if err != nil {
		return RespErr(models.CodeFail, err.Error())
	}
	return RespSuccess(conf)
}

func (gc *GlobalConfigApi) Set(req models.GlobalConfigReq) models.BaseResponse {
	if err := gc.gcSvc.Update(&models.GlobalConfig{
		JsonFormat: req.JsonFormat,
	}); err != nil {
		return RespErr(models.CodeFail, err.Error())
	}
	return RespSuccess(nil)
}

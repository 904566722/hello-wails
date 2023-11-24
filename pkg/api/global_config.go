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
		JsonFormat:   req.JsonFormat,
		EtcdEndPoint: req.EtcdEndPoint,
	}); err != nil {
		return RespErr(models.CodeFail, err.Error())
	}

	// 更新之后，同步到内存中
	if err := gc.gcSvc.Sync(); err != nil {
		return RespErr(models.CodeFail, "sync global config failed")
	}

	return RespSuccess(nil)
}

package exceptions

import "net/http"

type ErrorLogic struct {
	//ErrName  string
	ErrCode  int
	HttpCode int
	Message  string
}

// refer to for best practice: https://learn.microsoft.com/en-us/azure/architecture/best-practices/api-design
const (
	DataAlreadyExist  = 10001
	DataNotFound      = 10002
	DataCreateFailed  = 10003
	DataUpdateFailed  = 10004
	DataDeleteFailed  = 10005
	DataGetFailed     = 10006
	DataSaveFailed    = 10007
	DataPublishFailed = 10008
	InvalidValue      = 10009
	InvalidArgument   = 10010
	OutboundFailed    = 10011

	//
	ErpDataAlreadyExist  = 20001
	ErpDataNotFound      = 20002
	ErpDataCreateFailed  = 20003
	ErpDataUpdateFailed  = 20004
	ErpDataDeleteFailed  = 20005
	ErpDataGetFailed     = 20006
	ErpDataSaveFailed    = 20007
	ErpDataPublishFailed = 20008
	ErpInvalidValue      = 20009
	ErpInvalidArgument   = 20010
	ErpOutboundFailed    = 20011

	//OtherError         = 10007
)

var businessLogicReason = map[int]ErrorLogic{
	DataAlreadyExist:  {ErrCode: DataAlreadyExist, HttpCode: http.StatusUnprocessableEntity, Message: "data is already exist"},
	DataNotFound:      {ErrCode: DataNotFound, HttpCode: http.StatusNotFound, Message: "data not found"},
	DataCreateFailed:  {ErrCode: DataCreateFailed, HttpCode: http.StatusUnprocessableEntity, Message: "create data failed"},
	DataUpdateFailed:  {ErrCode: DataUpdateFailed, HttpCode: http.StatusUnprocessableEntity, Message: "update data failed"},
	DataSaveFailed:    {ErrCode: DataSaveFailed, HttpCode: http.StatusUnprocessableEntity, Message: "save data failed"},
	DataDeleteFailed:  {ErrCode: DataDeleteFailed, HttpCode: http.StatusUnprocessableEntity, Message: "delete data failed"},
	DataGetFailed:     {ErrCode: DataGetFailed, HttpCode: http.StatusUnprocessableEntity, Message: "get data failed"},
	DataPublishFailed: {ErrCode: DataPublishFailed, HttpCode: http.StatusInternalServerError, Message: "publish data failed"},
	InvalidValue:      {ErrCode: InvalidValue, HttpCode: http.StatusInternalServerError, Message: "invalid value"},
	InvalidArgument:   {ErrCode: InvalidArgument, HttpCode: http.StatusUnprocessableEntity, Message: "invalid argument"},
	OutboundFailed:    {ErrCode: OutboundFailed, HttpCode: http.StatusInternalServerError, Message: "call outbound failed"},

	//
	ErpDataAlreadyExist:  {ErrCode: DataAlreadyExist, HttpCode: http.StatusInternalServerError, Message: "data is already exist"},
	ErpDataNotFound:      {ErrCode: DataNotFound, HttpCode: http.StatusInternalServerError, Message: "data not found"},
	ErpDataCreateFailed:  {ErrCode: DataCreateFailed, HttpCode: http.StatusInternalServerError, Message: "create data failed"},
	ErpDataUpdateFailed:  {ErrCode: DataUpdateFailed, HttpCode: http.StatusInternalServerError, Message: "update data failed"},
	ErpDataSaveFailed:    {ErrCode: DataSaveFailed, HttpCode: http.StatusInternalServerError, Message: "save data failed"},
	ErpDataDeleteFailed:  {ErrCode: DataDeleteFailed, HttpCode: http.StatusInternalServerError, Message: "delete data failed"},
	ErpDataGetFailed:     {ErrCode: DataGetFailed, HttpCode: http.StatusInternalServerError, Message: "get data failed"},
	ErpDataPublishFailed: {ErrCode: DataPublishFailed, HttpCode: http.StatusInternalServerError, Message: "publish data failed"},
	ErpInvalidValue:      {ErrCode: InvalidValue, HttpCode: http.StatusInternalServerError, Message: "invalid value"},
	ErpInvalidArgument:   {ErrCode: InvalidArgument, HttpCode: http.StatusInternalServerError, Message: "invalid argument"},
	ErpOutboundFailed:    {ErrCode: OutboundFailed, HttpCode: http.StatusInternalServerError, Message: "call outbound failed"},

	//OtherError:         {ErrCode: ABC, HttpCode: http.StatusInternalServerError, Message: "your explanation of error EBL = error business logic"},
}

func BusinessLogicReason(code int) ErrorLogic {
	return businessLogicReason[code]
}

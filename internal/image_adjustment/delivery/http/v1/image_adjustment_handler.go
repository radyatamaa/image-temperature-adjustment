package v1

import (
	"context"
	"errors"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/radyatamaa/image-temperature-adjustment/internal"
	"github.com/radyatamaa/image-temperature-adjustment/internal/domain"
	"github.com/radyatamaa/image-temperature-adjustment/pkg/helper"
	"github.com/radyatamaa/image-temperature-adjustment/pkg/response"
	"github.com/radyatamaa/image-temperature-adjustment/pkg/validator"
	"github.com/radyatamaa/image-temperature-adjustment/pkg/zaplogger"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type ImageAdjustmentHandler struct {
	ZapLogger zaplogger.Logger
	internal.BaseController
	response.ApiResponse
	Usecase domain.ImageAdjustmentUseCase
}

func NewImageAdjustmentHandler(useCase domain.ImageAdjustmentUseCase, zapLogger zaplogger.Logger) {
	pHandler := &ImageAdjustmentHandler{
		ZapLogger: zapLogger,
		Usecase:   useCase,
	}
	beego.Router("/api/v1/image_adjustment/temperature", pHandler, "post:ImageAdjustmentTemperature")
}

func (h *ImageAdjustmentHandler) Prepare() {
	// check user access when needed
	h.SetLangVersion()
}

// ImageAdjustmentTemperature
// @Title ImageAdjustmentTemperature
// @Tags ImageAdjustment
// @Summary ImageAdjustmentTemperature
// @Produce json
// @Param Accept-Language header string false "lang"
// @Success 200 {object} swagger.BaseResponse{errors=[]object,data=object}
// @Failure 400 {object} swagger.BadRequestErrorValidationResponse{errors=[]swagger.ValidationErrors,data=object}
// @Failure 408 {object} swagger.RequestTimeoutResponse{errors=[]object,data=object}
// @Failure 500 {object} swagger.InternalServerErrorResponse{errors=[]object,data=object}
// @Param        file   formData  file    true  "file"
// @Param        adjustment_temperature  formData  string  true  "adjustment_temperature"
// @Param        preview  formData  string  false  "preview = true or false"
// @Router /v1/image_adjustment/temperature [post]
func (h *ImageAdjustmentHandler) ImageAdjustmentTemperature() {
	file, fileHeader, err := h.GetFile("file")
	if err != nil {
		h.Ctx.Input.SetData("stackTrace", h.ZapLogger.SetMessageLog(err))
		h.ResponseError(h.Ctx, http.StatusBadRequest, response.ApiValidationCodeError, response.ErrorCodeText(response.ApiValidationCodeError, h.Locale.Lang), err)
		return
	}

	request := domain.ImageAdjustmentRequest{
		File:                  file,
		FileHeader:            fileHeader,
		AdjustmentTemperature: helper.StringToFloat(h.GetString("adjustment_temperature")),
		Preview: 				h.GetString("preview"),
	}

	if err := validator.Validate.ValidateStruct(&request); err != nil {
		h.Ctx.Input.SetData("stackTrace", h.ZapLogger.SetMessageLog(err))
		h.ResponseError(h.Ctx, http.StatusBadRequest, response.ApiValidationCodeError, response.ErrorCodeText(response.ApiValidationCodeError, h.Locale.Lang), err)
		return
	}

	if err := request.ValidateFile(); err != nil {
		if errors.Is(err, response.ErrRequiredFile) {
			h.ResponseError(h.Ctx, http.StatusBadRequest, response.RequiredFileErrorCode, response.ErrorCodeText(response.RequiredFileErrorCode, h.Locale.Lang), err)
			return
		}
		if errors.Is(err, response.ErrInvalidFormatFileJpeg) {
			h.ResponseError(h.Ctx, http.StatusBadRequest, response.InvalidFormatFileJpegErrorCode, response.ErrorCodeText(response.InvalidFormatFileJpegErrorCode, h.Locale.Lang), err)
			return
		}
		h.Ctx.Input.SetData("stackTrace", h.ZapLogger.SetMessageLog(err))
		h.ResponseError(h.Ctx, http.StatusBadRequest, response.ApiValidationCodeError, response.ErrorCodeText(response.ApiValidationCodeError, h.Locale.Lang), err)
		return
	}

	result, err := h.Usecase.ImageAdjustmentTemperature(h.Ctx, request)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			h.ResponseError(h.Ctx, http.StatusRequestTimeout, response.RequestTimeoutCodeError, response.ErrorCodeText(response.RequestTimeoutCodeError, h.Locale.Lang), err)
			return
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.ResponseError(h.Ctx, http.StatusBadRequest, response.DataNotFoundCodeError, response.ErrorCodeText(response.DataNotFoundCodeError, h.Locale.Lang), err)
			return
		}
		h.ResponseError(h.Ctx, http.StatusInternalServerError, response.ServerErrorCode, response.ErrorCodeText(response.ServerErrorCode, h.Locale.Lang), err)
		return
	}
	if request.Preview == "true" {
		imageData, err := ioutil.ReadFile(result.OutputPathDirImage)
		if err != nil {
			h.ResponseError(h.Ctx, http.StatusInternalServerError, response.ServerErrorCode, response.ErrorCodeText(response.ServerErrorCode, h.Locale.Lang), err)
			return
		}
		h.Ctx.Output.Body(imageData)
	}else {
		h.Ok(h.Ctx, h.Tr("message.success"), result)
	}
	return
}

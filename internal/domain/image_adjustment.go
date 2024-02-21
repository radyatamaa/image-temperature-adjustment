package domain

import (
	beegoContext "github.com/beego/beego/v2/server/web/context"
	"github.com/radyatamaa/image-temperature-adjustment/pkg/response"
	"mime/multipart"
	"net/http"
)

type ImageAdjustmentRequest struct {
	File multipart.File `json:"file"`
	FileHeader *multipart.FileHeader `json:"file_header"`
	AdjustmentTemperature float64 `json:"adjustment_temperature" validate:"required"`
	Preview string `json:"preview"`
}

type ImageAdjustmentResponse struct {
	InputFileImage string `json:"input_file_image"`
	InputPathDirImage string `json:"input_path_dir_image"`
	OutputFileImage string `json:"output_file_image"`
	OutputPathDirImage string `json:"output_path_dir_image"`
}

// ImageAdjustmentUseCase UseCase Interface
type ImageAdjustmentUseCase interface {
	ImageAdjustmentTemperature(beegoCtx *beegoContext.Context, request ImageAdjustmentRequest) (res ImageAdjustmentResponse,err error)
}


func (f *ImageAdjustmentRequest) ValidateFile() error {
	if f.File == nil {
		return response.ErrRequiredFile
	}

	file, err := f.FileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 512) // 512 bytes should be enough to detect the file type
	_, err = file.Read(buffer)
	if err != nil {
		return err
	}

	fileType := http.DetectContentType(buffer)
	if fileType != "image/jpeg" {
		return response.ErrInvalidFormatFileJpeg
	}

	return nil
}

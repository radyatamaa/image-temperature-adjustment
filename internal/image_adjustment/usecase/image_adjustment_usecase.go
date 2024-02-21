package usecase

import (
	"context"
	"fmt"
	beegoContext "github.com/beego/beego/v2/server/web/context"
	"github.com/radyatamaa/image-temperature-adjustment/pkg/helper"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"os"
	"time"

	"github.com/nfnt/resize"
	"github.com/radyatamaa/image-temperature-adjustment/internal/domain"
	"github.com/radyatamaa/image-temperature-adjustment/pkg/zaplogger"
)

type imageAdjustmentUseCase struct {
	zapLogger                  zaplogger.Logger
	contextTimeout             time.Duration
}


func NewImageAdjustmentUseCase(timeout time.Duration,
	zapLogger zaplogger.Logger) domain.ImageAdjustmentUseCase {
	return &imageAdjustmentUseCase{
		contextTimeout:             timeout,
		zapLogger:                  zapLogger,
	}
}

func(i imageAdjustmentUseCase) adjustTemperature(beegoCtx *beegoContext.Context,file io.Reader, adjustment float64) (input,output *string,err error) {
	nameOfFile := helper.RandomString(10)
	outputPath := fmt.Sprintf("external/storage/%s-output.jpg",nameOfFile)
	inputPath := fmt.Sprintf("external/storage/%s-input.jpg",nameOfFile)

	out, err := os.Create(inputPath)
	if err != nil {
		beegoCtx.Input.SetData("stackTrace", i.zapLogger.SetMessageLog(err))
		return nil,nil,err
	}
	defer out.Close()

	// Copy the uploaded file data to the new file
	_, err = io.Copy(out, file)
	if err != nil {
		beegoCtx.Input.SetData("stackTrace", i.zapLogger.SetMessageLog(err))
		return nil,nil,err
	}

	// Open the input image file
	fileOriginal, err := os.Open(inputPath)
	if err != nil {
		beegoCtx.Input.SetData("stackTrace", i.zapLogger.SetMessageLog(err))
		return nil,nil,err
	}
	defer fileOriginal.Close()

	// Decode the input image
	img, _, err := image.Decode(fileOriginal)
	if err != nil {
		beegoCtx.Input.SetData("stackTrace", i.zapLogger.SetMessageLog(err))
		return nil,nil,err
	}

	// Create a new image with the same bounds as the original image
	bounds := img.Bounds()
	adjustedImg := image.NewRGBA(bounds)

	// Iterate over each pixel and adjust its temperature
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			r, g, b, _ := originalColor.RGBA()

			// Adjust the temperature by shifting the color channels
			rAdjust := uint8(float64(r>>8) * adjustment)
			gAdjust := uint8(float64(g>>8) * adjustment)
			bAdjust := uint8(float64(b>>8) * adjustment)

			adjustedColor := color.RGBA{
				R: rAdjust,
				G: gAdjust,
				B: bAdjust,
				A: 255,
			}

			adjustedImg.Set(x, y, adjustedColor)
		}
	}

	// Resize the image to the original dimensions
	resizedImg := resize.Resize(uint(bounds.Dx()), uint(bounds.Dy()), adjustedImg, resize.NearestNeighbor)

	// Create the output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		beegoCtx.Input.SetData("stackTrace", i.zapLogger.SetMessageLog(err))
		return nil,nil,err
	}
	defer outFile.Close()

	// Encode the adjusted image as JPEG
	err = jpeg.Encode(outFile, resizedImg, &jpeg.Options{Quality: 100})
	if err != nil {
		beegoCtx.Input.SetData("stackTrace", i.zapLogger.SetMessageLog(err))
		return nil,nil,err
	}

	return &inputPath,&outputPath,nil
}

func (i imageAdjustmentUseCase) ImageAdjustmentTemperature(beegoCtx *beegoContext.Context, request domain.ImageAdjustmentRequest) (res domain.ImageAdjustmentResponse, err error) {
	ctx, cancel := context.WithTimeout(beegoCtx.Request.Context(), i.contextTimeout)
	defer cancel()
	beegoCtx.Request.WithContext(ctx)

	inputFile,outputFile, err := i.adjustTemperature(beegoCtx,request.File,request.AdjustmentTemperature)
	if err != nil {
		return domain.ImageAdjustmentResponse{},err
	}

	return domain.ImageAdjustmentResponse{
		InputPathDirImage: *inputFile,
		InputFileImage:  fmt.Sprint("http://", beegoCtx.Request.Host ,"/", *inputFile),
		OutputPathDirImage: *outputFile,
		OutputFileImage: fmt.Sprint("http://", beegoCtx.Request.Host ,"/", *outputFile),
	},nil
}
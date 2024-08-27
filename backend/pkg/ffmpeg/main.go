package ffmpeg

import (
	"fmt"
	"log"
	"os"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
	"sade-backend/pkg/utility"
)

func New(audio, video, img string) *FFMpeg {
	return &FFMpeg{
		AudioWatermark: audio,
		VideoWatermark: video,
		ImgWatermark:   img,
	}
}

func (fw *FFMpeg) addVideoWatermark(input, output string) error {
	err := os.MkdirAll("../temp", os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating temp directory: %s", err)
	}

	tempFile, err := os.CreateTemp("../temp", "temp_video*.mp4")
	if err != nil {
		return fmt.Errorf("could not create temp file: %v", err)
	}
	defer func() {
		_ = tempFile.Close()
		err := os.Remove(tempFile.Name())
		if err != nil {
			log.Printf("Failed to delete temporary file: %v", err)
		}
	}()

	tempWatermark, err := os.CreateTemp("../temp", "temp_watermark*.mp4")
	if err != nil {
		return fmt.Errorf("could not create temp watermark file: %v", err)
	}
	defer func() {
		_ = tempWatermark.Close()
		err := os.Remove(tempWatermark.Name())
		if err != nil {
			log.Printf("Failed to delete temporary watermark file: %v", err)
		}
	}()

	tempInput, err := os.CreateTemp("../temp", "temp_input*.mp4")
	if err != nil {
		return fmt.Errorf("could not create temp input file: %v", err)
	}
	defer func() {
		_ = tempInput.Close()
		err := os.Remove(tempInput.Name())
		if err != nil {
			log.Printf("Failed to delete temporary input file: %v", err)
		}
	}()

	err = ffmpeg.Input(fw.VideoWatermark).Filter("scale", ffmpeg.Args{"1280:720"}).
		Output(tempWatermark.Name(), ffmpeg.KwArgs{"r": "30", "c:a": "aac", "b:a": "128k"}).OverWriteOutput().Run()
	if err != nil {
		return fmt.Errorf("could not set frame rate for watermark video: %v", err)
	}

	err = ffmpeg.Input(input).Filter("scale", ffmpeg.Args{"1280:720"}).
		Output(tempInput.Name(), ffmpeg.KwArgs{"r": "30", "c:a": "aac", "b:a": "128k"}).OverWriteOutput().Run()
	if err != nil {
		return fmt.Errorf("could not set frame rate for input video: %v", err)
	}

	err = ffmpeg.Concat([]*ffmpeg.Stream{
		ffmpeg.Input(tempWatermark.Name()),
		ffmpeg.Input(tempInput.Name()),
	}).Output(tempFile.Name(), ffmpeg.KwArgs{"c:a": "aac", "b:a": "128k"}).OverWriteOutput().Run()
	if err != nil {
		return fmt.Errorf("could not concatenate videos: %v", err)
	}

	overlay := ffmpeg.Input(fw.ImgWatermark).Filter("scale", ffmpeg.Args{"64:-1"})
	err = ffmpeg.Filter(
		[]*ffmpeg.Stream{
			ffmpeg.Input(tempFile.Name()),
			overlay,
		}, "overlay", ffmpeg.Args{"10:10", "enable='gte(t,1)'"}).
		Output(output, ffmpeg.KwArgs{"c:a": "copy"}).OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return fmt.Errorf("could not overlay the logo: %v", err)
	}

	return nil
}

func (fw *FFMpeg) addAudioWatermark(input, output string) error {

	err := os.MkdirAll("../temp", os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating temp directory: %s", err)
	}
	tempWatermarkFile, err := os.CreateTemp("../temp", "temp_watermark*.mp3")
	if err != nil {
		return fmt.Errorf("could not create temp watermark file: %v", err)
	}
	defer func() {
		_ = tempWatermarkFile.Close()
		err := os.Remove(tempWatermarkFile.Name())
		if err != nil {
			log.Printf("Failed to delete temporary watermark file: %v", err)
		}
	}()

	tempOutputFile, err := os.CreateTemp("../temp", "temp_output*.mp3")
	if err != nil {
		return fmt.Errorf("could not create temp output file: %v", err)
	}
	defer func() {
		_ = tempOutputFile.Close()
		err := os.Remove(tempOutputFile.Name())
		if err != nil {
			log.Printf("Failed to delete temporary output file: %v", err)
		}
	}()

	if strings.HasSuffix(fw.AudioWatermark, ".wav") {
		log.Println("Converting audio watermark to mp3...")
		err = ffmpeg.Input(fw.AudioWatermark).
			Output(tempWatermarkFile.Name(), ffmpeg.KwArgs{"c:a": "mp3"}).
			OverWriteOutput().Run()
		if err != nil {
			return fmt.Errorf("could not convert audio watermark to mp3: %v", err)
		}
	} else {
		_ = tempWatermarkFile.Close()
		tempWatermarkFile, err = os.Open(fw.AudioWatermark)
		if err != nil {
			return fmt.Errorf("could not open watermark file: %v", err)
		}
	}

	err = ffmpeg.Filter(
		[]*ffmpeg.Stream{
			ffmpeg.Input(input),
			ffmpeg.Input(tempWatermarkFile.Name()),
		}, "amix", ffmpeg.Args{"inputs=2:duration=longest:dropout_transition=2"}).
		Output(tempOutputFile.Name(), ffmpeg.KwArgs{"c:a": "mp3"}).OverWriteOutput().Run()
	if err != nil {
		return fmt.Errorf("could not overlay audio files: %v", err)
	}

	err = os.Rename(tempOutputFile.Name(), output)
	if err != nil {
		return fmt.Errorf("could not move audio file: %v", err)
	}

	return nil
}

func (fw *FFMpeg) addImgWatermark(input, output string, x, y int) error {
	log.Println("Starting to add image watermark...")
	logo := ffmpeg.Input(fw.ImgWatermark).Filter("scale", ffmpeg.Args{"64:-1"})
	err := ffmpeg.Filter(
		[]*ffmpeg.Stream{
			ffmpeg.Input(input),
			logo,
		}, "overlay", ffmpeg.Args{fmt.Sprintf("%d:%d", x, y)}).
		Output(output).OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return fmt.Errorf("could not overlay the logo: %v", err)
	}

	log.Println("Image watermark applied successfully.")
	return nil
}

func (fw *FFMpeg) ApplyWatermark(fileType utility.FileType, input, output string) error {
	var err error
	switch fileType {
	case utility.ImageType:
		err = fw.addImgWatermark(input, output, 10, 10)
	case utility.VideoType:
		err = fw.addVideoWatermark(input, output)
	case utility.AudioType:
		err = fw.addAudioWatermark(input, output)
	default:
		return fmt.Errorf("unsupported file type: %v", fileType)
	}

	if err != nil {
		return fmt.Errorf("cannot apply watermark: %v", err)
	}

	log.Printf("Watermark applied successfully to %s", output)
	return nil
}

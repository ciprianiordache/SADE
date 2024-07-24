package utility

import (
	"path/filepath"
	"strings"
)

type FileType int

const (
	Unknown FileType = iota
	ImageType
	AudioType
	VideoType
)

func CheckFileType(file string) FileType {
	ext := filepath.Ext(file)
	ext = strings.ToLower(ext)

	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
		return ImageType
	case ".mp3", ".wav", ".ogg", ".flac", ".aac", ".m4a":
		return AudioType
	case ".mp4", ".mov", ".avi", ".mkv", ".wmv", ".flv":
		return VideoType
	default:
		return Unknown
	}
}

func DetermineMediaType(filePath string) string {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
		return "image"
	case ".mp3", ".wav", ".ogg", ".flac", ".aac", ".m4a":
		return "audio"
	case ".mp4", ".mov", ".avi", ".mkv", ".wmv", ".flv":
		return "video"
	default:
		return "unsupported file type"
	}
}

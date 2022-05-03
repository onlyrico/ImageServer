package app

import (
	"flag"
)

var (
	Root      = "./"
	Port      = flag.Int("Port", 9420, "Listen Port")
	PageSize  = flag.Int("page", 1000, "Page size")
	Column    = flag.Int("column", 0, "Column")
	ExtCustom = flag.String("ext", "", "File extensions")
)

var (
	DefaultImageExt = []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".tif", ".svg", ".webp", ".ico"}
	DefaultVideoExt = []string{".mp4", ".mkv", ".mov", ".wmv", ".flv", ".avi", ".rmvb", ".mpg", ".mpeg", ".m4v", ".3gp", ".3g2"}
	DefaultAudioExt = []string{".mp3", ".wav", ".wma", ".ogg", ".flac"}
)

var (
	FileExtension = make(map[string]struct{})
)
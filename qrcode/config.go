package qrcode

import (
	"image/color"

	"github.com/skip2/go-qrcode"
)

type Config struct {
	Content         string
	RecoveryLevel   qrcode.RecoveryLevel
	Size            int
	BackgroundColor color.Color
	ForegroundColor color.Color
	FileName        string
}

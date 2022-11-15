package qrcode

import (
	"github.com/skip2/go-qrcode"
)

func Generate(config *Config) error {
	return qrcode.WriteColorFile(config.Content, config.RecoveryLevel, config.Size, config.BackgroundColor, config.ForegroundColor, config.FileName)
}

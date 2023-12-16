package cmd

import "errors"

const (
	commandDir      = ".commands"
	commandCfg      = "config.json"
	commandUrl      = "https://unpkg.com/linux-command/command/%s.md"
	commandDataJson = "https://unpkg.com/linux-command/dist/data.json"
)

const (
	maxCurrency = 6 // 最大并发数
	maxRetry    = 3 // 最大重试次数
)

// ErrCommandNotFound errors
var (
	ErrCommandNotFound = errors.New("command not found")
)

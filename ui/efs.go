package ui

import (
	"embed"
)

//go:embed static/*
var Resources embed.FS

//go:embed html/*
var Templates embed.FS

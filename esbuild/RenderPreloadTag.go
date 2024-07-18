package esbuild

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/intentt/assett/assetpath"
)

func renderAsAttributeValue(ext string) string {
	switch ext {
	case ".css":
		return "style"
	case ".js":
		return "script"
	case ".woff", ".woff2":
		return "font"
	case ".gif", ".jpg", ".jpeg", ".png", ".svg", ".webp":
		return "image"
	case ".mp4", ".webm", ".ogg":
		return "video"
	case ".mp3", ".wav", ".flac":
		return "audio"
	default:
		return "fetch"
	}
}

func renderRelAttribute(ext string) string {
	if ext == ".js" {
		return "modulepreload"
	}

	return "preload"
}

func RenderPreloadTag(
	PathTransformer assetpath.PathTransformer,
	path string,
) template.HTML {
	ext := strings.ToLower(filepath.Ext(path))

	asAttributeValue := renderAsAttributeValue(ext)
	relAttribute := renderRelAttribute(ext)

	var crossorigin string

	if "font" == asAttributeValue && assetpath.IsRemote(path) {
		crossorigin = " crossorigin"
	}

	var asAttribute string

	if "modulepreload" == relAttribute {
		asAttribute = ""
	} else {
		asAttribute = fmt.Sprintf(` as="%s"`, asAttributeValue)
	}

	return template.HTML(fmt.Sprintf(
		"\n"+`<link rel="%s" href="%s"%s%s>`,
		relAttribute,
		assetpath.TransformPath(PathTransformer, path),
		asAttribute,
		crossorigin,
	))
}

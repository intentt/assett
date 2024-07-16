package esbuild

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
)

type RenderingContext struct {
	CreateClientSideAssetPath CreateClientSideAssetPath
	MetafileIndex             *MetafileIndex
	preloadsList              []string
	stylesheets               []string
}

func (self *RenderingContext) Assets() template.HTML {
	assetsHtml := ""

	for _, stylesheet := range self.stylesheets {
		assetsHtml += stylesheet
	}

	return self.Preloads() + template.HTML(assetsHtml)
}

func (self *RenderingContext) Preloads() template.HTML {
	preloadsHtml := ""

	for _, preload := range self.preloadsList {
		ext := strings.ToLower(filepath.Ext(preload))

		asAttribute := self.getAsAttribute(ext)
		relAttribute := self.getRelAttribute(ext)

		var crossorigin string

		if "font" == asAttribute {
			crossorigin = "crossorigin"
		}

		preloadsHtml += fmt.Sprintf(
			`<link %s rel="%s" href="%s" as="%s">`+"\n    ",
			crossorigin,
			relAttribute,
			self.doCreateClientSideAssetPath(preload),
			asAttribute,
		)
	}

	return template.HTML(preloadsHtml)
}

func (self *RenderingContext) Stylesheet(entryPoint string) error {
	imports, err := self.MetafileIndex.GetImports(entryPoint)

	if err != nil {
		return err
	}

	self.preloadsList = append(self.preloadsList, imports...)

	path, err := self.MetafileIndex.GetPath(entryPoint)

	if err != nil {
		return err
	}

	self.stylesheets = append(self.stylesheets, fmt.Sprintf(
		`<link rel="stylesheet" type="text/css" href="%s">`+"\n    ",
		self.doCreateClientSideAssetPath(path),
	))

	return nil
}

func (self *RenderingContext) doCreateClientSideAssetPath(path string) string {
	if IsUrl(path) {
		return path
	}

	return self.CreateClientSideAssetPath(path)
}

func (self *RenderingContext) getAsAttribute(ext string) string {
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

func (self *RenderingContext) getRelAttribute(ext string) string {
	if ext == ".js" {
		return "modulepreload"
	}

	return "preload"
}

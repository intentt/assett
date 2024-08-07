package esbuild

import (
	"fmt"
	"html/template"

	"github.com/intentt/assett/assetpath"
)

type RenderingContext struct {
	MetafileIndex   *MetafileIndex
	PathTransformer assetpath.PathTransformer
	preloadables    []string
	scripts         string
	stylesheets     string
}

func (self *RenderingContext) RenderAssets() template.HTML {
	return template.HTML(self.stylesheets + self.scripts)
}

func (self *RenderingContext) RenderPreloads() template.HTML {
	var preloadsHtml template.HTML

	for _, preload := range self.preloadables {
		preloadsHtml += RenderPreloadTag(self.PathTransformer, preload)
	}

	return template.HTML(preloadsHtml)
}

func (self *RenderingContext) Script(entryPoint string) error {
	indexedOutput, err := self.registerEntryPoint(entryPoint)

	if err != nil {
		return err
	}

	if "" != indexedOutput.Output.CSSBundle {
		err = self.Stylesheet(indexedOutput.Output.CSSBundle)

		if err != nil {
			return err
		}
	}

	self.stylesheets += fmt.Sprintf(
		"\n"+`<script defer type="module" src="%s"></script>`,
		assetpath.TransformPath(self.PathTransformer, indexedOutput.OutputFilename),
	)

	return nil
}

func (self *RenderingContext) Stylesheet(entryPoint string) error {
	indexedOutput, err := self.registerEntryPoint(entryPoint)

	if err != nil {
		return err
	}

	self.stylesheets += fmt.Sprintf(
		"\n"+`<link rel="stylesheet" type="text/css" href="%s">`,
		assetpath.TransformPath(self.PathTransformer, indexedOutput.OutputFilename),
	)

	return nil
}

func (self *RenderingContext) registerEntryPoint(entryPoint string) (*IndexedOutput, error) {
	preloadables, err := self.MetafileIndex.GetPreloadables(entryPoint)

	if err != nil {
		return nil, err
	}

	self.preloadables = append(self.preloadables, preloadables...)

	indexedOutput, err := self.MetafileIndex.GetIndexedOutput(entryPoint)

	if err != nil {
		return nil, err
	}

	return indexedOutput, nil
}

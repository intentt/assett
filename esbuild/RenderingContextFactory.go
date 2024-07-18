package esbuild

import "github.com/intentt/assett/assetpath"

type RenderingContextFactory struct {
	PathTransformer assetpath.PathTransformer
	MetafileIndex   *MetafileIndex
}

func (self *RenderingContextFactory) NewRenderingContext() *RenderingContext {
	return &RenderingContext{
		PathTransformer: self.PathTransformer,
		MetafileIndex:   self.MetafileIndex,
	}
}

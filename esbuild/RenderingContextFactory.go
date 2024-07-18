package esbuild

type RenderingContextFactory struct {
	PathTransformer PathTransformer
	MetafileIndex   *MetafileIndex
}

func (self *RenderingContextFactory) NewRenderingContext() *RenderingContext {
	return &RenderingContext{
		PathTransformer: self.PathTransformer,
		MetafileIndex:   self.MetafileIndex,
	}
}

package esbuild

type RenderingContextFactory struct {
	CreateClientSideAssetPath CreateClientSideAssetPath
	MetafileIndex             *MetafileIndex
}

func (self *RenderingContextFactory) NewRenderingContext() *RenderingContext {
	return &RenderingContext{
		CreateClientSideAssetPath: self.CreateClientSideAssetPath,
		MetafileIndex:             self.MetafileIndex,
	}
}

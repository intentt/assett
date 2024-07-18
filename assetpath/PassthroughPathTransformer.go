package assetpath

type PassthroughPathTransformer struct{}

func (self *PassthroughPathTransformer) TransformLocalPath(path string) string {
	return path
}

func (self *PassthroughPathTransformer) TransformRemotePath(path string) string {
	return path
}

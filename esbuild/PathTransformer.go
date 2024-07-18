package esbuild

type PathTransformer interface {
	TransformLocalPath(path string) string
	TransformRemotePath(path string) string
}

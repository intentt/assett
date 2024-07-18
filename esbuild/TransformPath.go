package esbuild

func TransformPath(pathTransformer PathTransformer, path string) string {
	if IsRemote(path) {
		return pathTransformer.TransformRemotePath(path)
	}

	return pathTransformer.TransformLocalPath(path)
}

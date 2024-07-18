package esbuild

import "fmt"

type MetafileIndex struct {
	EntryPoints            map[string]*IndexedOutput
	EntryPointPreloadables map[string][]string
}

func (self *MetafileIndex) GetPreloadables(entryPoint string) ([]string, error) {
	imports := self.EntryPointPreloadables[entryPoint]

	if imports == nil {
		return nil, fmt.Errorf("entry point not found: %s", entryPoint)
	}

	return imports, nil
}

func (self *MetafileIndex) GetOutputPath(entryPoint string) (string, error) {
	indexedOutput := self.EntryPoints[entryPoint]

	if indexedOutput == nil {
		return "", fmt.Errorf("entry point not found: %s", entryPoint)
	}

	return indexedOutput.OutputFilename, nil
}

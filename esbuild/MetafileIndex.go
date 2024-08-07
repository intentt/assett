package esbuild

import "fmt"

type MetafileIndex struct {
	EntryPoints            map[string]*IndexedOutput
	EntryPointPreloadables map[string][]string
}

func (self *MetafileIndex) GetIndexedOutput(entryPoint string) (*IndexedOutput, error) {
	indexedOutput := self.EntryPoints[entryPoint]

	if indexedOutput == nil {
		return nil, fmt.Errorf("entry point not found: %s", entryPoint)
	}

	return indexedOutput, nil
}

func (self *MetafileIndex) GetPreloadables(entryPoint string) ([]string, error) {
	imports := self.EntryPointPreloadables[entryPoint]

	if imports == nil {
		return nil, fmt.Errorf("entry point not found: %s", entryPoint)
	}

	return imports, nil
}

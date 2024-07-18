package esbuild

import "encoding/json"

func NewMetafileIndex(metafileBytes []byte) (*MetafileIndex, error) {
	var metafile Metafile

	err := json.Unmarshal(metafileBytes, &metafile)

	if err != nil {
		return nil, err
	}

	metafileIndex := &MetafileIndex{
		EntryPoints:            make(map[string]*IndexedOutput),
		EntryPointPreloadables: make(map[string][]string),
	}

	for outputFilename, output := range metafile.Outputs {
		indexedOutput := &IndexedOutput{
			Output:         &output,
			OutputFilename: outputFilename,
		}
		metafileIndex.EntryPoints[output.EntryPoint] = indexedOutput

		preloadables := make([]string, 0, len(indexedOutput.Output.Imports))

		for _, imprt := range indexedOutput.Output.Imports {
			preloadables = append(preloadables, imprt.Path)
		}

		metafileIndex.EntryPointPreloadables[output.EntryPoint] = preloadables
	}

	return metafileIndex, nil
}

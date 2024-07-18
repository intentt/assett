package esbuild

import "os"

func NewMetafileIndexFromFile(filename string) (*MetafileIndex, error) {
	metafileIndexBytes, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return NewMetafileIndex(metafileIndexBytes)
}

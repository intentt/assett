package esbuild

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImports(t *testing.T) {
	metafileIndex, err := NewMetafileIndexFromFile("fixtures/esbuild-meta.json")

	assert.Nil(t, err)

	imports, err := metafileIndex.GetPreloadables("resources/css/page-common.css")

	assert.Nil(t, err)

	assert.Contains(t, imports, "static/test_6D5OPEBZ.svg")
	assert.Contains(t, imports, "https://fonts/font1.woff2")
}

func TestPaths(t *testing.T) {
	metafileIndex, err := NewMetafileIndexFromFile("fixtures/esbuild-meta.json")

	assert.Nil(t, err)

	indexedOutput, err := metafileIndex.GetIndexedOutput("resources/css/page-common.css")

	assert.Nil(t, err)

	assert.Equal(t, indexedOutput.OutputFilename, "static/page-common_DO3RNJ3I.css")
}

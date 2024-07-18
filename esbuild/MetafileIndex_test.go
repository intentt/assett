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
	assert.Contains(t, imports, "https://fonts.gstatic.com/s/notosans/v36/o-0bIpQlx3QUlC5A4PNB6Ryti20_6n1iPHjc5aDdu2ui.woff2")
}

func TestPaths(t *testing.T) {
	metafileIndex, err := NewMetafileIndexFromFile("fixtures/esbuild-meta.json")

	assert.Nil(t, err)

	path, err := metafileIndex.GetOutputPath("resources/css/page-common.css")

	assert.Nil(t, err)

	assert.Equal(t, path, "static/page-common_DO3RNJ3I.css")
}

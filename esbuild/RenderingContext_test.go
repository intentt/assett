package esbuild

import (
	"html/template"
	"testing"

	"github.com/intentt/assett/assetpath"
	"github.com/stretchr/testify/assert"
)

func TestRendering(t *testing.T) {
	metafileIndex, err := NewMetafileIndexFromFile("fixtures/esbuild-meta.json")

	assert.Nil(t, err)

	renderingContextFactory := &RenderingContextFactory{
		MetafileIndex:   metafileIndex,
		PathTransformer: &assetpath.PassthroughPathTransformer{},
	}

	renderingContext := renderingContextFactory.NewRenderingContext()

	assert.NotNil(t, renderingContext)

	renderingContext.Stylesheet("resources/css/page-common.css")

	assert.Equal(t, renderingContext.RenderPreloads(), template.HTML(`
<link rel="preload" href="https://fonts.gstatic.com/s/notosans/v36/o-0bIpQlx3QUlC5A4PNB6Ryti20_6n1iPHjc5aDdu2ui.woff2" as="font" crossorigin>
<link rel="preload" href="https://fonts.gstatic.com/s/notosans/v36/o-0bIpQlx3QUlC5A4PNB6Ryti20_6n1iPHjc5a7duw.woff2" as="font" crossorigin>
<link rel="preload" href="static/test_6D5OPEBZ.svg" as="image">`))

	assert.Equal(t, renderingContext.RenderAssets(), template.HTML(`
<link rel="stylesheet" type="text/css" href="static/page-common_DO3RNJ3I.css">`))
}

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

	err = renderingContext.Stylesheet("resources/css/page-common.css")

	assert.Nil(t, err)

	err = renderingContext.Script("resources/ts/controller_foo.tsx")

	assert.Nil(t, err)

	assert.Equal(
		t,
		template.HTML(`
<link rel="preload" href="https://fonts/font1.woff2" as="font" crossorigin>
<link rel="preload" href="https://fonts/font2.woff2" as="font" crossorigin>
<link rel="preload" href="static/test_6D5OPEBZ.svg" as="image">
<link rel="modulepreload" href="static/chunk-EMZKCXNJ.js">
<link rel="modulepreload" href="static/chunk-PI4ZFSEL.js">
<link rel="preload" href="static/logo_XSTJPNLH.png" as="image">
<link rel="preload" href="https://fonts/font3.woff2" as="font" crossorigin>`),
		renderingContext.RenderPreloads(),
	)

	assert.Equal(
		t,
		template.HTML(`
<link rel="stylesheet" type="text/css" href="static/page-common_DO3RNJ3I.css">
<link rel="stylesheet" type="text/css" href="static/controller_foo_CX2Z63ZH.css">
<script defer type="module" src="static/controller_foo_CTJMZK66.js"></script>`),
		renderingContext.RenderAssets(),
	)
}

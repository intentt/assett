package esbuild

import (
	"html/template"
	"testing"

	"github.com/intentt/assett/assetpath"
	"github.com/stretchr/testify/assert"
)

func TestRenderPreloadTag(t *testing.T) {
	pathTransformer := &assetpath.PassthroughPathTransformer{}

	assert.Equal(
		t,
		RenderPreloadTag(pathTransformer, "myfont.woff2"),
		template.HTML(`
<link rel="preload" href="myfont.woff2" as="font">`),
	)

	assert.Equal(
		t,
		RenderPreloadTag(pathTransformer, "http://example.com/myfont.woff2"),
		template.HTML(`
<link rel="preload" href="http://example.com/myfont.woff2" as="font" crossorigin>`),
	)

	assert.Equal(
		t,
		RenderPreloadTag(pathTransformer, "main.js"),
		template.HTML(`
<link rel="modulepreload" href="main.js">`),
	)
}

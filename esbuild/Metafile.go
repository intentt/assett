package esbuild

type Metafile struct {
	Inputs  map[string]Input  `json:"inputs"`
	Outputs map[string]Output `json:"outputs"`
}

package esbuild

type Output struct {
	Bytes      int                    `json:"bytes"`
	Inputs     map[string]OutputInput `json:"inputs"`
	Imports    []Import               `json:"imports"`
	Exports    []string               `json:"exports"`
	EntryPoint string                 `json:"entryPoint,omitempty"`
	CSSBundle  string                 `json:"cssBundle,omitempty"`
}

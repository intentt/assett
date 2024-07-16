package esbuild

type Import struct {
	Path     string            `json:"path"`
	Kind     string            `json:"kind"`
	External bool              `json:"external,omitempty"`
	Original string            `json:"original,omitempty"`
	With     map[string]string `json:"with,omitempty"`
}

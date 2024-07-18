package esbuild

import "github.com/intentt/netcfg/netcfg"

type ServerConfigurationPrefixTransformer struct{
	ServerConfiguration *netcfg.ServerConfiguration
}

func (self *ServerConfigurationPrefixTransformer) TransformLocalPath(path string) string {
	return self.ServerConfiguration.PrefixedPath(path)
}

func (self *ServerConfigurationPrefixTransformer) TransformRemotePath(path string) string {
	return path
}

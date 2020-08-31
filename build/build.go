package build

import (
	"path"

	"github.com/evanw/esbuild/pkg/api"
)

// Compile executes esbuild to bundle client-side app logic
func Compile(rootPath string) string {
	res := api.Build(api.BuildOptions{
		EntryPoints: []string{path.Join(rootPath, "index.js")},
		Outfile: "app.js",
		Bundle: true,
		Write: false,
		LogLevel: api.LogLevelInfo,
	})
	if len(res.Errors) > 0 {
		return ""
	}
	return string(res.OutputFiles[0].Contents)
}

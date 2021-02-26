package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sort"

	"github.com/ryanuber/columnize"
)

// Exported variables are expected to be set via -ldflags -X options at build-time
var (
	// Application is the application name
	Application string

	// Timestamp is the time at build
	Timestamp string

	// Commit is the commit hash
	Commit string

	// Tag is the most recent tag
	Tag string

	version Version
)

// Version describes the application version and build environment
type Version struct {
	Application string `json:"application"`

	Tag       string `json:"tag"`
	Timestamp string `json:"timestamp"`
	Commit    string `json:"commit"`
	Golang    string `json:"golang"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`

	Modules map[string]string `json:"modules"`
}

func init() {
	version = parse()
}

func parse() Version {
	modules := map[string]string{}
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		for _, req := range buildInfo.Deps {
			path := req.Path
			modules[path] = req.Version
		}
	}

	return Version{
		Application: Application,
		Tag:         Tag,
		Timestamp:   Timestamp,
		Commit:      Commit,
		Golang:      runtime.Version(),
		OS:          runtime.GOOS,
		Arch:        runtime.GOARCH,
		Modules:     modules,
	}
}

// Info returns version information
func Info() Version {
	return version
}

func (v Version) String(modules bool) string {
	cfg := columnize.DefaultConfig()
	cfg.Prefix = " "

	clientVersion := []string{
		fmt.Sprintf("Version:|%s", v.Tag),
		fmt.Sprintf("Built:|%s", v.Timestamp),
		fmt.Sprintf("Golang:|%s", v.Golang),
		fmt.Sprintf("Git commit:|%s", v.Commit),
		fmt.Sprintf("OS/Arch:|%s/%s", v.OS, v.Arch),
	}
	baseVersion := fmt.Sprintf("%s\n%s", v.Application, columnize.Format(clientVersion, cfg))

	if modules {
		moduleVersion := []string{}
		for path, version := range v.Modules {
			moduleVersion = append(moduleVersion, fmt.Sprintf("%s|%s", path, version))
		}
		sort.Strings(moduleVersion)
		baseVersion += fmt.Sprintf("\nPrivate modules\n%s", columnize.Format(moduleVersion, cfg))
	}

	return baseVersion
}

// Short returns short version string
func (v Version) Short() string {
	return fmt.Sprintf("%s built: %s", v.Tag, v.Timestamp)
}

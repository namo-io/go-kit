package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sort"

	"github.com/ryanuber/columnize"
)

var (
	Timestamp string
	Commit    string
	Tag       string
	version   Version
)

type Version struct {
	Tag       string `json:"tag"`
	Timestamp string `json:"timestamp"`
	Commit    string `json:"commit"`
	Golang    string `json:"golang"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`

	Modules map[string]string `json:"modules"`
}

func init() {
	modules := map[string]string{}
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		for _, req := range buildInfo.Deps {
			path := req.Path
			modules[path] = req.Version
		}
	}

	version = Version{
		Tag:       Tag,
		Timestamp: Timestamp,
		Commit:    Commit,
		Golang:    runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		Modules:   modules,
	}
}

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
	baseVersion := columnize.Format(clientVersion, cfg)

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

func (v Version) Short() string {
	return fmt.Sprintf("%s built: %s", v.Tag, v.Timestamp)
}

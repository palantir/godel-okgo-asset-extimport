// Copyright 2016 Palantir Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package integration_test

import (
	"os"
	"path"
	"testing"

	"github.com/nmiyake/pkg/dirs"
	"github.com/nmiyake/pkg/gofiles"
	"github.com/palantir/godel/pkg/products"
	"github.com/palantir/okgo/okgotester"
	"github.com/stretchr/testify/require"
)

const (
	okgoPluginLocator  = "com.palantir.okgo:okgo-plugin:0.3.0"
	okgoPluginResolver = "https://palantir.bintray.com/releases/{{GroupPath}}/{{Product}}/{{Version}}/{{Product}}-{{Version}}-{{OS}}-{{Arch}}.tgz"

	godelYML = `exclude:
  names:
    - "\\..+"
    - "vendor"
  paths:
    - "godel"
`
)

func TestExtimport(t *testing.T) {
	assetPath, err := products.Bin("extimport-asset")
	require.NoError(t, err)

	configFiles := map[string]string{
		"godel/config/godel.yml": godelYML,
		"godel/config/check.yml": "",
	}

	// create temporary GOPATH
	tmpGoPathDir, cleanup, err := dirs.TempDir("", "")
	require.NoError(t, err)
	defer cleanup()

	gopathSrcDir := path.Join(tmpGoPathDir, "src")
	err = os.MkdirAll(gopathSrcDir, 0755)
	require.NoError(t, err)

	_, err = gofiles.Write(gopathSrcDir, []gofiles.GoFileSpec{
		{
			RelPath: "github.com/org/project/bar/bar.go",
			Src:     `package bar`,
		},
	})
	require.NoError(t, err)

	newGoPath := os.Getenv("GOPATH")
	if newGoPath != "" {
		newGoPath += ":"
	}
	newGoPath += tmpGoPathDir
	err = os.Setenv("GOPATH", newGoPath)
	require.NoError(t, err)

	okgotester.RunAssetCheckTest(t,
		okgoPluginLocator, okgoPluginResolver,
		assetPath, "extimport",
		[]okgotester.AssetTestCase{
			{
				Name: "external import",
				Specs: []gofiles.GoFileSpec{
					{
						RelPath: "foo.go",
						Src:     `package main; import _ "github.com/org/project/bar"`,
					},
				},
				ConfigFiles: configFiles,
				WantError:   true,
				WantOutput: `Running extimport...
foo.go:1:22: imports external package github.com/org/project/bar
Finished extimport
`,
			},
			{
				Name: "external import in file from inner directory",
				Specs: []gofiles.GoFileSpec{
					{
						RelPath: "foo.go",
						Src:     `package main; import _ "github.com/org/project/bar"`,
					},
					{
						RelPath: "inner/bar",
					},
				},
				ConfigFiles: configFiles,
				Wd:          "inner",
				WantError:   true,
				WantOutput: `Running extimport...
../foo.go:1:22: imports external package github.com/org/project/bar
Finished extimport
`,
			},
		},
	)
}

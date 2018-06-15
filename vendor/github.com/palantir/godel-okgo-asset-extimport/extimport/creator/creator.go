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

package creator

import (
	"github.com/palantir/okgo/checker"
	"github.com/palantir/okgo/okgo"

	"github.com/palantir/godel-okgo-asset-extimport/extimport"
)

func Extimport() checker.Creator {
	return checker.NewCreator(
		extimport.TypeName,
		extimport.Priority,
		func(cfgYML []byte) (okgo.Checker, error) {
			return checker.NewAmalgomatedChecker(extimport.TypeName,
				checker.ParamPriority(extimport.Priority),
				checker.ParamIncludeProjectDirFlag(),
			), nil
		},
	)
}

/*
Copyright The Guard Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gitlab

import (
	"log"

	"github.com/appscode/go/term"

	"github.com/skratchdot/open-golang/open"
)

func IssueToken() {
	codeURurl := "https://gitlab.com/profile/personal_access_tokens"
	term.Infoln("Gitlab url for personal access tokens:", codeURurl)
	err := open.Start(codeURurl)
	if err != nil {
		log.Fatalln(err)
	}
}

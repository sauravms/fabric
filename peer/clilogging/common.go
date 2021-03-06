/*
Copyright IBM Corp. 2016 All Rights Reserved.

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

package clilogging

import (
	"fmt"

	"github.com/op/go-logging"
	"github.com/spf13/cobra"
)

func checkLoggingCmdParams(cmd *cobra.Command, args []string) error {
	var err error

	// check that at least one parameter is passed in
	if len(args) == 0 {
		return fmt.Errorf("no parameters provided")
	}

	if cmd.Name() == "setlevel" {
		if len(args) == 1 {
			err = fmt.Errorf("no log level provided")
		} else {
			// check that log level is valid. if not, err is set
			_, err = logging.LogLevel(args[1])
			if err != nil {
				err = fmt.Errorf("%s - %s", err, args[1])
			}
		}
	}

	return err
}

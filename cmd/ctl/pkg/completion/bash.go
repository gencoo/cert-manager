/*
Copyright 2021 The cert-manager Authors.

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

package completion

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/cmd/util"
)

func newCmdCompletionBash(ioStreams genericclioptions.IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "bash",
		Short: "Generate cert-manager CLI scripts for a Bash shell",
		Long: `To load completions:
Bash:
  $ source <(kubectl cert-manager completion bash)
  # To load completions for each session, execute once:
  # Linux:
  $ cert-manager completion bash > /etc/bash_completion.d/cert-manager

  # macOS:
  $ cert-manager completion bash > /usr/local/etc/bash_completion.d/cert-manager
`,
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(cmd.Root().GenBashCompletion(ioStreams.Out))
		},
	}
}

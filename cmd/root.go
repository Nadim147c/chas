/*
Copyright © 2026 Ephemeral

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"context"
	"os"
	"strings"

	"github.com/Nadim147c/chas/pkgs/chas"
	"github.com/Nadim147c/fang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chas <...characters>",
	Short: "Filter lines that contain a specific set of characters",
	Long: `chas (char + has) reads lines from standard input and only 
outputs lines that contain all of the characters specified in the arguments. 
It acts as a stream filter for verifying character presence.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		chars := strings.Join(args, "")
		in, out := cmd.InOrStdin(), cmd.OutOrStdout()
		return chas.Search(chars, in, out)
	},
}

func Execute(version string) {
	err := fang.Execute(
		context.Background(),
		rootCmd,
		fang.WithFlagTypes(),
		fang.WithVersion(version),
		fang.WithoutCompletions(),
		fang.WithoutManpage(),
	)
	if err != nil {
		os.Exit(1)
	}
}

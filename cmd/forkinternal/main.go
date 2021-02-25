package main

import (
	"fmt"
	"os"
	"path"

	"github.com/jsonnetmod/tools/pkg/forkinternal"
	"github.com/jsonnetmod/tools/version"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:     "forkinternal <import_path> [<import_path>...]",
		Short:   "fork go internal pkg",
		Version: version.Version,
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("at least one import path")
			}
			cwd, _ := os.Getwd()
			return forkinternal.ForkInternals(path.Join(cwd, "internal/forked"), args...)
		},
	}

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

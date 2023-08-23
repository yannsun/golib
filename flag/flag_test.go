package flag

import (
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestCobra(t *testing.T) {
	var cmdPull = &cobra.Command{
		Use:   "pull [OPTIONS] NAME[:TAG|@DIGEST]",
		Short: "Pull an image or a repository from a registry",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Pull: " + strings.Join(args, " "))
		},
	}

	var rootCmd = &cobra.Command{Use: "docker"}
	rootCmd.AddCommand(cmdPull)
	rootCmd.Execute()
}

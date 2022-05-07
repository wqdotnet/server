package cmd

import (
	"sync"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test demo",
	Long:  `test demo`,
	Run:   exectest,
}

func init() {
	rootCmd.AddCommand(testCmd)
}

type TestObject struct {
	Field1 int
	Field2 int
}

var (
	mp   = map[string]TestObject{}
	lock = sync.RWMutex{}
)

func exectest(cmd *cobra.Command, args []string) {

}

package armet

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "armet",
	Short: "armet serves APIs for viewing cluster resources",
	Long:  "armet serves APIs for viewing cluster resources",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func InitAndExecute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func main() {

}

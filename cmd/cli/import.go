package cli

import (
	"fmt"
	"os"

	"github.com/erminson/gitlab-vars/internal/usecase"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import list of variables (JSON)",
	Long:  `Import list of variables (JSON)`,
	Run: func(cmd *cobra.Command, args []string) {
		projectId := viper.GetInt64("project-id")
		uc := usecase.NewUseCase(Category, projectId, client)

		err := uc.ImportVariablesFromFile(Filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}

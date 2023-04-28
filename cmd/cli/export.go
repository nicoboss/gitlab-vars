package cli

import (
	"fmt"
	"github.com/erminson/gitlab-vars/internal/usecase"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export list of variables (json)",
	Long:  `Export list of variables (json)`,
	Run: func(cmd *cobra.Command, args []string) {
		projectId := viper.GetInt64("project-id")
		uc := usecase.NewUseCase(projectId, client)

		vars, err := uc.ListVariables()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(vars)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}

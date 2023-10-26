package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/erminson/gitlab-vars/internal/usecase"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export list of variables (JSON)",
	Long:  `Export list of variables (JSON)`,
	Run: func(cmd *cobra.Command, args []string) {
		projectId := viper.GetInt64("project-id")
		uc := usecase.NewUseCase(Category, projectId, client)

		vars, err := uc.ListVariables()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(vars)
		f, errCreate := os.Create(Filename)
		if errCreate != nil {
			log.Fatal(errCreate)
		}
		defer f.Close()
		_, errWrite := f.WriteString(vars)
		if errWrite != nil {
			log.Fatal(errWrite)
		}
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}

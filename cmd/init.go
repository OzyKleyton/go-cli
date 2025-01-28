package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [projectName]",
	Short: "Inicializa um novo Projeto Go.",
	Long: `o comando init cria a estrutura inicial do projeto,
				tendo como padrão as pastas cmd, config, internal e arquivos
				como padrões como Dockerfile, docker-compose e env.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		fmt.Printf("Criando o projeto '%s'...\n", projectName)
		createInitialProject(projectName)
	},
}

func createInitialProject(projectName string) {
	basePath := filepath.Join(".", projectName)
	folders := []string{
		"cmd/server",
		"config/db",
		"internal/api",
		"internal/api/handler",
		"internal/api/middleware",
		"internal/api/router",
		"internal/model",
		"internal/repository",
		"internal/service",
	}

	for _, folder := range folders {
		path := filepath.Join(basePath, folder)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			fmt.Printf("Erro ao criar pasta '%s': %v\n", path, err)
		} else {
			fmt.Printf("Criado: %s\n", path)
		}
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}

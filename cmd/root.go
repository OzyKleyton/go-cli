package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-cli",
	Short: "Ferramenta para scaffolding de projetos Go",
	Long: `cli-go é uma ferramenta CLI que facilita a criação de estruturas
	de projetos Golang com padrões específicos.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bem-vindo ao cli-go! Use --help para ver os comandos disponíveis.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [projectName]",
	Short: "Inicializa um novo Projeto Go.",
	Long: `O comando init cria a estrutura inicial do projeto,
tendo como padrão as pastas cmd, config, internal e arquivos
como padrões como Dockerfile, docker-compose e env.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		fmt.Print("Digite o nome do módulo Go (ex: github.com/seu-usuario/seu-projeto): ")
		reader := bufio.NewReader(os.Stdin)
		moduleName, _ := reader.ReadString('\n')
		moduleName = strings.TrimSpace(moduleName)

		fmt.Printf("Criando o projeto '%s' com módulo '%s'...\n", projectName, moduleName)

		createInitialProject(projectName, moduleName)

	},
}

func createInitialProject(projectName, moduleName string) {
	basePath := filepath.Join(".", projectName)

	folders := []string{
		"config/db",
		"internal/model",
		"internal/repository",
		"internal/service",
		"internal/api",
		"internal/api/router",
		"internal/api/handler",
		"cmd/server",
	}

	for _, folder := range folders {
		path := filepath.Join(basePath, folder)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			fmt.Printf("Erro ao criar pasta '%s': %v\n", path, err)
		} else {
			fmt.Printf("Criado: %s\n", path)
		}
	}

	files := map[string]string{
		"go.mod":                              "templates/go.mod.tmpl",
		".env":                                "templates/.env.tmpl",
		"cmd/server/main.go":                  "templates/project/cmd/server/main.go.tmpl",
		"config/config.go":                    "templates/project/config/config.go.tmpl",
		"config/db/db.go":                     "templates/project/config/db/db.go.tmpl",
		"internal/model/response.go":          "templates/project/internal/model/response.go.tmpl",
		"internal/model/user.go":              "templates/project/internal/model/user.go.tmpl",
		"internal/repository/userRepo.go":     "templates/project/internal/repository/userRepo.go.tmpl",
		"internal/service/userService.go":     "templates/project/internal/service/userService.go.tmpl",
		"internal/api/handler/userHandler.go": "templates/project/internal/api/handler/userHandler.go.tmpl",
		"internal/api/router/router.go":       "templates/project/internal/api/router/router.go.tmpl",
		"internal/api/api.go":                 "templates/project/internal/api/api.go.tmpl",
		"docker-entrypoint.sh":                "templates/docker-entrypoint.sh.tmpl",
		"Dockerfile":                          "templates/Dockerfile.tmpl",
		"docker-compose.yaml":                 "templates/docker-compose.yaml.tmpl",
		"makefile":                            "templates/makefile.tmpl",
		".gitignore":                          "templates/.gitignore.tmpl",
	}

	data := struct {
		Module string
	}{
		Module: moduleName,
	}

	for file, templatePath := range files {
		outputPath := filepath.Join(basePath, file)
		generateFileFromTemplate(templatePath, outputPath, data)
	}

	runCommand(basePath, "go", "mod", "tidy")
}

func runCommand(projectPath, command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Dir = projectPath

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Erro ao executar '%s %v': %v\nSaída:\n%s", command, args, err, string(output))
	}

	fmt.Printf("Comando executado com sucesso: %s %v\n", command, args)
}

func generateFileFromTemplate(templatePath, outputPath string, data interface{}) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(fmt.Sprintf("Erro ao carregar template '%s': %v", templatePath, err))
	}

	file, err := os.Create(outputPath)
	if err != nil {
		panic(fmt.Sprintf("Erro ao criar arquivo '%s': %v", outputPath, err))
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		panic(fmt.Sprintf("Erro ao preencher template '%s': %v", outputPath, err))
	}

	fmt.Printf("Arquivo gerado: %s\n", outputPath)
}

func init() {
	rootCmd.AddCommand(initCmd)
}

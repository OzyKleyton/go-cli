package cmd

import (
	"context"
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"text/template"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

//go:embed template/**/**/* template/* template/*/*/*
var templates embed.FS

var initCmd = &cobra.Command{
	Use:   "init [projectName]",
	Short: "Inicializa um novo Projeto Go.",
	Long: `O comando init cria a estrutura inicial do projeto,
tendo como padrão as pastas cmd, config, internal e arquivos
como Dockerfile, docker-compose e .env.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		figure.NewColorFigure("GO-CLI", "doom", "blue", true).Print()
		fmt.Println()

		projectName := args[0]

		var moduleName string
		prompt := &survey.Input{
			Message: "Digite o nome do módulo Go (ex: github.com/seu-usuario/seu-projeto):",
		}
		survey.AskOne(prompt, &moduleName)
		moduleName = strings.TrimSpace(moduleName)

		color.Cyan("\nCriando o projeto '%s' com módulo '%s'\n", projectName, moduleName)

		s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
		s.Color("cyan")
		s.Suffix = " Configurando estrutura do projeto\n"
		s.Start()

		goVersion, err := getInstalledGoVersion()

		if err == nil {
			color.Green("🧰 Go detectado no sistema: %s", goVersion)
		} else {
			color.Yellow("⚠️  Não foi possível detectar `go` no PATH: %v", err)
			color.Cyan("Versão do binário (compilado): %s", getRuntimeGoVersion())
		}

		createInitialProject(projectName, moduleName, goVersion)

		s.Stop()

		color.Green("\n✅ Projeto '%s' criado com sucesso!", projectName)
		color.Yellow("Para começar, rode os seguintes comandos:\n")
		fmt.Println(color.CyanString("  cd %s", projectName))
		fmt.Println(color.CyanString("  make up"))
		fmt.Println(color.CyanString("  make start"))
	},
}

func getInstalledGoVersion() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "version")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`go([0-9]+(?:\.[0-9]+){1,2})`)
	m := re.FindStringSubmatch(string(out))
	if len(m) >= 2 {
		return m[1], nil
	}
	return "", fmt.Errorf("não foi possível parsear a saída: %q", strings.TrimSpace(string(out)))
}

func getRuntimeGoVersion() string {
	return strings.TrimPrefix(runtime.Version(), "go")
}

func createInitialProject(projectName, moduleName, goVersion string) {
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
			color.Red("❌ Erro ao criar pasta '%s': %v\n", path, err)
		} else {
			color.Green("📁 Criado: %s", path)
		}
	}

	data := struct {
		Module    string
		GoVersion string
	}{
		Module:    moduleName,
		GoVersion: goVersion,
	}

	files := map[string]string{
		"go.mod":                              "template/go.mod.tmpl",
		".env":                                "template/.env.tmpl",
		"cmd/server/main.go":                  "template/cmd/server/main.go.tmpl",
		"config/config.go":                    "template/config/config.go.tmpl",
		"config/db/db.go":                     "template/config/db/db.go.tmpl",
		"internal/model/response.go":          "template/internal/model/response.go.tmpl",
		"internal/model/user.go":              "template/internal/model/user.go.tmpl",
		"internal/repository/userRepo.go":     "template/internal/repository/userRepo.go.tmpl",
		"internal/service/userService.go":     "template/internal/service/userService.go.tmpl",
		"internal/api/handler/userHandler.go": "template/internal/api/handler/userHandler.go.tmpl",
		"internal/api/router/router.go":       "template/internal/api/router/router.go.tmpl",
		"internal/api/api.go":                 "template/internal/api/api.go.tmpl",
		"docker-entrypoint.sh":                "template/docker-entrypoint.sh.tmpl",
		"Dockerfile":                          "template/Dockerfile.tmpl",
		"docker-compose.yaml":                 "template/docker-compose.yaml.tmpl",
		"makefile":                            "template/makefile.tmpl",
		".gitignore":                          "template/.gitignore.tmpl",
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
		color.Red("❌ Erro ao executar '%s %v': %v\nSaída:\n%s", command, args, err, string(output))
		return
	}

	color.Green("✅ Comando executado com sucesso: %s %v", command, args)
}

func generateFileFromTemplate(templatePath, outputPath string, data interface{}) {
	tmplContent, err := templates.ReadFile(templatePath)
	if err != nil {
		color.Red("❌ Erro ao carregar template '%s': %v", templatePath, err)
		return
	}

	tmpl, err := template.New("template").Parse(string(tmplContent))
	if err != nil {
		color.Red("❌ Erro ao parsear template '%s': %v", templatePath, err)
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		color.Red("❌ Erro ao criar arquivo '%s': %v", outputPath, err)
		return
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		color.Red("❌ Erro ao preencher template '%s': %v", outputPath, err)
		return
	}

	color.Green("📝 Arquivo gerado: %s", outputPath)
}

func init() {
	rootCmd.AddCommand(initCmd)
}

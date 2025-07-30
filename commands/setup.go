package commands

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"task-runner/utils"
)

func SetupPath() {
	binaryPath, err := os.Executable()
	if err != nil {
		utils.LogError("Não foi possível obter o caminho do executável.")
		return
	}
	binaryDir := filepath.Dir(binaryPath)
	binaryName := filepath.Base(binaryPath)

	aliasName := "trun"

	// Verifica se já está no PATH
	if isAlreadyPresent(binaryName, binaryDir) {
		utils.LogSuccess("✅ O TaskRunner já está no PATH!")
		return
	}

	if runtime.GOOS == "windows" {
		addToPathWindows(binaryDir)
	} else {
		addToPathUnix(binaryDir, aliasName, binaryName)
	}
}

func addToPathWindows(binaryDir string) {
	cmd := exec.Command("setx", "PATH", fmt.Sprintf("%s;%s", os.Getenv("PATH"), binaryDir))
	err := cmd.Run()
	if err != nil {
		utils.LogError("❌ Erro ao adicionar no PATH no Windows.")
	} else {
		utils.LogSuccess("✅ Adicionado ao PATH no Windows com sucesso.")
		utils.LogInfo("🔄 Reinicie o terminal para aplicar.")
	}
}

func addToPathUnix(binaryDir, aliasName, binaryName string) {
	profile := getShellProfile()

	exportLine := fmt.Sprintf("\n# Added by TaskRunner\nexport PATH=\"$PATH:%s\"\n", binaryDir)
	aliasLine := fmt.Sprintf("alias %s=\"%s\"\n", aliasName, binaryName)

	contentBytes, _ := os.ReadFile(profile)
	content := string(contentBytes)

	// Verifica se já existe
	if strings.Contains(content, exportLine) && strings.Contains(content, aliasLine) {
		utils.LogSuccess(fmt.Sprintf("✅ PATH e alias '%s' já estão configurados no %s.", aliasName, profile))
		return
	}

	file, err := os.OpenFile(profile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		utils.LogError("❌ Falha ao abrir arquivo de profile.")
		return
	}
	defer file.Close()

	_, err = file.WriteString(exportLine + aliasLine)
	if err != nil {
		utils.LogError("❌ Falha ao escrever no arquivo de profile.")
		return
	}

	utils.LogSuccess(fmt.Sprintf("✅ PATH e alias '%s' adicionados no arquivo %s.", aliasName, profile))
	utils.LogInfo(fmt.Sprintf("🔄 Reinicie o terminal ou execute 'source %s' para aplicar.", profile))
}

func isAlreadyPresent(command, binaryDir string) bool {
	var checkCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		checkCmd = exec.Command("where", command)
	} else {
		checkCmd = exec.Command("command", "-v", command)
	}

	if checkCmd.Run() == nil {
		return true
	}

	profile := getShellProfile()

	exportLine := fmt.Sprintf("export PATH=\"$PATH:%s\"", binaryDir)

	contentBytes, err := os.ReadFile(profile)
	if err != nil {
		return false
	}

	content := string(contentBytes)
	return strings.Contains(content, exportLine)
}

func getShellProfile() string {
	usr, _ := user.Current()
	homeDir := usr.HomeDir
	shell := os.Getenv("SHELL")

	switch {
	case strings.Contains(shell, "zsh"):
		return filepath.Join(homeDir, ".zshrc")
	case strings.Contains(shell, "bash"):
		return filepath.Join(homeDir, ".bashrc")
	default:
		return filepath.Join(homeDir, ".profile")
	}
}

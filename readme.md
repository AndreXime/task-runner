# TaskRunner

TaskRunner é um CLI desenvolvido em Go para gerenciar seus comandos favoritos no terminal. Salve, execute e organize comandos que você utiliza com frequência, de forma simples e rápida.

## Interface

<img width="1400" height="408" alt="image" src="https://github.com/user-attachments/assets/07cc9503-9b32-467a-ab8e-be234be111ea" />

## Funcionalidades

-   Salvar comandos favoritos com nome customizado.
-   Executar comandos rapidamente pelo nome.
-   Suporte a placeholders como $1 para parâmetros dinâmicos.
-   Adicionar o TaskRunner automaticamente ao seu $PATH.
-   Armazena seus comandos de forma persistente em um arquivo JSON na sua home.

## Instalação

### Download do binário

Acesse a página de [Releases](https://github.com/AndreXime/task-runner/releases) e
baixe o binário para seu sistema operacional:

```bash
# Dê permissão de execução (Linux/macOS):
chmod +x task-runner-linux-amd64
# Mova para uma pasta segura
mv taskrunner-linux-amd64 ~/.local/
# Adicione ao path automaticamente
./taskrunner-linux-amd64 setup
# Verifique o funcionamento
trun
```

### Compile você mesmo

```bash
# Entre na pasta do repositorio
git clone https://github.com/AndreXime/task-runner.git
cd task-runner
# Antes de compilar tenha Go instalado no seu PC
go build -o task-runner
# Adicione na Path
./task-runner setup
# Verifique o funcionamento
trun
```

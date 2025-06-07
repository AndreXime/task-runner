## Tasker CLI

Tasker é um CLI (Command Line Interface) desenvolvido em Go para gerenciar seus comandos favoritos no terminal. Salve, execute e organize comandos que você utiliza com frequência, de forma simples e rápida.

## Funcionalidades

- Salvar comandos favoritos com nome customizado.
- Executar comandos rapidamente pelo nome.
- Suporte a placeholders como $1 para parâmetros dinâmicos.
- Adicionar o Tasker automaticamente ao seu $PATH.
- Armazena seus comandos de forma persistente em um arquivo JSON na sua home.

### Exemplos dos comandos

```bash
# Comando para listar processos que usam certa porta
tasker add list-port 'lsof -i :$1'
tasker list
tasker run list-port 3000
tasker remove list-port
```

## Instalação

### Download do binário

Acesse a página de [Releases](https://github.com/AndreXime/tasker-cli/releases) e
baixe o binário para seu sistema operacional:

```bash
# Dê permissão de execução (Linux/macOS):
chmod +x tasker-linux-amd64
# Mova para uma pasta segura
mv tasker-linux-amd64 ~/.local/
# Adicione ao path automaticamente
./tasker setup
# Verifique o funcionamento
tasker
```

### Compile você mesmo

```bash
# Entre na pasta do repositorio
git clone https://github.com/AndreXime/tasker-cli.git
cd tasker-cli
# Antes de compilar tenha Go instalado no seu PC
go build -o tasker
# Adicione na Path
./tasker setup
# Verifique o funcionamento
tasker
```

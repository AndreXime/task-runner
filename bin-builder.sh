#!/bin/bash

set -e

APP_NAME="tasker"
OUTPUT_DIR="bin"

# Cria o diretório se não existir
mkdir -p "$OUTPUT_DIR"

# Lista de destinos: OS/ARCH/SUFIXO
targets=(
  "darwin/amd64"
  "darwin/arm64"
  "linux/amd64"
  "linux/arm64"
  "windows/amd64.exe"
  "windows/arm64.exe"
)

echo "🔨 Compilando binários para múltiplas plataformas..."

for target in "${targets[@]}"; do
  IFS="/" read -r GOOS GOARCH <<< "$(echo "$target" | sed 's/\.exe//')"
  SUFFIX=$(echo "$target" | sed "s/\//-/")

  echo "➡️  $GOOS/$GOARCH → $APP_NAME-$SUFFIX"

  GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 \
    go build -ldflags="-s -w" -o "$OUTPUT_DIR/$APP_NAME-$SUFFIX"
done

echo "✅ Binários salvos em: $OUTPUT_DIR/"

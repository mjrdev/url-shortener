#!/usr/bin/env bash
set -euo pipefail

usage() {
  echo "Usage: $(basename "$0") NAME"
  exit 1
}

if [[ $# -ne 1 ]]; then
  usage
fi

name="$1"
slug="$(echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' ' '_' | tr -cd '[:alnum:]_')"
if [[ -z "$slug" ]]; then
  echo "Error: migration name must contain alphanumeric characters."
  exit 1
fi

timestamp="$(date +%s)"
filename="migrations/${timestamp}_${slug}.go"

mkdir -p "$(dirname "$filename")"

cp ./templates/migration-template.go "$filename"
sed -i "s/ID: \"---\"/ID: \"${timestamp}\"/" "$filename"

echo "Created migration: $filename"

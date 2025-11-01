#!/bin/bash

if [ -z "$1" ]; then
  echo "Использование: $0 <путь_к_директории>"
  exit 1
fi

echo "Текущее значение PATH:"
echo "$PATH"

export PATH="$PATH:$1"

echo
echo "Новое значение PATH:"
echo "$PATH"


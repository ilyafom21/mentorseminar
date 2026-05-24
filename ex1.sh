#!/bin/bash

if [ -z "$1" ]; then
  echo "Использование: $0 <имя_файла>"
  exit 1
fi

target="$1"

echo "Список файлов и каталогов"
for item in *; do
  if [ -f "$item" ]; then
    type="файл"
  elif [ -d "$item" ]; then
    type="каталог"
  elif [ -L "$item" ]; then
    type="ссылка"
  else
    type="другое"
  fi
  echo "$item — $type"
done

echo
echo "Проверка наличия файла"
if [ -e "$target" ]; then
  echo "Файл '$target' найден."
else
  echo "Файл '$target' не найден."
fi

echo
echo "Информация о файлах"
for item in *; do
  perms=$(stat -f "%A" "$item" 2>/dev/null)
  echo "$item — права доступа: $perms"
done

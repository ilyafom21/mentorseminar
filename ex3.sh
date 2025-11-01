#!/bin/bash

read -p "Введите число: " num

if [ "$num" -gt 0 ]; then
  echo "Число положительное."
elif [ "$num" -lt 0 ]; then
  echo "Число отрицательное."
else
  echo "Это ноль."
fi

if [ "$num" -gt 0 ]; then
  i=1
  echo "Подсчет от 1 до $num:"
  while [ $i -le $num ]; do
    echo $i
    i=$((i + 1))
  done
fi


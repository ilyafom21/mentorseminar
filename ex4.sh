#!/bin/bash

greet() {
  echo "Hello, $1"
}

sum() {
  echo $(($1 + $2))
}

read -p "Введите имя: " name
greet "$name"

read -p "Введите первое число: " a
read -p "Введите второе число: " b
result=$(sum $a $b)
echo "Сумма чисел: $result"


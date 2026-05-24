#!/bin/bash

echo "Запуск трёх процессов sleep в фоне"
sleep 3 &
sleep 6 &
sleep 9 &

jobs


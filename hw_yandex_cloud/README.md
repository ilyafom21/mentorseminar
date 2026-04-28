# Yandex Cloud Homework

## Описание

В рамках домашнего задания был развернут кластер Yandex Data Processing в Yandex Cloud и настроен Apache Zeppelin.

В Zeppelin реализована витрина `mart_city_top_products` с использованием PySpark.

## Кластер Yandex Data Processing

<img src="./screenshots/yandex_cloud.png" width="900"/>

## Реализовано

- создание DataFrame users, orders, products
- вычисление revenue = qty * price
- join таблиц
- агрегация данных
- Window Functions
- Top-2 товаров по revenue_sum для каждого города
- сохранение в Parquet
- чтение результата из HDFS

## Результат выполнения

<img src="./screenshots/zeppelin.png" width="900"/>
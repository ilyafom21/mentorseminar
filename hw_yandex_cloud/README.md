# Yandex Cloud Homework

## Описание

В рамках домашнего задания был развернут кластер Yandex Data Processing в Yandex Cloud и настроен Apache Zeppelin.

В Zeppelin реализована витрина `mart_city_top_products` с использованием PySpark.

## Кластер Yandex Data Processing

<img width="1091" height="923" alt="yandex_cloud" src="https://github.com/user-attachments/assets/1330ed62-6537-4c0a-a8ad-30c9b792ebb6" />



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

<img width="510" height="866" alt="zeppelin" src="https://github.com/user-attachments/assets/db7ef422-cf42-475d-887f-897705ac90da" />


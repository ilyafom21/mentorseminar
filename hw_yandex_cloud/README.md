# Yandex Cloud Homework

## Описание

В рамках домашнего задания был развернут кластер Yandex Data Processing в Yandex Cloud и настроен Apache Zeppelin.

В Zeppelin реализована витрина `mart_city_top_products` с использованием PySpark.

## Кластер Yandex Data Processing

<img width="1091" height="923" alt="image" src="https://github.com/user-attachments/assets/89c5efd9-d7c1-4bb4-983a-90ff792e9410" />


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

<img width="510" height="866" alt="zeppelin" src="https://github.com/user-attachments/assets/14d6323f-51cf-451c-a44c-e95c6c9ee785" />

# IndustrialTechnology

Практические работы по дисциплине Технологии индустриального программирования  
Иванов Вячеслав, ЭФМО-01-24
## Практика 1. 
# Схема отношений сущностей базы данных ресторана
![Схема БД](https://raw.githubusercontent.com/tr1kK/IndustrialTechnology/refs/heads/prac1/%D1%81%D1%85%D0%B5%D0%BC%D0%B0-%D0%B1%D0%B4-%D1%80%D0%B5%D1%81%D1%82%D0%BE%D1%80%D0%B0%D0%BD%D0%B0_1.png)
# Описание сущностей
* customers — хранит информацию о клиентах ресторана.
* customer_addresses — хранит информацию об адресах клиентов ресторана, что полезно для доставки или дополнительной информации о клиентах.
* orders — содержит информацию о заказах, сделанных клиентами.
* menu_items — хранит информацию о блюдах и продуктах, предлагаемых в меню.
* order_items — хранит информацию о блюдах, заказанных в рамках каждого заказа.
* employees — хранит информацию о сотрудниках ресторана.
* order_assignments — связывает заказы и сотрудников, которые их обрабатывают.
* ingredients — хранит информацию об ингредиентах, используемых для приготовления блюд.
* menu_item_ingredients — связывает пункты меню и ингредиенты, используемые для их приготовления.
* suppliers — хранит информацию о поставщиках ингредиентов для ресторана.

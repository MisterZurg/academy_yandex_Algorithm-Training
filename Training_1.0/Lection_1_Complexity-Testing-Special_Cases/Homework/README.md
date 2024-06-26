# Яндекс. Тренировки по алгоритмам июнь 2021, занятие 1

> [!TIP]
> [РЕБЯТА, ДОМАШНЕЕ ЗАДАНИЕ СЛУШАЕМ!](https://contest.yandex.ru/contest/27393)

## ЗАДачи
- [A. Кондиционер](#A-Кондиционер)
- [B. Треугольник](#B-Треугольник)
- C. Телефонные номера
- D. Уравнение с корнем
- E. Скорая помощь
- F. Расстановка ноутбуков
- G. Детали
- H. Метро
- I. Узник замка Иф
- J. Система линейных уравнений - 2

## A. Кондиционер
В офисе, где работает программист Петр, установили кондиционер нового типа. Этот кондиционер отличается особой простотой в управлении. У кондиционера есть всего лишь два управляемых параметра: желаемая температура и режим работы.

Кондиционер может работать в следующих четырех режимах:

«freeze» — охлаждение. В этом режиме кондиционер может только уменьшать температуру. Если температура в комнате и так не больше желаемой, то он выключается.

«heat» — нагрев. В этом режиме кондиционер может только увеличивать температуру. Если температура в комнате и так не меньше желаемой, то он выключается.

«auto» — автоматический режим. В этом режиме кондиционер может как увеличивать, так и уменьшать температуру в комнате до желаемой.

«fan» — вентиляция. В этом режиме кондиционер осуществляет только вентиляцию воздуха и не изменяет температуру в комнате.

Кондиционер достаточно мощный, поэтому при настройке на правильный режим работы он за час доводит температуру в комнате до желаемой.

Требуется написать программу, которая по заданной температуре в комнате troom, установленным на кондиционере желаемой температуре tcond и режиму работы определяет температуру, которая установится в комнате через час.

### Формат ввода
Первая строка входного файла содержит два целых числа troom, и tcond, разделенных ровно одним пробелом $(–50 ≤ t_room ≤ 50, –50 ≤ t_cond ≤ 50)$.

Вторая строка содержит одно слово, записанное строчными буквами латинского алфавита — режим работы кондиционера.

### Формат вывода
Выходной файл должен содержать одно целое число — температуру, которая установится в комнате через час.

### Пример 1

Ввод

```
10 20
heat
```
Вывод
```
20
```

### Пример 2

Ввод
```
10 20
freeze
```
Вывод
```
10
```

### Примечания
В первом примере кондиционер находится в режиме нагрева. Через час он нагреет комнату до желаемой температуры в $20$ градусов.

Во втором примере кондиционер находится в режиме охлаждения. Поскольку температура в комнате ниже, чем желаемая, кондиционер самостоятельно выключается и температура в комнате не поменяется.


## B. Треугольник
Даны три натуральных числа. Возможно ли построить треугольник с такими сторонами. Если это возможно, выведите строку YES, иначе выведите строку NO.

Треугольник — это три точки, не лежащие на одной прямой.
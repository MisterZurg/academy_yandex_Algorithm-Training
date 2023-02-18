# Конспект Лекции 6: «Бинарный поиск»

## Что такое бинарный поиск

### Термины Михаила Густокашина

* **Левый бинпоиск** — первое подходящее значение

 ```python
# Всё плохо ... Всё хорошо
def lbinsearch(l, r, check, checkparams):
    while l < r:
        m = (l + r) // 2
        if check(m, checkparams):
            r = m
        else:
            l = m + 1
    return l
```

* **Правый бинпоиск** — последнее подходящее значение

```python
# Всё хорошо ... Всё плохо
def rbinsearch(l, r, check, checkparams):
    while l < r:
        m = (l + r + 1) // 2
        if check(m, checkparams):
            l = m
        else:
            r = m - 1
    return l
```

## Задача 1

> В управляющий совет школы входят родители, учителя и учащиеся школы,
> причем родителей должно быть не менее одной трети от общего числа членов совета.
> В настоящий момент в совет входит ```N``` человек, из них ```К``` родителей
>
> Определите, сколько родителей нужно дополнительно ввести в совет,
> чтобы их число стало составлять не менее трети от числа членов совета.

### Решение

> Будем искать минимальное количество родителей, которых нужно добавить,
> бинарным поиском. Не забываем, что **новые родители добавляются к общему числу членов совета**
> Не используем деление. ```L = 0, R = N```

 ```python
def lbinsearch(l, r, check, checkparams):
    while l < r:
        m = (l + r) // 2
        if check(m, checkparams):
            r = m
        else:
            l = m + 1
    return l


def checkendownment(m, params):
    n, k = params
    return (k + m) * 3 >= n + m
```

## Задача 2

> Юра решил подготовиться к собеседованию в Яндекс.
> Он выбрал на сайте **leetcode** ```N``` задач.
> В первый день Юра решил ```K``` задач,
> а в каждый следующий день Юра решал на одну задачу больше,
> чем в предыдущий день
>
> Определите, сколько дней уйдет у Юры на подготовку к собеседованию.

### Решение

> Будем искать минимальное количество дней, достаточное для решения не менее
> ```N``` задач, бинарным поиском.
> Нам понадобится формула арифметической прогрессии ```L = 0, R = N```
> ```(K, K + 1, ..., K + M - 1) = (2K + M - 1) * M / 2```

 ```python
def lbinsearch(l, r, check, checkparams):
    while l < r:
        m = (l + r) // 2
        if check(m, checkparams):
            r = m
        else:
            l = m + 1
    return l


def checkproblemcount(days, params):
    n, k = params
    return (k + +(k + days - 1)) * days // 2 >= n
```

## Задача 3

> Михаил читает лекции по алгоритмам. За кадром стоит доска размером ```W * Н```
> сантиметров. Михаилу нужно разместить на доске ```N``` квадратных стикеров со шпаргалками,
> при этом длина стороны стикера в сантиметрах должна быть целым числом
>
> Определите максимальную длину стороны стикера, чтобы все стикеры поместились на доске.

### Решение

> Будем искать максимальную сторону стикера бинарным поиском

```python
def rbinsearch(l, r, check, checkparams):
    while l != r:
        m = (l + r + 1) // 2
        if check(m, checkparams):
            l = m
        else:
            r = m - 1


def checkstickers(size, params):
    n, w, h = params
    return (w // size) * (h // size) >= n
```

## Задача 4

> Задана отсортированная по неубыванию последовательность из```N```чисел и число ```X```
>
> Необходимо определить индекс первого числа в последовательности, которое больше либо равно ```X```.
> Если такого числа нет, то вернуть число ```N```.

### Решение

> Воспользуемся левым бинпоиском для поиска первого подходящего числа.
> В случае, если бинпоиск сошелся к числу, меньшему ```X``` вернём ```N```

 ```python
def checkisge(index, params):
    seq, x = params
    return seq[index] >= x


def findfirstge(seq, x):
    ans = lbinsearch(0, len(seq) - 1, checkisge, (seq, x))
    if seq[ans] < x:
        return len(seq)
    return ans


# Кто забыл
def lbinsearch(l, r, check, checkparams):
    while l < r:
        m = (l + r) // 2
        if check(m, checkparams):
            r = m
        else:
            l = m + 1
    return l
```

## Задача 5

> Задана отсортированная по неубыванию последовательность из```N```чисел и число ```X```
>
> Необходимо определить сколько раз число ```X``` входит в последовательность.

### Решение

> Найдем одним левым бинпоиском правое число большее либо равное ```X```,
> а вторым левым бинпоиском — число строго большее ```X```.
> Разность индексов и будет количеством вхождений

```python
def checkisgt(index, params):
    seq, x = params
    return seq[index] > x


def checkisge(index, params):
    seq, x = params
    return seq[index] >= x


def findfirst(seq, x, check):
    ans = lbinsearch(0, len(seq) - 1, check, (seq, x))
    if not check(ans, (seq, x)):
        return len(seq)
    return ans


def countx(seq, x):
    indexgt = findfirst(seq, x, checkisgt)
    indexge = findfirst(seq, x, checkisge)
    return indexgt - indexge
    
    
# Кто забыл 2
def lbinsearch(l, r, check, checkparams):
    while l < r:
        m = (l + r) // 2
        if check(m, checkparams):
            r = m
        else:
            l = m + 1
    return l
```
## Задача 6 *из жизни*
> Задана процентная ставка по кредиту (```X%``` годовых), 
> срок кредитования (```N``` месяцев) и сумма кредита (```M``` рублей)
> 
> Необходимо рассчитать размер аннуетиного платежа ежемесячного платежа

### Решение подзадачи о ежемесячном проценте

> Ежемесячный процент не равен ```X/12```! Подберем его бинпоиском.
```py
def checkmonthlypaid(mperc, yperc):
    msum = 1 + mperc / 100
    ysum = 1 + yperc / 100
    return msum ** 12 >= ysum

    
def fbinsearch(l, r, eps, check, checkparams):
    while l + eps < r:
        m = (l + r) / 2
        if check(m, checkparams):
            r = m
        else:
            l = m + 1
    return l


x = 12
eps = 0.0001
mperc = fbinsearch(0, x, eps, checkmonthlypaid, x)
```

### Решение задачи о размере платежа
```py
def checkcredit(mpay, params):
    periods, creditsum, mperc = params
    for i in range(periods):
        percpay = creditsum * (mperc / 100)
        creditsum  -= mpay - percpay
    return creditsum <= 0


monthlypay = fbinsearch(0, m, eps, checkcredit, (n, m, mperc))
print(monthlypay)
```

## Задача 7 *про Тернанарный поиск (которого нет)*
> Велосипедисты, участвующие в шоссейной гонке, в некоторый момент времени,
> который называется начальным, оказались в точках,
> удаленных от места старта на ```x1, x2, ..., хn``` метров (```n``` - общее количество велосипедистов, не превосходит 100 000).
> Каждый велосипедист двигается со своей постоянной скоростью ```v1, v2, ..., vn``` метров в секунду.
> Все велосипедисты двигаются в одну и ту же сторону.
> Репортер, освещающий ход соревнований, хочет определить момент времени,
> в который расстояние между лидирующим в гонке велосипедистом и замыкающим гонку велосипедистом станет минимальным, 
> чтобы с вертолета сфотографировать сразу всех участников велогонки
> 
> Необходимо найти момент времени, когда расстояние станет минимальным.

### Решение
> Определим функцию ```dist(t)```, которая будет за ```O(N)``` определять расстояние между лидером
> и замыкающим в момент времени ```t```. Если ```dist(t + ε) > dist(t)```,
> то функция растёт и надо сдвинуть левую границу поиска, иначе — правую
```python
def dist(t, params):
    x, v = params
    minpos = maxpos = x[0] + v[0] * t
    for i in range(1,len(x)):
        nowpos = x[i] + v[i] * t
        minpos = min(minpos, nowpos)
        maxpos = max(maxpos, nowpos)
    return maxpos - minpos

def checkasc(t, eps, params):
    return dist(t + eps, params) >= dist(t, params)

def fbinsearch(l, r, eps, check, params):
    while l + eps < r:
        m = (l + r) / 2
        if check(m, params):
            r = m
        else:
            l = m + 1
    return l
```
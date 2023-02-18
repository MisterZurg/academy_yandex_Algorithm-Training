# Конспект Лекции 5: «Префиксные суммы и два указателя»

## Что такое префиксные суммы

* Пусть у нас массив ```nums``` из ```N``` чисел и необходимость отвечать на запросы «Чему равна сумма элементов на
  полуинтервале [L, R)»

* Подсчитаем массив ```prefixsum``` длинной ```N + 1```, где ```prefixsum[k]```
  будет хранить сумму всех чисел из ```nums``` с индексами от ```0``` до ```k-1```

```  
 Индекс   | 0 | 1 | 2 |  3 |  4 |  5 |  6 |
  nums    | 5 | 3 | 8 |  1 |  4 |  4 |    |
prefixsum | 0 | 5 | 8 | 16 | 17 | 21 | 27 |
```

## Построение массивов префиксных сумм

* Массив можно построить за ```O(N): prefixsum[i] = prefixsum[i-1] + nums[i-1]```
* Не забывать про отличие размера!
* Переполнение!

## Ответ на запрос суммы на отрезке

* Ответ за ```O(1): sum(L, R) = prefixsum[R] - prefixsum[L]```
* Найдем ```sum(2, 5) = prefixsum[5] - prefixsum[2] = 21 - 8 = 13```

## Реализация RSQ через префиксные суммы

```py
def makeprefixsum(nums):
    prefixsum = [0] * (len(nums) + 1)
    for i in range(1, len(nums) + 1):
        prefixsum[i] = prefixsum[i - 1] + nums[i - 1]
    return prefixsum


# Range Some Query
def rsq(prefixsum, l, r):
    return prefixsum[r] - prefixsum[l]
```

## Задача 1

> Дана последовательность чисел длинной ```N``` и ```M``` запросов
>
> Запросы «Сколько нулей на полуинтервале ```[L, R)```»

### Решение за O(NM)

> Для каждого запроса перебираем все числа от ```L``` до ```R``` (не включительно)
> и считаем количество нулей. В худшем случае каждый запрос за ```O(N)```

```py
def countzeros(nums, l, r):
    cnt = 0
    for i in range(l, r):
        if nums[i] == 0:
            cnt += 1
    return cnt
```

### Решение за O(N + M)

> Для каждого префикса посчитаем количества нулей на нём (```prefixzeros```).
> Тогда ответ на запрос на полуинтервале ```[L, R)```:
> ```prefixzeros[R] - prefixzeros[L]```

```py
def makeprefixzeros(nums):
    prefixzeros = [0] * (len(nums) + 1)
    for i in range(1, len(nums) + 1):
        if nums[i - 1] == 0:
            prefixzeros[i] = prefixzeros[i - 1] + 1
        else:
            prefixzeros[i] = prefixzeros[i - 1]
    return prefixzeros


def countzeros(prefixzeros, l, r):
    return prefixzeros[r] - prefixzeros[l]
```

## Задача 2

> Дана последовательность чисел длинной ```N```
>
> Необходимо найти количество отрезков с нулевой суммой

### Наивное решение за O(N ^ 3)

> Перебираем начало и конец отрезка и просто просуммируем все его элементы

```py
def countzerossumranges(nums):
    cntranges = 0
    for i in range(len(nums)):
        for j in range(i + 1, len(nums) + 1):
            rangesum = 0
            for k in range(i, j):
                rangesum += nums[k]
            if rangesum == 0:
                cntranges += 1
    return cntranges
```

### Решение за O(N ^ 2)

> Перебираем начало и будем двигать конец, суммируя элементы

```py
def countzerossumranges(nums):
    cntranges = 0
    for i in range(len(nums)):
        rangesum = 0
        for j in range(i, len(nums) + 1):
            rangesum += nums[j]
            if rangesum == 0:
                cntranges += 1
    return cntranges
```

### Решение за O(N)

> Посчитаем префиксные суммы. Одинаковые префиксные суммы означают,
> что сумма на отрезке с началом и концом в позициях,
> на которых достигаются одинаковые префиксные,
> будет равна нулю

```py
def countprefixsums(nums):
    prefixsumbyvalue = {0: 1}
    nowsum = 0
    for now in nums:
        nowsum += now
        if nowsum not in prefixsumbyvalue:
            prefixsumbyvalue[nowsum] = 0
        prefixsumbyvalue += 1
    return prefixsumbyvalue


def countzerossumranges(prefixsumbyvalue):
    cntranges = 0
    for nowsum in prefixsumbyvalue:
        cntsum = prefixsumbyvalue[nowsum]
        cntranges += cntsum * (cntsum - 1) // 2
    return cntranges
```

## Два указателя

## Задача 3

> Дана отсортированная последовательность чисел длинной ```N``` и число ```K```
>
> Необходимо найти количество пар ```A, B```, таких что ```B - A > K```.

### Решение за O(N ^ 2)

> Перебираем все пары чисел и для каждой проверим условие

```py
def cntpairswithdiffgtk(sortednums, k):
    cntpairs = 0
    for first in range(len(sortednums)):
        for last in range(first, len(sortednums)):
            if sortednums[last] - sortednums[first] > k:
                cntpairs += 1
    return cntpairs
```

### Решение за O(N)

> Возьмем наименьшее число и найдем для него первое подходящее большее.
> Все ещё большие числа точно подходят.
> Возьмем в качестве меньшего числа следующее, а указатель первого подходящего
> большего будем двигать начиная с той позиции, где он находится сейчас

```py
def cntpairswithdiffgtk(sortednums, k):
    cntpairs = 0
    last = 0
    for first in range(len(sortednums)):
        while last < len(sortednums) and sortednums[last] - sortednums[first] <= k:
            last += 1
        cntpairs += len(sortednums) - last
    return cntpairs
```

# Задача 4

> Игрок в футбол обладает одной числовой характеристикой — профеосионализмом.
> Команда называется сплоченной, если профессионализм любого игрока не превосходит
> суммарный профессионализм любых двум других игроков из команды.
> Команда может состоять из любого количества игроков.
> Дана тсортированная последовательность чисел длиной ```N``` - профеосионалиэм игроков
>
> Необходимо найти максимальный суммарный профессионализм сплоченной команды.

# Решение

> Два указателя

```py
def bestteamsum(players):
    bestsum = 0
    nowsum = 0
    last = 0
    for first in range(len(players)):
        while last < len(players) and (last == first or players[first] + players[first + 1] >= players[last]):
            nowsum += players[last]
            last += 1
        bestsum = max(bestsum, nowsum)
        nowsum -= players[first]
    return bestsum
```

# Задача 5

> Даны две **отсортированные** последовательности чисел (длинной ```N``` и ```M```)
>
> Необходимо слить их в одну отсортированную последовательность.

```java
// От себя, Merge Sort спасибо подготовке и обучению Жабе
```

## Решение

> Поставим два указателя на начало каждой из последовательностей.
> Выберем тот, который указывает на меньшее число,
> запишем это число в результат и сдвинем указатель.

### Неидеальная реализация

```py
def merge(nums1, nums2):
    merged = [0] * (len(nums1) + len(nums2))
    first1 = first2 = 0
    inf = max(nums1[-1], nums2[-1]) + 1
    nums1.append(inf)
    nums2.append(inf)
    for k in range(len(nums1) + len(nums2) - 2):
        if nums1[first1] <= nums2[first2]:
            merged[k] = nums1[first1]
            first1 += 1
        else:
            merged[k] = nums2[first2]
            first2 += 1
    nums1.pop()
    nums2.pop()
    return merged
```

### Более близкая к идеалу реализация

```py
def merge(nums1, nums2):
    merged = [0] * (len(nums1) + len(nums2))
    first1 = first2 = 0
    for k in range(len(nums1) + len(nums2)):
        if first1 != len(nums1) and (first2 == len(nums2) or nums1[first1] < nums2[first2]):
            merged[k] = nums1[first1]
            first1 += 1
        else:
            merged[k] = nums2[first2]
            first2 += 1
    return merged
```
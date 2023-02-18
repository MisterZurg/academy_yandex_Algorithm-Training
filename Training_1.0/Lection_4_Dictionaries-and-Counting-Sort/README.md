# Конспект Лекции 4: «Словари и сортировка подсчётом»

## Что такое сортировка подсчётом

* Пусть необходимо массив из N целых чисел, каждое от 0 до K
* Обычная сортировка займет O(N logN)
* Будем считать количество вхождений каждого числа, а затем выводить каждое число столько раз, сколько оно встречалось.
  Это займет О(N+К) и О(К) дополнительной памяти
* Интервал значений можно сдвинуть, чтобы он был не от 0 до К, а от минимального до максимального значения в массиве

```py
def countSort(seq):
    minval = min(seq)
    maxval = max(seq)
    k = (maxval - minval + 1)  # кол-во возможных значений
    count = [0] * k
    for now in seq:
        count[now - minval] += 1
    nowpos = 0
    for val in range(0, k):
        for i in range(count[val]):
            seq[nowpos] = val + minval
            nowpos += 1
```

## Задача 1

> Дано два числа X и Y без ведущих нулей
>
> Необходимо проверить, можно ли получить первое из второго перестановкой цифр

### Решение

> Посчитаем количество вхождений каждой цифры в каждое из чисел и сравним.
> Цифры будем постепенно добывать из числа справо налево с помощью ```% 10``` и ```/ 10```

```py
def isDigitPermutation(x, y):
    def countDigits(num):
        digitcount = [0] * 10
        while num > 0:
            lastdigit = num % 10
            digitcount[lastdigit] += 1
            num //= 10
        return digitcount

    digitsx = countDigits(x)
    digitsy = countDigits(y)
    for digit range(10):
        if digitsx[digit] != digitsy[digit]:
            return False
    return True
```

## Словари

* Словарь — он как *множество*, но к каждому ключу приписано значение
* Искать по значению в словаре **нельзя**!
* Константа в сложности словарей заметно больше, чем у массивов, поэтому где можно — **лучше использовать** сортировку
  подсчётом
* Сортировку подсчётом неразумно использовать, если **данные разреженные**

## Задача 2

> На шахматной доске размером N x N находятся M ладей (ладья бьет клетки на той же горизонтали или вертикали до ближайшей занятой)
>
> Определите, сколько пар ладей бьют друг друга. Ладьи задаются парой чисел ```I``` и ```J```, обозначающих координаты клетки
> ```1 <= N <= 10^9, 0 <= M <= 2 * 10^5```

### Решение

> Для каждой занятой горизонтали и вертикали будем хранить количество ладей на них.
> Количество пар по горизонтали (вертикали) равно количеству ладей минус 1.
> Суммируем это количество пар для всех горизонталей и вертикалей.

```py
def countBeatingRooks(rookcoords):
    def addRook(rowOrCol, key):
        if key not in rowOrCol:
            rowOrCol[key] = 0
        rowOrCol[key] += 1

    def countPairs(rowOrCol):
        pairs = 0
        for key in rowOrCol:
            pairs += rowOrCol[key] - 1
        return pairs

    rooksInRow = {}
    rooksInCol = {}
    for row, col in rookcoords:
        addRook(rooksInRow, row)
        addRook(rooksInCol, col)
    return countPairs(rooksInRow) + countPairs(rooksInCol)
```

## Задача 3

> Дана строка **S**
>
> Выведите гистограмму как в примере (коды символов отсортированы)
> ```S = Hello, world```
> ```py
>       # 
>       ##
> ##########
>  !,Hdelorw
> ```

Time Complexity O(N^2)

### Решение

> Для каждого символа в словаре посчитаем, сколько раз он встречался.
> Найдем самый частый и переберем количество от этого числа до 1.
> Пройдём по всем отсортированным ключам и если количество больше счётчика — выведем #

```py
def printChart(s):
    symcount = {}
    maxsymcount = 0
    for sym in s:
        if sym not in symcount:
            symcount[sym] = 0
        symcount[sym] += 1
        maxsymcount = max(maxsymcount, symcount[sym])
    sorteduniqsyms = sorted(symcount.keys())
    for row in range(maxsymcount, 0, -1):
        for sym in sorteduniqsyms:
            if symcount[sym] >= row:
                print('#', end='')
            else:
                print(' ', end='')
        print()
    print(''.join(sorteduniqsyms))
```

## Некоторые критерии качества алгоритма

* Потребление памяти
* Время на риализацию
* Сложность поддержки
* Возможность распараллеливания
* Необходимая квалификация сотрудника
* Стоимость оборудования

## Задача 4 (C Leetcode)

> Сгруппировать слова по общим буквам
>
> ```
> Sample input: ["eat", "tea", "tan", "ate", "nat", "bat"]
> Sample output: [["ate", "eat", "tea"], ["nat", "bat"], ["tan"]]
> ```

### Решение здорового человека

> Отсортируем в каждом слове буквы и это будет выступать в роли ключа,
> а значение будет список слов

```py
def groupWords(words):
    groups = {}
    for word in words:
        sortedword = ''.join(sorted(word))
        if sortedword not in groups:
            groups[sortedword] = []
        groups[sortedword].append(word)
    ans = []
    for sortedword in groups:
        ans.append(groups[sortedword])
    return ans
```

### Подозрения здорового человека

```py
sortedword = ''.join(sorted(word))
```

> Вдруг слово будет длинное (N)? Сортировка займет O(N logN).
> Количество различных букв в слове K <= N, можем посчитать
> количество каждой за O(N) и отсортировать за O(K logK), теоритически словарь

### Оптимизация (?) здорового человека

```py
def groupWords(words):
    def keybyword(word):
        symcnt = {}
        for sym in word:
            if sym not in symcnt:
                symcnt[sym] = 0
            symcnt[sym] += 1
        lst = []
        for sym in ''.join(symcnt.keys()):
            lst.append(sym)
            lst.append(str(symcnt[sym]))
        return ''.join(lst)

    groups = {}
    for word in words:
        groupkey = keybyword(word)
        if groupkey not in groups:
            groups[groupkey] = []
        groups[groupkey].append(word)
    ans = []
    for groupkey in groups:
        ans.append(groups[groupkey])
    return ans
```
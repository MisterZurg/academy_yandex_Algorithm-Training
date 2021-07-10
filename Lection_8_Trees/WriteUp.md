# Конспект Лекции 8: «Деревья»

## Менеджмент памяти

> У нас есть заранее **неизвестное** количество структур с двумя ссылками на другие структуры
>
> Мы знаем заранее, какое **максимальное** количество структур может существовать одновременно
>
> Хотим научиться **выделять и освобождать** память
> ![alt tag](initState.JPG "Состояние после инициализации")
>
> ![alt tag](elemone.JPG "Состояние после инициализации")
>
> ![alt tag](elemtwo.JPG "Состояние после инициализации")
>
> ![alt tag](elemzeroindex.JPG "Состояние после инициализации")

## Код менеджера памяти

```py
def initmemory(maxn):
    memory = []
    for i in range(maxn):
        memory.append([0, i + 1, 0])
    return [memory, 0]


def newnode(memstruct):
    memory, firstfree = memstruct
    memstruct[1] = memory[firstfree][1]
    return firstfree


def delnode(memstruct, index):
    memory, firstfree = memstruct
    memory[index][1] = firstfree
    memstruct[1] = index

```

## Что такое бинарное дерево поиска

> У каждого узла есть ключ и два сына — левый и правый
>
> В левом поддереве ключи меньше, а в правом — больше
>
> Если ключи поступают в случайном порядке, то глубина дерева будет ```O(log N)```
> ![alt tag](binarytrees.JPG "Состояние после инициализации")
> ![alt tag](binarytreesearch.JPG "Состояние после инициализации")
> ![alt tag](binarytreeadd.JPG "Состояние после инициализации")
> ![alt tag](binarytreedelete.JPG "Состояние после инициализации")
> ![alt tag](binarytreedelete2.JPG "Состояние после инициализации")
> ![alt tag](binarytreedelete3.JPG "Состояние после инициализации")
> ![alt tag](treeafterremoving.JPG "Состояние после инициализации")

## Реализация поиска

```py
def find(memstruct, root, x):
    key = memstruct[0][root][0]
    if x == key:
        return root
    elif x < key:
        left = memstruct[0][root][1]
        if left == -1:
            return -1
        else:
            return find(memstruct, left, x)
    elif x > key:
        right = memstruct[0][root][2]
        if right == -1:
            return -1
        else:
            return find(memstruct, right, x)
```

## Реализация добавления

```py
def createandfillnode(memstruct, key):
    index = newnode(memstruct)
    memstruct[0][index][0] = key
    memstruct[0][index][1] = -1
    memstruct[0][index][2] = -1


def add(memstruct, root, x):
    key = memstruct[0][root][0]
    if x < key:
        left = memstruct[0][root][1]
        if left == -1:
            memstruct[0][root][1] = createandfillnode(memstruct, x)
        else:
            add(memstruct, left, x)
    elif x > key:
        right = memstruct[0][root][2]
        if right == -1:
            memstruct[0][root][2] = createandfillnode(memstruct, x)
        else:
            add(memstruct, right, x)


# Как создать какое-нибудь дерево?
memstruct = initmemory(20)
root = createandfillnode(memstruct, 8)
add(memstruct, root, 10)
add(memstruct, root, 9)
add(memstruct, root, 14)
add(memstruct, root, 13)
add(memstruct, root, 3)
add(memstruct, root, 1)
add(memstruct, root, 6)
add(memstruct, root, 4)
add(memstruct, root, 7)
```
![alt tag](treeinPython.JPG "Представление дерева в Python")

## Не бинарные деревья
> У узлов дерева может быть больше двух сыновей, тогда их нужно хранить **списком**
> 
> **Примеры:** дерево каталогов и файлов, html-документы (DOM-дерево), дерево классов в программе и т.д.
> 
> Обходим также как и бинрное, просто запуская **рекурсивную функцию** для всех детей

## Сериализация дерева Хаффмана
> Алгоритм Хаффмана позволяет сопоставить более часто встречающимся символам болев короткий код
> 
> Каждый раз берем два самых редко встречающихся символа и объединяем их в один узел 
> 
> Строим бинарное дерево, кладем буквы в листья. Переход в левого сына кодируетса числом О, в правого -1, а код символа - это все ребра на пути от корня до листа 
>
> В примере буква «а» встречается 4 раза, «б» — 3 раза, а «в» и «г» по одному разу
> ![alt tag](huffman.JPG "Дерево Хаффмана")


## Как сохранить структуру дерева ввиде строки?
> L — в левого ребёнка, R — в правого, U — в предка.
> 
> **LURLLURUURUU**
> 
> D — в наиболее левого непосещённого ребенка (детей всегда либо два, либо ноль)
> 
> **DUDDDUDUUDUU**
> 
> Теперь U означает, что мы поднимаемся вверх до тех пор, пока приходимиз правого ребенка.
> Если пришли в вершину из левого — сразу пойдем в правого
> 
> **DUDDUU**

## Восстановление дерева по записи
> D — в наиболее левого непосещённого ребенка (детей всегда либо два, либо ноль)
> U — поднимаемся вверх до тех пор, пока приходим из правого ребенка.
> Если пришли в вершину из левого — сразу пойдем в правого
> 
> **DUDDUU**

## Задача
> По сериализованной записи восстановить код листьев.

## Создание дерева и его обход
```python
def maketree(serialized):
    tree = {'left' : None, 'right' : None, 'up' : None, 'type' : 'root'}
    nownode = tree
    for sym in serialized:
        if sym == 'D':
            newnode = {'left' : None, 'right' : None, 'up' : nownode, 'type' : 'left'}
            nownode['left'] = newnode
            nownode = newnode
        elif sym == 'U':
            while nownode['type'] == 'right':
                nownode = newnode['up']
            nownode = newnode['up']
            newnode = {'left' : None, 'right' : None, 'up' : nownode, 'type' : 'right'}
            nownode['right'] = newnode
            nownode = newnode
    return tree


def traverse(root, prefix):
    if root['left'] is None and root['right'] is None:
        return [''.join(prefix)]
    prefix.append('0')
    ans = traverse(root['left'], prefix)
    prefix.pop()
    prefix.append('1')
    
    ans.extend(traverse(root['right'], prefix))
    prefix.pop()
    return ans
```
def sort2(first, second):
    if first < second:
        return first, second
    return second, first


a = int(input())
b = int(input())
c = int(input())
d = int(input())  # меньшая сторона отверстия
e = int(input())  # большая сторона отверстия
# Топим самый больший элемент
a, b = sort2(a, b)
b, c = sort2(b, c)
a, b = sort2(a, b)

d, e = sort2(d, e)

if a <= d and b <= e:
    print("YES")
else:
    print("NO")
# A. Покраска деревьев
> [!NOTE]
> Командная олимпиада Высшая проба 2019 Москва, Липецк, Самара, Ижевск, Челябинск 21 апреля 2019


# B. Футбольный комментатор


# C. Форматирование файла
```
n = int(input())
a = (int(input()) for _ in range(n))

print(sum(k // 4 + min(k % 4, 2) for k in a))
```
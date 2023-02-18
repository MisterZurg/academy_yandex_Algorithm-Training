# Скандально НЕ известный алгоритм половинка от QuickSort
# катая порядковая статистика
def kth_rearrange(seq, left, right, k):
    left = 0
    right = len(seq) - 1
    while left < right:
        # Выбираем "случайное" число x
        x = seq[(left + right) // 2]
        eqxfirst = left
        gtxfirst = left
        for i in range(left, right + 1):
            now = seq[i]
            # Если текущий элемент = x, просто свапаем
            if now == x:
                seq[i] = seq[gtxfirst]
                seq[gtxfirst] = now
                gtxfirst += 1
            # Если очередное число меньше, чем x
            elif now < x:
                seq[i] = seq[gtxfirst]
                seq[gtxfirst] = seq[eqxfirst]
                seq[eqxfirst] = now
                gtxfirst += 1
                eqxfirst += 1
        # Меняем левую/правую границы
        if k < eqxfirst:
            right = eqxfirst - 1
        elif k >= gtxfirst:
            left = gtxfirst
        else:
            return


seq = list(map(int, input().split()))
kth_rearrange(seq, 0, len(seq), len(seq) - 1)  # Max
kth_rearrange(seq, 0, len(seq) - 1, len(seq) - 2)  # Max 2, 3
kth_rearrange(seq, 0, len(seq) - 3, 2)  # Min1, 2

if seq[-1] * seq[-2] * seq[-3] >= seq[-1] * seq[0] * seq[1]:
    print(seq[-1], seq[-2], seq[-3])
else:
    print(seq[-1], seq[0], seq[1])

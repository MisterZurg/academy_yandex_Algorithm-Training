# СТРОКА ГРЕЯ
# 1 2 3 4 5 4 3 2 1 1 2 1 2 2 1 2 1

# Префикс функция
# Скандально известный алгоритм Кнута-Морриса-Пратта

# Time Complexity O(N^2)
def slow(seq):
    for start in range(len(seq)):
        i = start
        j = len(seq) - 1
        while i < len(seq) and j >= 0 and seq[i] == seq[j] and i <= j:
            i += 1
            j += 1
        if i > j:
            ans = []
            for i in range(start - 1, -1, -1):
                ans.append(seq[i])
            return ans

n = int(input())
seq = list(map(int, input().split()))
ans = slow(seq)
print(len(ans))
print(*ans)
a = int(input())
b = int(input())
n = int(input())
m = int(input())

min1 = (a + 1) * (n - 1) + 1
max1 = min1 + 1 + 2 * a

min2 = (b + 1) * (m - 1) + 1
max2 = min2 + 1 + 2 * b

minans = max(min1, min2)
maxans = min(max1, max2)

if maxans < minans:
    print(-1)
else:
    print(minans, maxans)

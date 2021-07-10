def extend(rect, d):
    minPlus, maxPlus, minMius, maxMins = rect
    return (minPlus - d, maxPlus + d, minMius - d, maxMins + d)


def intersect(rect1, rect2):
    ans = [max(rect1[0], rect2[0]),
           min(rect1[1], rect2[1]),
           max(rect1[2], rect2[2]),
           min(rect1[3], rect2[3])
           ]
    return ans


t, d, n = int(input().split())
posRect = (0, 0, 0, 0)
for _ in range(n):
    posRect = extend(posRect, t)
    navX, navY = map(int, input().split())
    navRect = extend((navX + navY, navX + navY, navX - navY, navX - navY), d)
    posRect = intersect(posRect, navRect)

points = []

for xPlusY in range(posRect[0],posRect[1] + 1):
    for xMinusY in range(posRect[2],posRect[3] + 1):
        if (xPlusY + xMinusY) % 2 == 0:
            x = (xPlusY + xMinusY) // 2
            y = xPlusY - x
            points.append((x, y))

print(len(points))
for point in points:
    print(*point)
import random

def getentranceandfloor(flantno, flatsonfloor, floors):
    floorsbefore = (flantno - 1) // flatsonfloor
    entrance = floorsbefore // floors + 1
    floor = floorsbefore % floors + 1
    return entrance, floor


def check(k1, m, k2, p2, n2, flatsonfloor):
    entrance2, floor2 = getentranceandfloor(k2, flatsonfloor, m)
    if entrance2 == p2 and floor2 == n2:
        return getentranceandfloor(k1, flatsonfloor, m)
    return -1, -1


def slow(k1, m, k2, p2, n2):
    ent = -1
    floor = -1
    goodflag = False
    for flatsonfloor in range(1, maxrandval + 1):   # Вместо 10^6 (миллиона)
        nent, nfloor = check(k1, m, k2, p2, n2, flatsonfloor)
        if nent != -1:
            goodflag = True
            if ent == -1:
                ent,floor = nent, nfloor
            elif ent!= nent and ent != 0:
                ent = 0
            elif floor != nfloor and floor !=0:
                floor = 0
    if goodflag:
        return ent, floor
    else:
        return -1, -1


def fast(k1, m, k2, p2, n2):
    f2 = (p2 -1) * m + n2
    if f2 == 1:
        if k1 <= k2:
            return 1, 1
        elif m == 1:
            return 0, 1
        else:
            return 0, 0
    else:
        onfmax = (k2 + f2 - 2) // (f2 - 1) - 1
        onfmin = (k2 + f2 - 1) // f2
        if onfmax < onfmin:
            return -1, -1
        else:
            f1min = (k1 - 1) // onfmax + 1
            f1max = (k1 - 1) // onfmin + 1

            p1min = (f1min - 1) // m + 1
            p1max = (f1max - 1) // m + 1

            n1min = (f1min - 1) % m + 1
            n1max = (f1max - 1) % m + 1

            if p1min == p1max and n1min == n1max:
                return p1min, n1min
            elif p1min == p1max:
                return p1min, 0
            elif n1min == n1max:
                return 0, n1min

# Стрессовое тестирование
maxrandval = 10
while True:
    randvals = [0] * 5
    for i in range(5):
        randvals[i] = random.randint(1, maxrandval)
    k1, m, k2, p2, n2 = randvals
    slowans = slow(k1, m, k2, p2, n2)
    fastans = fast(k1, m, k2, p2, n2)

    if slowans == fastans:
        print("OK")
    else:
        print("WA", fastans, slowans)
        break
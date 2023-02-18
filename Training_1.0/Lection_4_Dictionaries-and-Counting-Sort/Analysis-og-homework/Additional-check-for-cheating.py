#import sys
# sys stdin
def onlyLettersDigitsAndUndderscores(line):
    ans = []
    for c in line:
        if c.isalnum() or c.isdigit() or c == '_':
            ans.append(c)
        else:
            ans.append(' ')
    return ''.join(ans)


def isCorrect(word, sr):
    if word.isdigit():
        return False
    if not word[0].isdigit() or stDigit:
        return True
    return False


with open('input.txt', 'r') as fin:
    n, caseSens, stDigit = fin.readlines().strip().split()
    n = int(n)
    caseSens = caseSens == 'yes'
    stDigit = stDigit == 'yes'
    keywords = set()
    for _ in range(n):
        keyword = fin.readline().strip()
        if not caseSens:
            keyword = keyword.lower()
        keywords.add(keyword)

    cntPosIds = {}
    wordNumber = 0
    for line in fin.readlines():
        line = onlyLettersDigitsAndUndderscores(line.strip())
        words = line.split()
        for word in words:
            if not caseSens:
                word = word.lower()
            if word in keywords:
                continue
            if isCorrect(word, stDigit):
                wordNumber += 1
                if word not in cntPosIds:
                    cntPosIds[word] = [0, wordNumber]
                cntPosIds[word][0] += 1

    bestWord = ''
    maxCntPos = [0, 0]
    for word in cntPosIds:
        if cntPosIds[word][0] > maxCntPos[0]:
            maxCntPos = cntPosIds[word]
            bestWord = word
        elif cntPosIds[word][0] == maxCntPos[0] and cntPosIds[word][1] < maxCntPos[1]:
            maxCntPos = cntPosIds[word]
            bestWord = word
    print(bestWord)
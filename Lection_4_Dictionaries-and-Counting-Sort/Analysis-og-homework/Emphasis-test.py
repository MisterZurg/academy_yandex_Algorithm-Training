n = int(input())
dictlower = {}
for _ in range(n):
    s = input()
    sLower = s.lower()
    if sLower not in dictlower:
        dictlower[sLower] = set()
    dictlower[sLower].add(s)

text = input()
mistakes = 0
for word in text.split():
    wordLower = word.lower()
    if wordLower in dictlower:
        if word not in dictlower[wordLower]:
            mistakes += 1
    else:
        stresses = 0
        for c in word:
            if c.isupper():
                stresses += 1
        if stresses != 1:
            mistakes += 1
print(mistakes)


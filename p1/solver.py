lines = []

with open('./input.txt') as file:
    lines = file.readlines()

maxElf, curElf = 0, 0

for line in lines:
    if line == '\n':
        if curElf > maxElf:
            maxElf = curElf
        curElf = 0
    else:
        curElf += int(line)

print(maxElf)
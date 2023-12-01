lines = []

with open('./input.txt') as file:
    lines = file.readlines()

elves, curElf = [], 0

for line in lines:
    if line == '\n':
        elves.append(curElf)
        curElf = 0
    else:
        curElf += int(line)

elves.sort(reverse=True)

print(sum(elves[:3]))
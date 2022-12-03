lines = []

with open('./input.txt') as file:
    lines = file.readlines()

score, conv, wins = 0, {"A":"X","B":"Y","C":"Z"}, {"X":"Z","Y":"X","Z":"Y"}

for line in lines:
    oppo, me, cur = conv[line[0]], line[2], 0

    if me == "X":
        cur += 1
    elif me == "Y":
        cur += 2
    else:
        cur += 3
    
    if oppo == me:
        cur += 3
    elif wins[me] == oppo:
        cur += 6
    
    score += cur

print(score)

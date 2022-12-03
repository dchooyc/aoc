lines = []

with open('./input.txt') as file:
    lines = file.readlines()

score = 0
loses = {"A":3,"B":1,"C":2}
value = {"A":1,"B":2,"C":3}
wins = {"A":2,"B":3,"C":1}

for line in lines:
    oppo, end, cur = line[0], line[2], 0

    if end == "X":
        cur += loses[oppo]
    elif end == "Y":
        cur += value[oppo] + 3
    else:
        cur += wins[oppo] + 6
    
    score += cur

print(score)

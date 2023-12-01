class Solver:
  def __init__(self, lines):
    self.lines = lines
    self.priority = self.generatePriority()


  def generatePriority(self):
    priority = {}

    for i in range(1, 27):
      lower = chr((i - 1) + ord("a"))
      upper = chr((i - 1) + ord("A"))
      priority[lower] = i
      priority[upper] = i + 26
    
    return priority


  def sumPriority(self):
    res = 0

    for line in self.lines:
      mid = len(line) >> 1
      first = line[:mid]
      second = line[mid:]
      items = self.findCommonItems(first, second)

      for item in items:
        res += self.priority[item]

    return res
  

  def sumPriorityGroup(self):
    res = 0

    for i in range(0, len(self.lines), 3):
      first = self.lines[i]
      second = self.lines[i + 1]
      third = self.lines[i + 2]

      d1 = self.findCommonItems(first, second)
      d2 = self.findCommonItems(second, third)

      badge = ""

      for item in d1:
        if item in d2:
          badge = item
          break
      
      if badge == "":
        continue
      
      res += self.priority[badge]
    
    return res


  def findCommonItems(self, string1, string2):
    res, d = {}, {}

    for char in string1:
      d[char] = True
    
    for char in string2:
      if char in d:
        res[char] = True
    
    return res


lines = []

with open('./input.txt') as file:
    lines = file.readlines()

res = Solver(lines)

print(res.sumPriority())
print(res.sumPriorityGroup())
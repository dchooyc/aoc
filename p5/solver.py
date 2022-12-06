class Solver:
  def __init__(self, lines):
    self.lines = lines


  def topCrates(self):
    crates, moves = [], []

    for i in range(len(self.lines)):
      if self.lines[i] == '\n':
        crates = self.lines[:i - 1]
        moves = self.lines[i + 1:]
        break

    stacked = self.stackCrates(crates)

    for move in moves:
      a = move.split(" ")
      quant, src, dst = int(a[1]), int(a[3]) - 1, int(a[5]) - 1

      for i in range(quant):
        box = stacked[src].pop(0)
        stacked[dst].insert(0, box)
    
    res = ""

    for i in range(9):
      if stacked[i][0] != "":
        res += stacked[i][0]
    
    return res


  def newerCrane(self):
    crates, moves = [], []

    for i in range(len(self.lines)):
      if self.lines[i] == '\n':
        crates = self.lines[:i - 1]
        moves = self.lines[i + 1:]
        break

    stacked = self.stackCrates(crates)

    for move in moves:
      a = move.split(" ")
      quant, src, dst = int(a[1]), int(a[3]) - 1, int(a[5]) - 1

      load = stacked[src][:quant]
      stacked[src] = stacked[src][quant:]
      stacked[dst] = load + stacked[dst]
    
    res = ""

    for i in range(9):
      if stacked[i][0] != "":
        res += stacked[i][0]
    
    return res


  def stackCrates(self, crates):
    res = [[] for i in range(9)]

    for crate in crates:
      i, index = 1, 0
      
      while i < len(crate):
        if crate[i] != " ":
          res[index].append(crate[i])
        
        index += 1
        i += 4

    return res


lines = []

with open('./input.txt') as file:
    lines = file.readlines()

res = Solver(lines)
print(res.topCrates())
print(res.newerCrane())
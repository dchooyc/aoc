class Solver:
  def __init__(self, lines):
    self.lines = lines
    self.grid = self.generateGrid()

  
  def generateGrid(self):
    grid = []

    for line in self.lines:
      row = []

      for char in line[:len(line) - 1]:
        row.append(int(char))
      
      grid.append(row)
    
    return grid

  
  def countVisible(self):
    m, n = len(self.grid), len(self.grid[0])

    res = (m * 2) + (n * 2) - 4

    for i in range(1, m - 1):
      for j in range(1, n - 1):
        if self.checkCell(i, j, m, n):
          res += 1
    
    return res
  

  def findMaxScore(self):
    m, n = len(self.grid), len(self.grid[0])

    res = 0

    for i in range(1, m - 1):
      for j in range(1, n - 1):
        res = max(res, self.checkScore(i, j, m, n))
    
    return res


  def checkScore(self, i, j, m, n):
    cell = self.grid[i][j]
    up, down, left, right = -1, -1, -1, -1
    u, d, l, r = i - 1, i + 1, j - 1, j + 1

    while u >= 0:
      if self.grid[u][j] >= cell:
        up = i - u
        break
      u -= 1

    if up == -1:
      up = i
    
    while d < m:
      if self.grid[d][j] >= cell:
        down = d - i
        break
      d += 1
    
    if down == -1:
      down = m - 1 - i
    
    while l >= 0:
      if self.grid[i][l] >= cell:
        left = j - l
        break
      l -= 1
    
    if left == -1:
      left = j
      
    while r < n:
      if self.grid[i][r] >= cell:
        right = r - j
        break
      r += 1
    
    if right == -1:
      right = n - 1 - j
    
    return up * down * left * right

  
  def checkCell(self, i, j, m, n):
    cell = self.grid[i][j]
    up, down, left, right = True, True, True, True
    u, d, l, r = i - 1, i + 1, j - 1, j + 1

    while u >= 0:
      if self.grid[u][j] >= cell:
        up = False
        break
      u -= 1
    
    if up:
      return True
    
    while d < m:
      if self.grid[d][j] >= cell:
        down = False
        break
      d += 1
    
    if down:
      return True
    
    while l >= 0:
      if self.grid[i][l] >= cell:
        left = False
        break
      l -= 1
    
    if left:
      return True
      
    while r < n:
      if self.grid[i][r] >= cell:
        right = False
        break
      r += 1

    if right:
      return True
    
    return False


lines = []

with open('./input.txt') as file:
    lines = file.readlines()

res = Solver(lines)
print(res.countVisible())
print(res.findMaxScore())
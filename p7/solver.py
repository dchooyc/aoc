class Solver:
  def __init__(self, lines):
    self.lines = lines
    self.d = self.generateDirectories()

  
  def generateDirectories(self):
    d, path = {}, []

    for line in self.lines:
      a = line[:len(line) - 1].split(" ")

      if a[0] == "$" and a[1] == "cd":
        if a[2] == "..":
          path.pop()
        else:
          path.append(a[2])
      elif a[0] != "$" and a[0] != "dir":
        size = int(a[0])

        for i in range(len(path), 0, -1):
          curPath = "/".join(path[:i])
          
          if curPath in d:
            d[curPath] += size
          else:
            d[curPath] = size
    
    return d


  def findDelete(self):
    target = self.d["/"] - 40000000
    res  = 70000000

    for directory in self.d:
      if self.d[directory] >= target:
        res = min(res, self.d[directory])
    
    return res


  def findSmall(self):
    res = 0

    for directory in self.d:
      if self.d[directory] <= 100000:
        res += self.d[directory]
    
    return res


lines = []

with open('./input.txt') as file:
    lines = file.readlines()

res = Solver(lines)
print(res.findSmall())
print(res.findDelete())
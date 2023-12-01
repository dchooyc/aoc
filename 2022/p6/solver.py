class Solver:
  def __init__(self, lines):
    self.lines = lines

  def packetMarker(self):
    s = self.lines[0]
    letters = [s[0], s[1], s[2]]

    for i in range(3, len(s)):
      letters.append(s[i])

      if self.check(letters):
        return i + 1
      
      letters.pop(0)
    
    return len(s)
  

  def startMessage(self):
    s = self.lines[0]
    letters = []

    for i in range(13):
      letters.append(s[i])
    
    for i in range(13, len(s)):
      letters.append(s[i])

      if self.check(letters):
        return i + 1
      
      letters.pop(0)

    return len(s)

  
  def check(self, letters):
    d = {}

    for letter in letters:
      if letter in d:
        return False
      else:
        d[letter] = True
    
    return True


lines = []

with open('./input.txt') as file:
    lines = file.readlines()

res = Solver(lines)
print(res.packetMarker())
print(res.startMessage())
class Solver:
  def __init__(self, lines):
    self.lines = lines

  
  def countContainedRanges(self):
    res = 0

    for line in self.lines:
      sections = line.split(",")
      first = sections[0].split("-")
      second = sections[1].split("-")
      start1, end1 = int(first[0]), int(first[1])
      start2, end2 = int(second[0]), int(second[1])

      if (self.within(start1, end1, start2, end2) or 
          self.within(start2, end2, start1, end1)):
        res += 1
    
    return res


  def countOverlappedRanges(self):
    res = 0

    for line in self.lines:
      sections = line.split(",")
      first = sections[0].split("-")
      second = sections[1].split("-")
      start1, end1 = int(first[0]), int(first[1])
      start2, end2 = int(second[0]), int(second[1])

      if self.overlaps(start1, end1, start2, end2):
        res += 1
    
    return res


  def within(self, start1, end1, start2, end2):
    return start2 >= start1 and end2 <= end1
  

  def overlaps(self, start1, end1, start2, end2):
    return start2 <= end1 and start1 <= end2


lines = []

with open('./input.txt') as file:
    lines = file.readlines()

res = Solver(lines)
print(res.countContainedRanges())
print(res.countOverlappedRanges())
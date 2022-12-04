class Solver:
  def __init__(self, lines):
    self.lines = lines


lines = []

with open('./input.txt') as file:
    lines = file.readlines()

res = Solver(lines)
print(res.lines)
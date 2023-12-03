class Solver:
	def __init__(self, input):
		self.lines = self.get_lines(input)
		
	def get_lines(self, target):
		lines = []
		
		with open(target) as file:
			lines = file.readlines()
		
		stripped_lines = [line.strip() for line in lines]
    
		return stripped_lines
	
	def sum_parts(self):
		lines = [list(l) for l in self.lines]
		res = 0
		cur, adjacent = 0, False

		for i in range(len(lines)):
			for j in range(len(lines[0])):
				char = lines[i][j]

				if ord(char) >= ord("0") and ord(char) <= ord("9"):
					cur *= 10
					cur += int(char)
					if self.has_adjacent_symbol(lines, i, j):
						adjacent = True
				else:
					if cur != 0 and adjacent:
						res += cur
					cur, adjacent = 0, False

			if cur != 0 and adjacent:
				res += cur
			cur, adjacent = 0, False
		
		return res
	
	def has_adjacent_symbol(self, lines, i, j):
		dirs = [[-1, 0], [1, 0], [0, -1], [0, 1], [-1, -1], [1, 1], [-1, 1], [1, -1]]

		for dir in dirs:
			x, y = dir[0] + i, dir[1] + j
			if x >= 0 and x < len(lines) and y >= 0 and y < len(lines[0]):
				char = lines[x][y]
				if char != "." and (ord(char) < ord("0") or ord(char) > ord("9")):
					return True
		
		return False
	
	def common_gears(self):
		lines = [list(l) for l in self.lines]
		cur, agears = 0, {}
		gears = {}

		for i in range(len(lines)):
			for j in range(len(lines[0])):
				char = lines[i][j]

				if ord(char) >= ord("0") and ord(char) <= ord("9"):
					cur *= 10
					cur += int(char)
					ag = self.adjacent_gears(lines, i, j)
					for g in ag:
						agears[g] = True
				else:
					if cur != 0 and len(agears) != 0:
						for gear in agears:
							if gear in gears:
								gears[gear].append(cur)
							else:
								gears[gear] = [cur]
					cur, agears = 0, {}

			if cur != 0 and len(agears) != 0:
				for gear in agears:
					if gear in gears:
						gears[gear].append(cur)
					else:
						gears[gear] = [cur]
			cur, agears = 0, {}

		res = 0
		
		for key in gears:
			if len(gears[key]) == 2:
				res += gears[key][0] * gears[key][1]

		return res
	
	def adjacent_gears(self, lines, i, j):
		dirs = [[-1, 0], [1, 0], [0, -1], [0, 1], [-1, -1], [1, 1], [-1, 1], [1, -1]]
		gears = []

		for dir in dirs:
			x, y = dir[0] + i, dir[1] + j
			if x >= 0 and x < len(lines) and y >= 0 and y < len(lines[0]):
				char = lines[x][y]
				if char == "*":
					gears.append(str(x) + "," + str(y))
		
		return gears

res = Solver("./input.txt")
print(res.sum_parts())
print(res.common_gears())

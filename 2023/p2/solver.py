class Solver:
	def __init__(self, input):
		self.lines = self.get_lines(input)
		
	def get_lines(self, target):
		lines = []
		
		with open(target) as file:
			lines = file.readlines()
		
		stripped_lines = [line.strip() for line in lines]
    
		return stripped_lines
	
	def possible(self):
		# 12 red cubes, 13 green cubes, and 14 blue cubes
		r, g, b = 12, 13, 14
		res = 0

		for i in range(len(self.lines)):
			line = self.lines[i]
			parts = line.split(": ")
			handfuls = parts[1].split("; ")
			cubes = {}

			for handful in handfuls:
				colours = handful.split(", ")

				for colour in colours:
					vals = colour.split(" ")
					cnt = int(vals[0])
					col = vals[1]

					if col in cubes:
						cubes[col] = max(cubes[col], cnt)
					else:
						cubes[col] = cnt
			
			passed = True

			for col in cubes:
				if col == "red" and cubes[col] > r:
					passed = False
				elif col == "green" and cubes[col] > g:
					passed = False
				elif col == "blue" and cubes[col] > b:
					passed = False

			if passed:
				res += (i + 1)
		
		return res
	
	def power(self):
		res = 0

		for i in range(len(self.lines)):
			line = self.lines[i]
			parts = line.split(": ")
			handfuls = parts[1].split("; ")
			cubes = {
				"red": 0,
				"green": 0,
				"blue": 0
			}

			for handful in handfuls:
				colours = handful.split(", ")

				for colour in colours:
					vals = colour.split(" ")
					cnt = int(vals[0])
					col = vals[1]
					cubes[col] = max(cubes[col], cnt)
			
			p = 1

			for col in cubes:
				p *= cubes[col]
			
			res += p
		
		return res

 
res = Solver("./input.txt")
print(res.possible())
print(res.power())

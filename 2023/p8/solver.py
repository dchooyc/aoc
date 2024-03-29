class Solver:
	def __init__(self, input):
		self.lines = self.get_lines(input)
		self.instruct =  self.lines[0]
		self.nodes = self.get_nodes(self.lines[2:])
		
	def get_lines(self, target):
		lines = []
		
		with open(target) as file:
			lines = file.readlines()
		
		stripped_lines = [line.strip() for line in lines]
    
		return stripped_lines
	
	def get_nodes(self, lines):
		nodes = {}

		for line in lines:
			parts = line.split(" = ")
			node = parts[0]
			children = parts[1].split(", ")
			left = children[0][1:]
			right = children[1]
			right = right[:len(right) - 1]
			nodes[node] = [left, right]
		
		return nodes
	
	def steps(self):
		point, steps = "AAA", 0
		index = 0

		while point != "ZZZ":
			if index == len(self.instruct):
				index = 0
			if self.instruct[index] == "L":
				point = self.nodes[point][0]
			else:
				point = self.nodes[point][1]
			index += 1
			steps += 1
		
		return steps
	
	def ghosts(self):
		points = []

		for node in self.nodes:
			if node[2] == "A":
				points.append(node)
		
		steps = []

		for point in points:
			s = self.steps_point(point)
			steps.append(s)
		
		lcm = steps[0]

		for i in range(1, len(steps)):
			lcm = self.lcm(lcm, steps[i])
		
		return lcm

	def gcd(self, num, den):
		if num < den:
			num, den = den, num
		rem = num % den
		while rem != 0:
			num, den = den, rem
			rem = num % den
		return den
	
	def lcm(self, num, den):
		return (num * den) // self.gcd(num, den)
	
	def steps_point(self, point):
		steps = 0
		index = 0

		while point[2] != "Z":
			if index == len(self.instruct):
				index = 0
			if self.instruct[index] == "L":
				point = self.nodes[point][0]
			else:
				point = self.nodes[point][1]
			index += 1
			steps += 1
		
		return steps

res = Solver("./input.txt")
print(res.steps())
print(res.ghosts())

class Solver:
	def __init__(self, input):
		self.lines = self.get_lines(input)
		self.list1 = []
		self.list2 = []
		
	def get_lines(self, target):
		lines = []
		
		with open(target) as file:
			lines = file.readlines()
		
		stripped_lines = [line.strip() for line in lines]
    
		return stripped_lines
	
	def get_lists(self, lines):
		list1, list2 = [], []

		for line in lines:
			vals = line.split("   ")
			list1.append(int(vals[0]))
			list2.append(int(vals[1]))
		
		list1.sort()
		list2.sort()

		self.list1 = list1
		self.list2 = list2
	
	def part1(self):
		res = 0

		for i in range(len(self.list1)):
			res += abs(self.list1[i] - self.list2[i])
		
		return res
	
	def part2(self):
		a = {}
		b = {}

		for num in self.list1:
			a[num] = True
		
		for num in self.list2:
			if num in b:
				b[num] += 1
			else:
				b[num] = 1
			
		res = 0

		for key in a:
			if key in b:
				res += key * b[key]
		
		return res

  
res = Solver("./input.txt")
res.get_lists(res.lines)
p1 = res.part1()
p2 = res.part2()
print(p1)
print(p2)

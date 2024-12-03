class Solver:
	def __init__(self, input):
		self.lines = self.get_lines(input)

	def part1(self):
		res = 0
		for line in self.lines:
			res += self.mull(line)
		return res
	
	def part2(self):
		res = 0
		lines = self.clean_lines(self.lines)
		for line in lines:
			res += self.mull(line)
		return res

	def clean_lines(self, lines):
		line = ""
		for l in lines:
			line += l
		res = []

		parts = line.split("don't()")
		res.append(parts[0])
		for i in range(1, len(parts)):
			cur = parts[i].split("do()")
			for i in range(1, len(cur)):
				res.append(cur[i])
			
		return res

	def mull(self, line):
		m = "mul("
		res = 0
		i = 0
		while i < (len(line) - 3):
			if line[i:i+4] == m:
				j = i + 4
				while j < len(line) and line[j] != ")":
					j += 1
				if j < len(line) and line[j] == ")":
					vals = line[i+4:j]
					nums = vals.split(",")
					if len(nums) == 2:
						res += self.conv(nums[0]) * self.conv(nums[1])
			i += 1
		
		return res

	def conv(self, num):
		res = 0

		for i in range(len(num)):
			char = num[i]
			if ord(char) >= ord("0") and ord(char) <= ord("9"):
				res *= 10
				res += int(char)
			else:
				return 0
		
		return res

	def get_lines(self, target):
		lines = []
		
		with open(target) as file:
			lines = file.readlines()
		
		stripped_lines = [line.strip() for line in lines]
    
		return stripped_lines
  
res = Solver("./input.txt")
p1 = res.part1()
p2 = res.part2()
print(p1)
print(p2)

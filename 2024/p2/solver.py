class Solver:
	def __init__(self, input):
		self.lines = self.get_lines(input)

	def part_1(self):
		res = 0
		for line in self.lines:
			vals = line.split(" ")
			nums = [int(val) for val in vals]
			safe = True
			inc = nums[0] < nums[1]

			for i in range(1, len(nums)):
				diff = nums[i] - nums[i - 1]
				if (diff > 0 and not inc) or (diff < 0 and inc):
					safe = False
					break
				if abs(diff) < 1 or abs(diff) > 3:
					safe = False
					break
			
			if safe:
				res += 1
		return res
	
	def part_2(self):
		res = 0
		for line in self.lines:
			vals = line.split(" ")
			nums = [int(val) for val in vals]
			for i in range(len(nums)):
				if self.check_safe(nums[:i] + nums[i + 1:]):
					res += 1
					break
		return res
	
	def check_safe(self, nums):
			safe = True
			inc = nums[0] < nums[1]

			for i in range(1, len(nums)):
				diff = nums[i] - nums[i - 1]
				if (diff > 0 and not inc) or (diff < 0 and inc):
					safe = False
					break
				if abs(diff) < 1 or abs(diff) > 3:
					safe = False
					break
			
			return safe
		
	def get_lines(self, target):
		lines = []
		
		with open(target) as file:
			lines = file.readlines()
		
		stripped_lines = [line.strip() for line in lines]
    
		return stripped_lines
  
res = Solver("./input.txt")
p1 = res.part_1()
p2 = res.part_2()
print(p1)
print(p2)

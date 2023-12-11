class Solver:
	def __init__(self, input):
		self.lines = self.get_lines(input)
		
	def get_lines(self, target):
		lines = []
		
		with open(target) as file:
			lines = file.readlines()
		
		stripped_lines = [line.strip() for line in lines]
    
		return stripped_lines
	
	def oasis(self, backwards):
		res = 0

		for line in self.lines:
			nums = [int(val) for val in line.split(" ")]
			extr = self.extrapolate(nums, backwards)

			if backwards:
				res += nums[0] - extr[0]
			else:
				res += nums[-1] + extr[-1]
		
		return res
	
	def extrapolate(self, nums, backwards):
		res = []
		all_zero = True

		for i in range(1, len(nums)):
			diff = nums[i] - nums[i - 1]
			res.append(diff)
			if diff != 0:
				all_zero = False
		
		if all_zero:
			return res
		
		extr = self.extrapolate(res, backwards)

		if backwards:
			front = res[0] - extr[0]
			return [front] + res

		last = res[-1] + extr[-1]
		return res + [last]
			
res = Solver("./input.txt")
print(res.oasis(False))
print(res.oasis(True))

class Solver:
	def __init__(self, input):
		self.lines = self.get_lines(input)
		
	def get_lines(self, target):
		lines = []
		
		with open(target) as file:
			lines = file.readlines()
		
		stripped_lines = [line.strip() for line in lines]
    
		return stripped_lines
  
res = Solver("./input.txt")
print(res.lines)

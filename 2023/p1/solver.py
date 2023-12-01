class Solver:
	def __init__(self, input):
		self.lines = self.get_lines(input)
		
	def get_lines(self, target):
		lines = []
		
		with open(target) as file:
			lines = file.readlines()
		
		stripped_lines = [line.strip() for line in lines]
    
		return stripped_lines
	
	def calib_val(self):
		res = 0

		for line in self.lines:
			nums = []

			for char in line:
				if ord(char) <= ord("9") and ord(char) >= ord("0"):
					nums.append(char)
			
			res += int(nums[0] + nums[-1])
		
		return res
	
	def calib_val_words(self):
		res = 0
		words = {
			"one": "1",
			"two": "2",
			"three": "3",
			"four": "4",
			"five": "5",
			"six": "6",
			"seven": "7",
			"eight": "8",
			"nine": "9"
		}

		for line in self.lines:
			nums = []
			index = 0

			while index < len(line):
				char = line[index]

				if ord(char) <= ord("9") and ord(char) >= ord("0"):
					nums.append(char)
				else:
					end = index + 3
					while end <= len(line) and end <= index + 5:
						if line[index:end] in words:
							nums.append(words[line[index:end]])
							break
						end += 1
				
				index += 1
			
			res += int(nums[0] + nums[-1])
		
		return res
			
  
res = Solver("./input.txt")
print(res.calib_val())
print(res.calib_val_words())

class Solver:
	def __init__(self, input):
		self.lines = self.get_lines(input)
		
	def get_lines(self, target):
		lines = []
		
		with open(target) as file:
			lines = file.readlines()
		
		stripped_lines = [line.strip() for line in lines]
    
		return stripped_lines
	
	def winnings(self):
		res = 0

		for line in self.lines:
			matches = self.get_matches(line)

			if matches != 0:
				res += 1 << (matches - 1)
		
		return res
	
	def scratchcards(self):
		n = len(self.lines)
		card_to_matches = [0] * n
		
		for i in range(n):
			card_to_matches[i] = self.get_matches(self.lines[i])
		
		res = 0
		copies = [1] * n

		for i in range(n):
			copies_of_card = copies[i]
			matches_of_card = card_to_matches[i]
			j = i + 1

			while j < n and j <= i + matches_of_card:
				copies[j] += copies_of_card
				j += 1
			
			res += copies[i]
		
		return res
			
	def get_matches(self, line):
			parts = line.split(": ")
			numbers = parts[1].split(" | ")
			winners = self.convert_nums(numbers[0])
			scratch = self.convert_nums(numbers[1])
			matches = 0

			for num in scratch:
				if num in winners:
					matches += 1
			
			return matches

	def convert_nums(self, numbers):
		res, cur = {}, 0

		for char in numbers:
			if char != " ":
				cur *= 10
				cur += int(char)
			else:
				if cur != 0:
					res[cur] = True
				cur = 0
		
		if cur != 0:
			res[cur] = True
		
		return res

res = Solver("./input.txt")
print(res.winnings())
print(res.scratchcards())

class Solver:
	def __init__(self, input):
		self.lines = self.get_lines(input)
		
	def get_lines(self, target):
		lines = []
		
		with open(target) as file:
			lines = file.readlines()
		
		stripped_lines = [line.strip() for line in lines]
    
		return stripped_lines
	
	def ways_to_beat(self):
		times = self.get_nums(self.lines[0])
		dists = self.get_nums(self.lines[1])
		res = 1

		for i in range(len(times)):
			t = times[i]
			count = 0
			for j in range(1, t):
				speed = j
				travel = speed * (t - j)
				if travel > dists[i]:
					count += 1
			res *= count

		return res
	
	def ways_to_beat_one(self):
		time = self.get_num(self.lines[0])
		dist = self.get_num(self.lines[1])
		cant, i, j = 0, 1, time
		while i < time:
			if i * (time - i) <= dist:
				cant += 1
			else:
				break
			i += 1

		while j >= 0:
			if j * (time - j) <= dist:
				cant += 1
			else:
				break
			j -= 1

		return time - cant
	
	def get_num(self, line):
		cur = 0

		for char in line:
			if ord(char) >= ord("0") and ord(char) <= ord("9"):
				cur *= 10
				cur += int(char)

		return cur
	
	def get_nums(self, line):
		res, cur = [], 0

		for char in line:
			if ord(char) >= ord("0") and ord(char) <= ord("9"):
				cur *= 10
				cur += int(char)
			else:
				if cur != 0:
					res.append(cur)
				cur = 0
		
		res.append(cur)
		return res

res = Solver("./input.txt")
print(res.ways_to_beat())
print(res.ways_to_beat_one())

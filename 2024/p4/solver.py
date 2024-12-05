class Solver:
	def __init__(self, input):
		self.lines = self.get_lines(input)

	def part1(self):
		res = 0
		mat = self.get_mat()
		a = self.conv_v(mat)
		b = self.conv_h(mat)
		c = self.conv_d(mat)
		rot = self.rotate(mat)
		d = self.conv_d(rot)
		for s in [a, b, c, d]:
			print(s)
			for target in ["XMAS", "SAMX"]:
				res += self.contains(s, target)
		return res

	def contains(self, s, target):
		t = len(target)
		res = 0
		for i in range(len(s) - t + 1):
			if s[i:i+t] == target:
				res += 1
		return res
	
	def get_mat(self):
		res = []

		for line in self.lines:
			res.append(list(line))
		
		return res
	
	def conv_h(self, mat):
		res = ""
		m = len(mat)
		for i in range(m):
			res += "".join(mat[i]) + "#"
		return res
	
	def conv_v(self, mat):
		res = ""
		m, n = len(mat), len(mat[0])
		for j in range(n):
			col = ""
			for i in range(m):
				col += mat[i][j]
			res += col + "#"
		return res
	
	def conv_d(self, mat):
		res = ""
		m, n = len(mat), len(mat[0])

		for i in range(m):
			cur = ""
			for j in range(i + 1):
				cur += mat[i - j][j]
			res += cur + "#"
		
		for j in range(1, n):
			cur = ""
			ind = j
			for i in range(m - 1, j - 1, -1):
				cur += mat[i][ind]
				ind += 1

			res += cur + "#"
		
		return res
	
	def rotate(self, mat):
		res = []
		m, n = len(mat), len(mat[0])

		for j in range(n):
			row = []
			for i in range(m - 1, -1, -1):
				row.append(mat[i][j])
			res.append(row)
		
		return res

	def get_lines(self, target):
		lines = []
		
		with open(target) as file:
			lines = file.readlines()
		
		stripped_lines = [line.strip() for line in lines]
    
		return stripped_lines
  
res = Solver("./input.txt")
p1 = res.part1()
print(p1)

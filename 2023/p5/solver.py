class Solver:
	def __init__(self, input):
		self.lines = self.get_lines(input)
		self.mappings = self.get_mappings(self.lines[3:])
		self.seeds = self.stia(self.lines[0].split(": ")[1])

	def get_lines(self, target):
		lines = []
		
		with open(target) as file:
			lines = file.readlines()
		
		stripped_lines = [line.strip() for line in lines]
    
		return stripped_lines

	def get_seed_locations(self):
		seeds = self.seeds.copy()

		for mapping in self.mappings:
			for i in range(len(seeds)):
				seed = seeds[i]

				for cores in mapping:
					des_start = cores[0]
					src_start = cores[1]
					k = cores[2]

					if seed >= src_start and seed < src_start + k:
						diff = seed - src_start
						dest = des_start + diff
						seeds[i] = dest
						break
			
		return seeds
		
	def get_mappings(self, data):
		mappings = []
		cur_mapping = []

		for line in data:
			if len(line) != 0 and line[-1] == ":":
				continue
			
			if len(line) == 0:
				mappings.append(cur_mapping)
				cur_mapping = []
			else:
				cur_mapping.append(self.stia(line))
		
		if len(cur_mapping) != 0:
			mappings.append(cur_mapping)
		
		return mappings
	
	# string to int arr
	def stia(self, s):
		res = []

		for val in s.split(" "):
			res.append(int(val))

		return res
  
res = Solver("./input.txt")
print(min(res.get_seed_locations()))

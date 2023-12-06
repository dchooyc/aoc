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

	def seed_start_end(self):
		res = []

		for i in range(0, len(self.seeds), 2):
			start = self.seeds[i]
			end = start + self.seeds[i + 1]
			res.append([start, end])
		
		return res
	
	def get_seed_by_ranges(self):
		seeds = self.seed_start_end()

		for mapping in self.mappings:
			new_seeds = []

			while len(seeds) > 0:
				seed = seeds.pop()
				seed_start = seed[0]
				seed_end = seed[1]
				overlap = False

				for cores in mapping:
					des_start = cores[0]
					src_start = cores[1]
					k = cores[2]
					src_end = src_start + k
					overlap_start = max(seed_start, src_start)
					overlap_end = min(seed_end, src_end)
					
					if overlap_start < overlap_end:
						new_des_start = des_start + (overlap_start - src_start)
						new_des_end = des_start + (overlap_end - src_start)
						new_seeds.append([new_des_start, new_des_end])
						overlap = True
					
						if overlap_start > seed_start:
							seeds.append([seed_start, overlap_start])
					
						if overlap_end < seed_end:
							seeds.append([overlap_end, seed_end])

						break
				
				if not overlap:
					new_seeds.append(seed)
			
			seeds = new_seeds

		res = seeds[0][0]

		for seed in seeds:
			res = min(res, seed[0])
		
		return res
		
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
print(res.get_seed_by_ranges())

class Solver:
	def __init__(self, input):
		self.lines = self.get_lines(input)
		
	def get_lines(self, target):
		lines = []
		
		with open(target) as file:
			lines = file.readlines()
		
		stripped_lines = [line.strip() for line in lines]
    
		return stripped_lines
	
	def total_winnings(self, use_joker):
		hand_to_bid = {}
		hand_types = [[] for _ in range(7)]
		encoding = self.get_encoding()
		if use_joker:
			encoding = self.get_encoding_joker()

		for line in self.lines:
			parts = line.split(" ")
			hand = parts[0]
			bid = int(parts[1])
			encoded = ""
			tally, max_count = {}, 0

			for card in hand:
				encoded += encoding[card]
				if card in tally:
					tally[card] += 1
				else:
					tally[card] = 1
				max_count = max(max_count, tally[card])
			
			hand_to_bid[encoded] = bid
			ht = self.get_hand_type(tally, max_count)
			if use_joker:
				ht = self.get_hand_type_joker(tally, max_count)
			hand_types[ht].append(encoded)
		
		res, rank = 0, 1

		for hand_type in hand_types:
			hand_type.sort()
			for hand in hand_type:
				res += rank * hand_to_bid[hand]
				rank += 1
		
		return res

	def get_hand_type(self, tally, max_count):
		if max_count == 5:
			return 6
		elif max_count == 4:
			return 5
		elif max_count == 3:
			if len(tally) == 2:
				return 4
			else:
				return 3
		elif max_count == 2:
			if len(tally) == 3:
				return 2
			else:
				return 1
		return 0
	
	def get_hand_type_joker(self, tally, max_count):
		if max_count == 5:
			return 6
		elif max_count == 4:
			if "J" in tally:
				return 6
			return 5
		elif max_count == 3:
			if len(tally) == 2:
				if "J" in tally:
					return 6
				return 4
			else:
				if "J" in tally:
					return 5
				return 3
		elif max_count == 2:
			if len(tally) == 3:
				if "J" in tally:
					if tally["J"] == 2:
						return 5
					return 4
				return 2
			else:
				if "J" in tally:
					return 3
				return 1
		
		if "J" in tally:
			return 1
		
		return 0

	def get_encoding(self):
		cards = ["A", "K", "Q", "J", "T"]
		res = {}

		for i in range(8):
			res[str(i + 2)] = chr(ord("a") + i)
		
		for i in range(5):
			res[cards[i]] = chr(ord("a") + 12 - i)
		
		return res
	
	def get_encoding_joker(self):
		cards = ["A", "K", "Q", "T"]
		res = {}
		res["J"] = "a"

		for i in range(8):
			res[str(i + 2)] = chr(ord("a") + i + 1)
		
		for i in range(4):
			res[cards[i]] = chr(ord("a") + 12 - i)
		
		return res
  
res = Solver("./input.txt")
print(res.total_winnings(False))
print(res.total_winnings(True))
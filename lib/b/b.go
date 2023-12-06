package b

func InRange(target, start, end byte) bool {
	return target >= start && target <= end
}

func IsNum(char byte) bool {
	return InRange(char, '0', '9')
}

func IsLower(char byte) bool {
	return InRange(char, 'a', 'z')
}

func IsUpper(char byte) bool {
	return InRange(char, 'A', 'Z')
}

func IsLetter(char byte) bool {
	return IsLower(char) || IsUpper(char)
}

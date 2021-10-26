package _7_letterCombinations

var allMap = map[byte][]string{
	'2':{"a","b","c"},'3':{"d","e","f"},'4':{"g","h","i"},'5':{"j","k","l"},
	'6':{"m","n","o"},'7':{"p","q","r","s"},'8':{"t","u","v"},'9':{"w","x","y","z"},
}

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	res := allMap[digits[0]]
	for i := 1; i < len(digits); i++ {
		temp := make([]string, 0)
		for j := 0; j < len(res); j ++ {
			for k := 0; k < len(allMap[digits[i]]); k++{
				temp = append(temp, res[j]+allMap[digits[i]][k])
			}
		}
		res = temp
	}

	return res
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	backTrack(0, digits, "", &res)
	return res
}

func backTrack(offset int, digits, item string, res *[]string) {
	if len(item) == len(digits) {
		*res = append(*res, item)
		return
	}
	curr := allMap[digits[offset]]
	for i := 0; i < len(curr); i++ {
		backTrack(offset+1, digits, item+curr[i], res)
	}
}



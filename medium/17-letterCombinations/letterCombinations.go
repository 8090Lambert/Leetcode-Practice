package _7_letterCombinations

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	allMap := map[byte][]string{
		'2':{"a","b","c"},'3':{"d","e","f"},'4':{"g","h","i"},'5':{"j","k","l"},
		'6':{"m","n","o"},'7':{"p","q","r","s"},'8':{"t","u","v"},'9':{"w","x","y","z"},
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
	if len(digits) == 0 {
		return []string{}
	}
	allMap := map[byte][]string{
		'2':{"a","b","c"},'3':{"d","e","f"},'4':{"g","h","i"},'5':{"j","k","l"},
		'6':{"m","n","o"},'7':{"p","q","r","s"},'8':{"t","u","v"},'9':{"w","x","y","z"},
	}
	res := allMap[digits[0]]
	for i := 1; i < len(digits); i++ {
		tmp := make([]string, 0)
		for j := 0; j < len(res); j++ {
			for k := 0; k < len(allMap[digits[i]]); k++ {
				tmp = append(tmp, res[j] + allMap[digits[i]][k])
			}
		}
		res = tmp
	}
	return res
}
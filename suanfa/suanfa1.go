package main

//	func romanToInt(s string) int {
//		var ace int = 0
//		ac := "IVXLCDM"
//		var a int =0
//		var b int =0
//		for i := 0; i < len(s); i++ {
//			if(s[i])
//		}
//		return ace
//	}
func romanToInt(s string) int {
	ret := 0
	i := 0
	length := len(s)
	for ; i < length-1; i++ {
		//fmt.Println(s[i : i+2])
		switch s[i : i+2] {
		case "IV":
			ret += 4
			i += 1
		case "IX":
			ret += 9
			i += 1
		case "XL":
			ret += 40
			i += 1
		case "XC":
			ret += 90
			i += 1
		case "CD":
			ret += 400
			i += 1
		case "CM":
			ret += 900
			i += 1
		default:
			switch s[i] {
			case 'I':
				ret += 1
			case 'V':
				ret += 5
			case 'X':
				ret += 10
			case 'L':
				ret += 50
			case 'C':
				ret += 100
			case 'D':
				ret += 500
			case 'M':
				ret += 1000
			}
		}
	}
	if i == length-1 {
		switch s[i] {
		case 'I':
			ret += 1
		case 'V':
			ret += 5
		case 'X':
			ret += 10
		case 'L':
			ret += 50
		case 'C':
			ret += 100
		case 'D':
			ret += 500
		case 'M':
			ret += 1000
		}
	}
	return ret
}

func main() {

}

// unicode 对字符的判断函数

package main
import "unicode"


func main() {
	var l = unicode.IsLetter('a')
	var d = unicode.IsDigit('a')
	var s = unicode.IsSpace('a')

	print("l: ", l)
	print("d: ", d)
	print("s: ", s)
}
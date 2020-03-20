package problems

import(
	"github.com/alexrv11/algorithms/datastructure"
)

/*
Validates the string contains charcters that close together 
'{[()]}' -> true
'{[]}{()}' -> false
'{{{{{' -> false
']]]]' -> false
*/

type Problem1 struct {
	stack datastructure.StackRune
	text string
}

func (p *Problem1) Validate() bool {
	openSymbolMap := map[rune]rune{ '{':'}', '[':']', '(':')'}
	for _, value := range p.text {
		if _, ok := openSymbolMap[value]; ok { //value is a open symbol
			p.stack.Push(value)
		} else if p.stack.Empty() {// value is close symbol but we don't have element to compare
			return false
		} else {
			if x, _ := openSymbolMap[p.stack.Top()]; value == x {
				p.stack.Pop()
			} else {
				return false
			}
		}
	}
	return p.stack.Empty()
}







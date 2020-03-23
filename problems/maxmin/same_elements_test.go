package maxmin

import(
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSameElementsCase1(t *testing.T)  {
	A := []int{10, 20,100, 100, 200 }
	B := []int{20, 100, 200, 100, 400, 300}
 
	expected := []int{20, 100, 100, 200}
	fmt.Print(SameElements(A, B))
	assert.Equal(t, expected, SameElements(A,B))
 }
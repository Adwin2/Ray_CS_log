//结构体
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts in place.
func Sort() {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(value []int, t *tree) []int {

}

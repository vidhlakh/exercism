package tree

// Node will contain node fot he tree and its children
//struct for output
type Node struct {
	ID       int
	Children []*Node
}

// Record will contain node's ID and its Parents ID
//struct for input
type Record struct {
	ID, Parent int
}

//Build builds the tree
func Build(records Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

}

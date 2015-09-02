package layout

type Iterator struct {
	stack []*Node
	path  []string // name down to current
}

type Node struct {
	name     string
	children []L // pointer to the parent' children
	index    int // index into the parent's child layers
}

type Path []string

func NewIterator(root []L) *Iterator {
	it := Iterator{stack: []*Node{&Node{"$", root, 0}}}
	return &it
}

func (it *Iterator) Empty() bool {
	return len(it.stack) == 0
}

func (it *Iterator) Layout() (ret L) {
	if top := it.top(); top != nil {
		ret = top.children[top.index]
	}
	return
}

func (it *Iterator) Path() (ret Path) {
	if top := it.top(); top != nil {
		curr := top.children[top.index]
		ret = append(it.path, curr.Name)
	}
	return ret
}

// next descend into children, or if no children: advance to a valid sibling
func (it *Iterator) Next() bool {
	if top := it.top(); top != nil {
		curr := top.children[top.index]
		if len(curr.Layers) > 0 {
			it.push(curr)
		} else {
			it.advance(top)
		}
	}
	return !it.Empty()
}

// advance across children and siblings to find a new valid node.
func (it *Iterator) advance(top *Node) {
	for top != nil {
		// first try moving to a new child
		if next := top.index + 1; next < len(top.children) {
			top.index = next
			break
		}
		// out of children, return to parent
		// we will immediately loop around to try a sibling
		top = it.pop()
	}
}

// return the layer list at the top of the stack
func (it *Iterator) top() (ret *Node) {
	if cnt := len(it.stack); cnt > 0 {
		ret = it.stack[cnt-1]
	}
	return ret
}

// push the child list of the passed layer
func (it *Iterator) push(curr L) {
	it.stack = append(it.stack, &Node{curr.Name, curr.Layers, 0})
	it.path = append(it.path, curr.Name)
}

// exit the current layer list
func (it *Iterator) pop() (top *Node) {
	it.stack = it.stack[0 : len(it.stack)-1]
	if t := it.top(); t != nil {
		it.path = it.path[0 : len(it.path)-1]
		top = t
	}
	return top
}

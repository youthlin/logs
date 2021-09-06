package trie

type TireNode struct {
	data interface{}
	chd  map[rune]*TireNode
}

type Tire struct {
	root *TireNode
}

func newNode(data interface{}) *TireNode {
	return &TireNode{data: data, chd: map[rune]*TireNode{}}
}

func NewTire(rootData interface{}) *Tire {
	return &Tire{newNode(rootData)}
}

func (t *Tire) Insert(path string, data interface{}) {
	n := t.root
	for _, r := range path {
		chd, ok := n.chd[r]
		if !ok {
			chd = newNode(nil)
			n.chd[r] = chd
		}
		n = chd
	}
	n.data = data
}

func (t *Tire) Search(path string) interface{} {
	n := t.root
	result := n.data
	for _, r := range path {
		if chd, ok := n.chd[r]; ok {
			n = chd
			if n.data != nil {
				result = n.data
			}
		} else {
			break
		}
	}
	return result
}

func (t *Tire) Dump() map[string]interface{} {
	path := ""
	n := t.root
	result := make(map[string]interface{})
	dump(path, n, result)
	return result
}

func dump(path string, node *TireNode, result map[string]interface{}) {
	if node.data != nil {
		result[path] = node.data
	}
	for r, chd := range node.chd {
		dump(path+string(r), chd, result)
	}
}

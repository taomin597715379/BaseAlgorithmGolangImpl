package rbtree

// Rb tree struct
type TreeNode struct {
	rchild *TreeNode
	lchild *TreeNode
	pchild *TreeNode
	color  string
	data   interface{}
}

type RBTree struct {
	root   *TreeNode
	cur    *TreeNode
	create *TreeNode
}

func NewTreeNode(data interface{}) *TreeNode {
	return &TreeNode{
		color: `red`,
		data:  data,
	}
}

func (rb *RBTree) addNode(data interface{}) {
	rb.create = NewTreeNode(data)
	if !rb.isEmpty() {
		//找到自己的位置
		rb.cur = rb.root
		for {
			if data.(int) < rb.cur.data.(int) {
				if rb.cur.lchild == nil {
					rb.cur.lchild = rb.create
					rb.create.pchild = rb.cur
				} else {
					rb.cur = rb.cur.lchild
				}
			} else if data.(int) > rb.cur.data.(int) {
				if rb.cur.rchild == nil {
					rb.cur.rchild = rb.create
					rb.create.pchild = rb.cur
				} else {
					rb.cur = rb.cur.rchild
				}
			} else {
				return
			}
		}
	} else {
		rb.root = rb.create
		rb.root.color = `black`
		return
	}
	//插入节点之后，需要对红黑树性质进行修复
	rb.insertBalanceFix(rb.create)
}

func (rb *RBTree) deleteNode(data interface{}) {

}

func (rb *RBTree) insertBalanceFix(inode *TreeNode) {
	if inode.color == `red` && inode.pchild.pchild.color == `red` {

	}
}

func (rb *RBTree) deleteBalanceFix(dnode *TreeNode) {

}

func (rb *RBTree) getRoot() *TreeNode {

}

func (rb *RBTree) isEmpty() {

}

func (rb *RBTree) leftRotate(node *TreeNode) {

}

func (rb *RBTree) rightRotate(node *TreeNode) {

}

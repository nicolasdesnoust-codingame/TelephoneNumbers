package domain

type TelephoneNumbersTree struct {
	rootNode *TreeNode
	size     int
}

func NewTelephoneNumbersTree() *TelephoneNumbersTree {
	return &TelephoneNumbersTree{NewTreeNode(-1), 0}
}

func (tree *TelephoneNumbersTree) RegisterTelephoneNumber(telephoneNumber string) {
	insertionPoint := tree.findInsertionPoint(telephoneNumber)
	tree.insertRemainingNumbers(insertionPoint, telephoneNumber)
}

type insertionPoint struct {
	numberIndex int
	node        *TreeNode
}

func (tree *TelephoneNumbersTree) findInsertionPoint(telephoneNumber string) *insertionPoint {
	currentNode := tree.rootNode

	numberIndex := 0
	for ; numberIndex < len(telephoneNumber); numberIndex++ {
		if !currentNode.hasChildWithValue(int(telephoneNumber[numberIndex])) {
			break
		}

		childNode := currentNode.getChildByValue(int(telephoneNumber[numberIndex]))
		currentNode = childNode
	}

	return &insertionPoint{numberIndex: numberIndex, node: currentNode}
}

func (tree *TelephoneNumbersTree) insertRemainingNumbers(insertionPoint *insertionPoint, telephoneNumber string) {
	previousNode := insertionPoint.node
	var currentNode *TreeNode

	for numberIndex := insertionPoint.numberIndex; numberIndex < len(telephoneNumber); numberIndex++ {
		currentValue := int(telephoneNumber[numberIndex])
		currentNode = previousNode.addChildWithValue(currentValue)
		tree.size++

		previousNode = currentNode
	}
}

func (tree *TelephoneNumbersTree) Size() int {
	return tree.size
}

package domain

type TelephoneNumbersTree struct {
	rootNode *TreeNode
	size     int
}

func NewTelephoneNumbersTree() *TelephoneNumbersTree {
	const rootNodeValuePlaceholder = -1
	return &TelephoneNumbersTree{NewTreeNode(rootNodeValuePlaceholder), 0}
}

func (tree *TelephoneNumbersTree) RegisterTelephoneNumber(telephoneNumber *TelephoneNumber) {
	insertionPoint := tree.findInsertionPoint(telephoneNumber)
	tree.insertRemainingNumbers(insertionPoint, telephoneNumber)
}

type insertionPoint struct {
	numberIndex int
	node        *TreeNode
}

func (tree *TelephoneNumbersTree) findInsertionPoint(telephoneNumber *TelephoneNumber) *insertionPoint {
	currentNode := tree.rootNode

	numberIndex := 0
	for ; numberIndex < telephoneNumber.Size(); numberIndex++ {
		number := telephoneNumber.GetNumberByIndex(numberIndex)

		if !currentNode.hasChildWithValue(number) {
			break
		}

		childNode := currentNode.getChildByValue(number)
		currentNode = childNode
	}

	return &insertionPoint{numberIndex: numberIndex, node: currentNode}
}

func (tree *TelephoneNumbersTree) insertRemainingNumbers(insertionPoint *insertionPoint, telephoneNumber *TelephoneNumber) {
	previousNode := insertionPoint.node
	var currentNode *TreeNode

	for i := insertionPoint.numberIndex; i < telephoneNumber.Size(); i++ {
		currentValue := telephoneNumber.GetNumberByIndex(i)
		currentNode = previousNode.addChildWithValue(currentValue)
		tree.size++

		previousNode = currentNode
	}
}

func (tree *TelephoneNumbersTree) Size() int {
	return tree.size
}

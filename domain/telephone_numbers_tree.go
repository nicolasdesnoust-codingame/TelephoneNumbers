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
	currentNode := tree.rootNode

	numberIndex := 0
	for ; numberIndex < telephoneNumber.Size(); numberIndex++ {
		number := telephoneNumber.GetNumberByIndex(numberIndex)

		if !currentNode.hasChildWithValue(number) {
			currentNode.addChildWithValue(number)
			tree.size++
		}

		currentNode = currentNode.getChildByValue(number)
	}
}

func (tree *TelephoneNumbersTree) Size() int {
	return tree.size
}

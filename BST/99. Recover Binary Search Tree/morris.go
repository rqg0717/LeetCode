package main

func recoverTreeMorris(root *TreeNode) {
	var first, second, last, max *TreeNode
	for root != nil {
		if root.Left == nil {
			if last != nil && root.Val <= last.Val {
				if first == nil {
					first = last
				}
				second = root
			}
			last = root
			root = root.Right
		} else {
			// searching left
			max = root.Left
			for max.Right != nil && max.Right != root {
				max = max.Right
			}

			if max.Right != root {
				// connects to root
				max.Right = root
				root = root.Left
			} else {
				// disconnect
				max.Right = nil
				if last != nil && root.Val <= last.Val {
					if first == nil {
						first = last
					}
					second = root
				}
				last = root
				root = root.Right
			}
		}
	}
	first.Val, second.Val = second.Val, first.Val
}

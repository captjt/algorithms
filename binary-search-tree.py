# Code snippet for the binary search tree validation problem in hackerrank
# Link : https://www.hackerrank.com/challenges/ctci-is-binary-search-tree
# TODO: fails at weird edge cases

""" Node is defined as
class node:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None
"""
def tree_max(node):
    if not node:
        return float("-inf")
    maxleft  = tree_max(node.left)
    maxright = tree_max(node.right)
    return max(node.data, maxleft, maxright)

def tree_min(node):
    if not node:
        return float("inf")
    minleft  = tree_min(node.left)
    minright = tree_min(node.right)
    return min(node.data, minleft, minright)

def check_binary_search_tree_(root):
    if not root:
        return True
    if (tree_max(root.left) <= root.data <= tree_min(root.right) and
        check_binary_search_tree_(root.left) and check_binary_search_tree_(root.right)):
        return True
    else:
        return False

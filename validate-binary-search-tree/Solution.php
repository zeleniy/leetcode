<?php

namespace Zeleniy\Leetcode\ValidateBinarySearchTree;

class Solution {

    function isValidBST($root) {

        return $this->validate($root, null, null);
    }

    private function validate($node, $min, $max) {

        if ($node === null) {
            return true;
        }

        if ($min !== null && $node->val <= $min) {
            return false;
        }

        if ($max !== null && $node->val >= $max) {
            return false;
        }

        return $this->validate($node->left, $min, $node->val) &&
            $this->validate($node->right, $node->val, $max);
    }
}

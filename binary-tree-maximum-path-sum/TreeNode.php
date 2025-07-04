<?php

namespace Zeleniy\Leetcode\BinaryTreeMaximumPathSum;

class TreeNode {

    public $val   = null;
    public $left  = null;
    public $right = null;

    function __construct($val = 0, $left = null, $right = null) {

        $this->val   = $val;
        $this->left  = $left;
        $this->right = $right;
    }
}

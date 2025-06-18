<?php

namespace Zeleniy\Leetcode\ValidateBinarySearchTree;

class DataSet
{
    public static function getData() {

        return [
            [true,  null],
            [true,  new TreeNode(0)],
            [true,  new TreeNode(2, new TreeNode(1), new TreeNode(3))],
            [false, new TreeNode(2, new TreeNode(3), new TreeNode(1))],
            [false, new TreeNode(2, new TreeNode(2), new TreeNode(2))],
            [false, new TreeNode(5, new TreeNode(1), new TreeNode(4, new TreeNode(3), new TreeNode(6)))],
            [false, new TreeNode(5, new TreeNode(1), new TreeNode(6, new TreeNode(3), new TreeNode(5)))],
        ];
    }
}
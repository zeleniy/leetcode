<?php

namespace Zeleniy\Leetcode\ValidateBinarySearchTree;

use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase
{
    private Solution $solution;

    public function setUp(): void
    {
        $this->solution = new Solution();
    }

    #[DataProvider('withIsValidBST')]
    function testIsValidBST(bool $answer, $root)
    {
        $this->assertEquals($answer, $this->solution->isValidBST($root));
    }

    public static function withIsValidBST() {

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
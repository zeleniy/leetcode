<?php

namespace Zeleniy\Leetcode\BinaryTreeMaximumPathSum;

use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase
{
    private Solution $solution;

    public function setUp(): void
    {
        $this->solution = new Solution();
    }

    #[DataProvider('withMaxPathSum')]
    function testMaxPathSum(int $answer, TreeNode $root)
    {
        $this->assertEquals($answer, $this->solution->maxPathSum($root));
    }

    public static function withMaxPathSum() {

        return [
            [0, new TreeNode()],
        ];
    }
}
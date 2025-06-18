<?php

namespace Zeleniy\Leetcode\ValidateBinarySearchTree;

use PHPUnit\Framework\Attributes\DataProviderExternal;
use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase
{
    private Solution $solution;

    public function setUp(): void
    {
        $this->solution = new Solution();
    }

    #[DataProviderExternal(DataSet::class, 'getData')]
    function testIsValidBST(bool $answer, $root)
    {
        $this->assertEquals($answer, $this->solution->isValidBST($root));
    }
}
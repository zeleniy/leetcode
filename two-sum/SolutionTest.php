<?php

namespace Zeleniy\Leetcode\TwoSum;

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
    function testTwoSum(array $answer, int $target, array $nums)
    {
        $this->assertEquals($answer, $this->solution->twoSum($nums, $target));
    }
}

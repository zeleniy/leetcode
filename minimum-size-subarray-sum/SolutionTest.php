<?php

namespace Zeleniy\Leetcode\MinimumSizeSubarraySum;

use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase
{
    private Solution $solution;

    public function setUp(): void
    {
        $this->solution = new Solution();
    }

    #[DataProvider('withMinSubArrayLen')]
    function testMinSubArrayLen(int $answer, int $target, array $nums)
    {
        $this->assertEquals($answer, $this->solution->minSubArrayLen($target, $nums));
    }

    public static function withMinSubArrayLen() {

        return [
            [2,  7, [2,3,1,2,4,3]],
            [1,  4, [1,4,4]],
            [0, 11, [1,1,1,1,1,1,1,1]],
            [2, 15, [5,1,3,5,10,7,4,9,2,8]],
        ];
    }
}
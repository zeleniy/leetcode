<?php

namespace Zeleniy\Leetcode\SubarraySumEqualsK;

use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase
{
    private Solution $solution;

    public function setUp(): void
    {
        $this->solution = new Solution();
    }

    #[DataProvider('withContainsDuplicate')]
    function testContainsDuplicate(int $answer, array $numbers, int $k)
    {
        $this->assertEquals($answer, $this->solution->subarraySum($numbers, $k));
    }

    public static function withContainsDuplicate() {

        return [
            [2, [1,1,1], 2],
            [2, [1,2,3], 3],
            [3, [1,-1,0], 0],
            [0, [1], 0],
        ];
    }
}
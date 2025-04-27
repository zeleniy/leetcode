<?php

namespace Zeleniy\Leetcode\ClimbingStairs;

use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase
{
    private Solution $solution;

    public function setUp(): void
    {
        $this->solution = new Solution();
    }

    #[DataProvider('withClimbStairs')]
    function testClimbStairs(int $answer, int $n)
    {
        $this->assertEquals($answer, $this->solution->climbStairs($n));
    }

    public static function withClimbStairs() {

        return [
            [2, 2],
            [3, 3],
        ];
    }
}
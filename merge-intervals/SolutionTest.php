<?php

namespace Zeleniy\Leetcode\MergeIntervals;

use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase
{
    private Solution $solution;

    public function setUp(): void
    {
        $this->solution = new Solution();
    }

    #[DataProvider('withMerge')]
    function testMerge(array $answer, array $intervals)
    {
        $this->assertEquals($answer, $this->solution->merge($intervals));
    }

    public static function withMerge() {

        return [
            [[[1,6],[8,10],[15,18]], [[1,3],[2,6],[8,10],[15,18]]],
            [[[1,5]], [[1,4],[4,5]]],
            [[[0,4]], [[1,4],[0,4]]],
            [[[1,4]], [[1,4],[2,3]]],
            [[[1,10]], [[2,3],[4,5],[6,7],[8,9],[1,10]]],
        ];
    }
}
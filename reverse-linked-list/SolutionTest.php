<?php

namespace Zeleniy\Leetcode\ReverseLinkedList;

use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase
{
    private Solution $solution;

    public function setUp(): void
    {
        $this->solution = new Solution();
    }

    #[DataProvider('withReverseList')]
    function testReverseList(array $answer, array $intervals)
    {
        $this->markTestIncomplete();
        $this->assertEquals($answer, $this->solution->reverseList($intervals));
    }

    public static function withReverseList() {

        return [
            [[1, 2, 3], [3, 2, 1]]
        ];
    }
}
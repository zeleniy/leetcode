<?php

namespace Zeleniy\Leetcode\ContainsDuplicate;

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
    function testContainsDuplicate(bool $answer, array $numbers)
    {
        $this->assertEquals($answer, $this->solution->containsDuplicate($numbers));
    }

    public static function withContainsDuplicate() {

        return [
            [true,  [1,2,3,1]],
            [false, [1,2,3,4]],
            [true,  [1,1,1,3,3,4,3,2,4,2]],
        ];
    }
}
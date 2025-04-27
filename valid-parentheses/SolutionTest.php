<?php

namespace Zeleniy\Leetcode\ValidParentheses;

use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase
{
    private Solution $solution;

    public function setUp(): void
    {
        $this->solution = new Solution();
    }

    #[DataProvider('withIsValid')]
    function testIsValid(bool $answer, string $string)
    {
        $this->assertEquals($answer, $this->solution->isValid($string));
    }

    public static function withIsValid() {

        return [
            [true,  ''],
            [true,  '()'],
            [true,  '()[]{}'],
            [false, '(]'],
            [true,  '([])'],
            [false, '['],
            [false, '(){}}{'],
        ];
    }
}

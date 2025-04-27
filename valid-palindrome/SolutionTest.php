<?php

namespace Zeleniy\Leetcode\ValidPalindrome;

use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase
{
    private Solution $solution;

    public function setUp(): void
    {
        $this->solution = new Solution();
    }

    #[DataProvider('withIsPalindrome')]
    function testIsPalindrome(bool $answer, string $string)
    {
        // $this->markTestSkipped();
        $this->assertEquals($answer, $this->solution->isPalindrome($string));
    }

    public static function withIsPalindrome() {

        return [
            [true,  'a'],
            [true,  'aa'],
            [true,  'ama'],
            [true,  'amma'],
            [true,  'amnma'],
            [true,  'A man, a plan, a canal: Panama'],
            [false, 'race a car'],
            [false, 'good luck'],
            [true,  ' '],
            [true,  ''],
            [true, '.,'],
            [true,  '.'],
        ];
    }
}
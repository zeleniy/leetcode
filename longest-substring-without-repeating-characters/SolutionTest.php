<?php

namespace Zeleniy\Leetcode\LongestSubstringWithoutRepeatingCharacters;

use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase
{
    private Solution $solution;

    public function setUp(): void
    {
        $this->solution = new Solution();
    }

    #[DataProvider('withLengthOfLongestSubstring')]
    function testLengthOfLongestSubstring(int $answer, string $string)
    {
        $this->assertEquals($answer, $this->solution->lengthOfLongestSubstring($string));
    }

    public static function withLengthOfLongestSubstring() {

        return [
            [0, ''],
            [1, 'x'],
            [3, 'pwwkew'],
            [3, 'abcabcbb'],
            [1, 'bbbbb'],
            [7, 'geeksforgeeks'],
            [6, 'abcdefabcbb'],
        ];
    }
}

<?php

namespace Zeleniy\Leetcode\ValidAnagram;

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
    function testContainsDuplicate(bool $answer, string $s1, string $s2)
    {
        $this->assertEquals($answer, $this->solution->isAnagram($s1, $s2));
    }

    public static function withContainsDuplicate() {

        return [
            [true,  'anagram', 'nagaram'],
            [false, 'rat', 'car'],
            [false, 'aacc', 'ccac'],
        ];
    }
}
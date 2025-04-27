<?php

namespace Zeleniy\Leetcode\FirstUniqueCharacterInString;

use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase
{
    private Solution $solution;

    public function setUp(): void
    {
        $this->solution = new Solution();
    }

    #[DataProvider('withFirstUniqChar')]
    function testFirstUniqChar(int $answer, string $string)
    {
        $this->assertEquals($answer, $this->solution->firstUniqChar($string));
    }

    public static function withFirstUniqChar() {

        return [
            [0,  'leetcode'],
            [2,  'loveleetcode'],
            [-1, 'aabb'],
        ];
    }
}
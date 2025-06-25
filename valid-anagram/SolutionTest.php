<?php

namespace Zeleniy\Leetcode\ValidAnagram;

use PHPUnit\Framework\Attributes\DataProviderExternal;
use PHPUnit\Framework\TestCase;

class SolutionTest extends TestCase
{
    private Solution $solution;

    public function setUp(): void
    {
        $this->solution = new Solution();
    }

    #[DataProviderExternal(DataSet::class, 'getData')]
    function testIsAnagramWithCountChars(bool $answer, string $s1, string $s2)
    {
        $this->assertEquals($answer, $this->solution->isAnagramWithCountChars($s1, $s2));
    }

    #[DataProviderExternal(DataSet::class, 'getData')]
    function testIsAnagramWithHashMap(bool $answer, string $s1, string $s2)
    {
        $this->assertEquals($answer, $this->solution->isAnagramWithHashMap($s1, $s2));
    }

    #[DataProviderExternal(DataSet::class, 'getData')]
    function testIsAnagramWithFixedArray(bool $answer, string $s1, string $s2)
    {
        $this->assertEquals($answer, $this->solution->isAnagramWithFixedArray($s1, $s2));
    }
}
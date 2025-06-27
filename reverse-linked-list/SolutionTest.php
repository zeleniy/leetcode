<?php

namespace Zeleniy\Leetcode\ReverseLinkedList;

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
    function testReverseList(array $answer, null|ListNode $head)
    {
        $head = $this->solution->reverseList($head);

        foreach ($answer as $value) {
            $this->assertEquals($value, $head->val);
            $head = $head->next;
        }

        $this->assertNull($head);
    }
}
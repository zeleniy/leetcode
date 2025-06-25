<?php

namespace Zeleniy\Leetcode\ValidAnagram;

/**
 * ./phpvendor/bin/phpbench run ./valid-anagram/ --retry-threshold=5
 * @Revs(1000)
 * @Iterations(5)
 */
class SolutionBench
{
    private Solution $solution;

    public function __construct()
    {
        $this->solution = new Solution();
    }

    function benchIsAnagramWithCountChars()
    {
        foreach (DataSet::getData() as $data) {
            $this->solution->isAnagramWithCountChars($data[1], $data[2]);
        }
    }

    function benchIsAnagramWithHashMap()
    {
        foreach (DataSet::getData() as $data) {
            $this->solution->isAnagramWithHashMap($data[1], $data[2]);
        }
    }

    function benchIsAnagramWithFixedArray()
    {
        foreach (DataSet::getData() as $data) {
            $this->solution->isAnagramWithFixedArray($data[1], $data[2]);
        }
    }
}

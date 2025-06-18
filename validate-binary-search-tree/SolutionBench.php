<?php

namespace Zeleniy\Leetcode\ValidateBinarySearchTree;

/**
 * ./phpvendor/bin/phpbench run validate-binary-search-tree/ --report=default
 */
class SolutionBench
{
    private Solution $solution;

    public function __construct()
    {
        $this->solution = new Solution();
    }

    /**
     * @Revs(1000)
     * @Iterations(5)
     */
    function benchIsValidBST()
    {
        foreach (DataSet::getData() as $data) {
            $this->solution->isValidBST($data[1]);
        }
    }
}
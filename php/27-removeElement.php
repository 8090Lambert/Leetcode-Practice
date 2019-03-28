<?php

/**
 * Leetcode 第27题
 * 
 * O(1)的空间复杂度删除数组指定的元素，并返回数组的长度
 */
function removeElement(&$nums, $val)
{
    $count = 0;
    $len = count($nums);
    for ($i=0; $i < $len; $i++) { 
        if ($nums[$i] != $val) {
            $nums[$count] = $nums[$i];
            $count++;
        }
    }

    return $count;
}


$custom = function (&$nums, $val) {
    $i = 0;
    $len = count($nums);
    while ($i < $len) {
        if ($nums[$i] == $val) {
            $nums[$i] = $nums[$len-1];
            $len--;
        } else {
            $i++;
        }
    }

    return $i;
};


$a = [0, 1, 2, 2, 3, 0, 4, 2];

echo removeElement($a, 2);



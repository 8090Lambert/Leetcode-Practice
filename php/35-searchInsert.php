<?php

/**
 * Leetcode 35题
 * 数组中找到相同元素的位置或是应该插入的位置
 */
function searchInsert($nums, $target)
{
    $count = count($nums);
    if ($nums[0] > $target) {
        return 0;
    } elseif ($nums[$count-1] < $target) {
        return $count;
    } else {
        for ($i=0; $i < $count; $i++) { 
            if ($nums[$i] == $target) {
                return $i;
            } elseif ($nums[$i] > $target) {
                while ($nums[$i] > $target && $i > 1) {
                    $i--;
                }
                return $i;
            }
        }
    }
}


echo searchInsert([1, 3, 5, 6], 2);



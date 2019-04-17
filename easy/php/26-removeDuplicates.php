<?php

function removeDuplicates(&$nums)
{
    $count = count($nums);
    $prev = $nums[0];
    for ($i=1; $i < $count; $i++) {
        $current = $nums[$i];
        if ($prev == $current) {
            unset($nums[$i]);
        } else {
            $prev = $current;
        }
    }

    return count($nums);
}

$a = [0,0,1,1,1,2,2,3,3,4];
removeDuplicates($a);
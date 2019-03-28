<?php

/**
 * Leetcode 28题
 * 实现类似 php 的 strpos
 */
function strStr1($haystack, $needle)
{
    if ($needle == '') {
        return 0;
    }
    
    $flag = false;
    $first = $needle[0];
    $needLength = strlen($needle);
    foreach (str_split($haystack) as $key => $value) {
        if ($value == $first && substr($haystack, $key, $needLength) == $needle) {
            return $flag = $key;
        }
    }

    return $flag;
}


var_dump(strStr1('hello', 'll'));
<?php

/**
 * @param array $strings
 */
function longestCommonPrefix(array $strs) {
    $number = count($strs);
    if ($number == 0) {
        return '';
    }
    if ($number == 1) {
        return $strs[0];
    }

    $first = $strs[0];
    $firstLen = strlen($first);
    for ($i=0; $i < $firstLen; $i++) { 
        $str = $first[$i];
        for ($j=1; $j < $stringsCount; $j++) { 
            if ($strs[$j][$i] != $str) {
                return substr($first, 0, $i);
            }
        }
    }

    return $first;
}


function longestCommonPrefix1(array $strs)
{
    $number = count($strs);
    if ($number == 0) {
        return '';
    }
    if ($number == 1) {
        return $strs[0];
    }

    $first = $strs[0];
    for ($i=1; $i < $number; $i++) {
        while(strpos($strs[$i], $first) !== 0) {
            $first = substr($first, 0, strlen($first) - 1);
            if (!$first) {
                return '';
            }
        }
    }

    return $first;
}

//longestCommonPrefix1(["flower","flow","flight"]);


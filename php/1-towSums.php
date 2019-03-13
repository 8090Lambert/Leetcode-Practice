<?php

/**
 * 两数之和：从集合中找出
 */
function twoSums(array $sums, int $target) {
    $map = [];
    $response = [];
    
    foreach ($sums as $key => $value) {
        # code...
        $need = $target - $value;
        if (isset($map[$need])) {
            $response = array_merge($response, [$key, $map[$need]]);
        }
        $map[$value] = $key;
    }

    return $response;
}
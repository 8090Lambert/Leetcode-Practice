<?php

/**
 * Leetcode 第20题
 */
function isValidWithTwoStack($s)
{
    $orderStack = [];
    $reversedOrderStack = [];
    $strLen = strlen($s);
    $map = ['(' => ')', '[' => ']', '{' => '}'];
    for ($i=0; $i < $strLen; $i++) {
        if (in_array($s[$i], array_keys($map))) {
            if ($s[$i+1] == $map[$s[$i]]) {
                $i = $i + 1;
                continue;
            } else {
                array_push($orderStack, $map[$s[$i]]);
            }
        } else {
            if (!empty($orderStack) && empty($reversedOrderStack) && $s[$i] == $orderStack[count($orderStack)-1]) {
                array_pop($orderStack);
            } else {
                array_unshift($reversedOrderStack, $s[$i]);
            }
        }
    }
    
    return $orderStack == $reversedOrderStack;
}


/**
 * 单个栈
 */
function isValid($s)
{
    $stack = [];
    $strlen = strlen($s);
    $map = ['(' => ')', '[' => ']', '{' => '}'];
    for ($i=0; $i < $strlen; $i++) { 
        if (in_array($s[$i], array_keys($map))) {   // 所有开括号方向
            if ($s[$i+1] == $map[$s[$i]]) {
                $i += 1;
                continue;
            }
            array_unshift($stack, $map[$s[$i]]);
        } else {    // 所有闭括号
            if ($s[$i] != array_slice($stack, 0, 1)[0]) {
                return false;
            } else {
                array_shift($stack);
            }
        }
    }

    return count($stack) == 0;
}

//var_dump(isValid("()[]{}"));
//var_dump(isValid("[({(())}[()])]"));
//var_dump(isValid("([)]"));

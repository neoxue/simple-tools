<?php

$a = file_get_contents("t.txt");
//print_r($a);
preg_match_all("/[\x{4e00}-\x{9fa5}]/u", $a, $chinese);
//print_r($chinese);
echo count($chinese[0]);
echo "\n";

$hash = [];
//print_r(array_keys($chinese));
foreach ($chinese[0] as $k => $v) {
    if (empty($hash[$v])) {
        $hash[$v] = 1;
    } else {
        $hash[$v] += 1;
    }
}
foreach ($hash as $k => $v ) {
    if ($v > 100) {
        print_r($v);
        echo "\t";
        print_r($k);
        echo "\n";
    }
}

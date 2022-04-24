<?php
error_reporting(E_ALL);
ini_set("display_errors", 1);

$url = "http://crmgo.loc:8001/api";



function request($url, $post) {
    $ch=curl_init($url);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($ch, CURLOPT_POSTFIELDS, $post);

    echo curl_exec($ch);
    curl_close($ch);
}

$post = [
    "method" => "getEntry",
    "module" => "auth",
    "rest_data" => json_encode([
        [
            "Name" => "name",
            "Value"   => "alex"
        ],
        [
            "Name" => "password",
            "Value"   => "1234"
        ],
        [
            "Name" => "language",
            "Value"   => "ru"
        ]
    ]),
];
echo "\n" . $post['rest_data'] . "\n";
request($url, $post);





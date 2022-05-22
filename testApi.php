<?php
error_reporting(E_ALL);
ini_set("display_errors", 1);

$url = "http://localhost:8001/api";

class testApi {

    private $url;
    private $token;

    public function __construct($url)
    {
        $this->url = $url;
    }

    public function isAuth()
    {
        return !empty($this->token);
    }

    public function getUserById()
    {
        $post = [
            "method" => "getEntry",
            "module" => "user",
            "token" => $this->token,
            "rest_data" => json_encode([
                [
                    "Name" => "id",
                    "Value" => "a6da9bc8-d619-11ec-9d64-0242ac120002"
                ]
            ]),
        ];
        $result = $this->request($this->url, $post);
        $result = json_decode($result,true);
        $result['result_data'] = json_decode($result['result_data'],true);

    }

    public function auth()
    {
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
        $result = $this->request($this->url, $post);
        $result = json_decode($result,true);
        $result['result_data'] = json_decode($result['result_data'],true);

        if ($result['result_data'][0]['Name'] == 'token') {
            $this->token = $result['result_data'][0]['Value'];
        }
    }

    public function request($url, $post) {
        $ch=curl_init($url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $post);
        $response = curl_exec($ch);
        echo "\n Response: " . $response . "\n";
        curl_close($ch);
        return $response;
    }
}

while (true) {
    $test = new testApi($url);
    $test->auth();
    if ($test->isAuth()) {
        $test->getUserById();
    }
    usleep(1000);
}




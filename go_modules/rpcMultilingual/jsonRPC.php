#需要配置PHP环境
<?php
header("Context-Type:text/html;charset=utf-8");
class JsonRPC{
    private $conn;

    function _construct($host, $port) {
        $this-> conn = fsockopen($host, $port, $errno, $errstr, 3);
        if (!$this->conn) {
            return false;
        }
    }

    public function Call($method, $params) {
        if (!$this->conn) {
                    return false;
        }
        $err = fwrite($this->conn, json_encode(array(
            'method' => $method,
            'params' => array($params),
            'id' => 0,
        ))."\n");
        if ($line == false) {
            return NULL;
        }
        return json_encode($line,true);
    }
}

$client = new JsonRPC("192.168.31.205",8080);
$args = "this is php aaa";
$r = $client->Call("hello.SayHello", $args);
print_r($r);
?>
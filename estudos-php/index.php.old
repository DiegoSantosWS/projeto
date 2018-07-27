<?php

$dsn = "mysql:host=172.25.0.2;port=3306;dbname=cepwsite";
$usr = "root";
$pass = "1234";

try {
    $pdo = new PDO($dsn, $usr, $pass);
    $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);

    $sql = "SELECT * FROM sig_conteudos";
    $res = $pdo->prepare($sql);
    $res->execute();
    
    $bd = $res->fetch(PDO::FETCH_OBJ);
    print_r($bd);

} catch(Exception $e) {
   echo "CONNECTION failed: ".$e->getMessage();
}
phpinfo();
?>

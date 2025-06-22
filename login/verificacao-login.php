<?php

$email = $_GET['ffemail'];
$senha = $_GET['ffsenha'];
$confirmar = $_GET ['confirmar2'];
$nome = "guizin";
if(  $email == $nome  && $senha == 4609 && $confirmar == 4609){
    header("location: ../index.html");
}
else{
    echo "nome de usuario e senha incorretos ";
}
?>
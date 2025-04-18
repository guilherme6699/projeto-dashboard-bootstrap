<?php

$email = $_GET['femail'];
$senha = $_GET['fsenha'];
$confirmar = $_GET ['confirm'];
$nome = "guizin";
if(  $email == $nome  && $senha == 4609 && $confirmar == 4609){
    header("location: ../index.html");
}
else{
    echo "nome de usuario e senha ja existente ";
}
?>
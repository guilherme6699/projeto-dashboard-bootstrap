window.onload = function(){
    let segundos = 0;
    const contadorElement = document.getElementById('contador');

    const intervalo = setInterval(function(){
        segundos++;
    
    const horas = Math.floor(segundos /3600);
    const minutos = Math.floor((segundos % 3600) / 60);
    const segs = segundos % 60;


    const horaFormatado = horas .toString().padStart(2, '0');
    const minutoFormatado = minutos.toString().padStart(2, '0');
    const segundoFormatado = segs.toString().padStart(2, '0');


    contadorElement.textContent = `${horaFormatado}:${minutoFormatado}:${segundoFormatado}`;
    }, 1000);

    window.addEventListener('beforeunload', function(){
        clearInterval(intervalo);
    });
};
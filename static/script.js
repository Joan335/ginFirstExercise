function showCurrentTime() {
    fetch('/api/time')
        .then(response => response.json())
        .then(data => {
            
            const opciones = {
                weekday: 'long',   // día de la semana
                year: 'numeric',
                month: 'long',
                day: 'numeric',
                hour: '2-digit',
                minute: '2-digit',
                second: '2-digit'
            };

            const fecha = new Date(data.time);
            let horaES = fecha.toLocaleString('es-CO', opciones);
            horaES = horaES.charAt(0).toUpperCase() + horaES.slice(1);
            document.getElementById("hora").textContent = horaES;

        })
        .catch(err => console.error("Error al obtener la hora", err));
}

const button = document.getElementById('getResults');

const dom = getCookie("dominio") 
const urlResults = 'http://' + dom + '/v1/results'

button.addEventListener('click', f => {
    f.preventDefault();

    const token = getCookie("Token")
    const ipc = getCookie("IPC_cargado")	
    const terminado = getCookie("JuegoTerminado")
   

    if (terminado === "1" && ipc === "1") {

    fetch(urlResults, {
        method: 'GET',
        headers: {
                'Authorization': token 			
        }
    }).then(async (result) => {

        const data = await result.text()
     
        
        if (result.status === 200) {
    
   
    let dataNew = data.replace("\n", "" )
   
    setCookie("resultados", dataNew, 1)
	
	window.location.assign('./resultsPage')
   

		}  else {
             
            throw  data 
        }
      })    
        .catch(error => {

            Swal.fire({
                toast: true,
                position: "top-right",
                text: "Error: " + error.message,
                timer: 10000,
                showConfirmButton: false
            });
        })
    }else{

        Swal.fire({
            title: "",
            toast: true,
            icon: "warning",
            color: "#716add",
            text: "Aún no estan disponibles los resultados"
    })
    
}

})


      // eliminar cookies
       setCookie("resultado", " ", -1)

function getCurrentDateFormatted() {
    const date = new Date();

    const year = date.getFullYear();                      // 4-digit year
    const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are 0-indexed
    const day = String(date.getDate()).padStart(2, '0');       // Day of the month

    return `${year}-${month}-${day}`;
}



const fechaActual = document.getElementById('actual-prediction-date');
const usuario = document.getElementById('username');

fechaActual.textContent = getCurrentDateFormatted();

const dominio = getCookie("dominio") 
const urlCausas = 'http://' + dominio + '/v1/causas'


const token = getCookie("token")	
const user = getCookie("usuario") 

usuario.textContent = user



// ****** OBTENER LISTADO DE CAUSAS ****** //

const botonConsulta = document.getElementById('consultas');

botonConsulta.addEventListener('submit', e => {
    e.preventDefault();

    const data = new FormData(botonConsulta);
    const obj = {};
    data.forEach((value, key) => obj[key] = value);

let info 

fetch(urlCausas, {
    method: 'POST',
    body: JSON.stringify(obj),
    headers: {     
        'Authorization':token 			
    }
}).then(async (result) => {

    info = await result.text()
    
    if (result.status === 200) {
        
        console.log(info)
        console.log(result.text)
        console.log(result.status)
        //const parsedObject = JSON.parse(info)
        //console.log(parsedObject)
        setCookie("resultado", info, 1)

        window.location.assign('./result')
       

    }  else {
              console.log(info)
              console.log(result.text)
              console.log(result.status)
              const data = result.text
              throw data 
    }
  })    
    .catch(error => {

        Swal.fire({
            toast: true,
            position: "top-right",
            text: "Error: " + error.Msg,
            timer: 10000,
            showConfirmButton: true
        });
    })

})    
// ****** CARGA DE PREDICCION EN BD ****** //

// const results = getCookie("JuegoTerminado")
// const form = document.getElementById('predictionsForm');

// const urlPredictions = 'http://' + dominio + '/v1/predictions'

// form.addEventListener('submit', e => {
//     e.preventDefault();

    
//     const noPrediction = isEmpty('prediction-input', 0)
    
//     if (noPrediction)
//     return 

    
//     if (results === "1" && info.email !== "gastonhch@gmail.com")  {

//         const msg = "Juego Finalizado"
       
//         Swal.fire({
//             toast: true,
            
//             //text:  msg,
//             html: "<h1>" + msg + "</h1>",
//             timer: 20000,
//             showConfirmButton: true
//  //           didOpen: () => {
// //
//  //               const clas = Swal.getPopup().querySelector("div");
//  //               clas. = "class=\"text-sm\" "
// //
//  //           }

//         });

//         return
//     }

    
//     const data = new FormData(form);
	
//     const obj = {};
//     data.forEach((value, key) => obj[key] = value);

//     fetch(urlPredictions, {
//         method: 'POST',
//         body: JSON.stringify(obj),
//         headers: {
//             'Content-Type': 'application/json' ,
//             'Authorization': token 			
//         }
//     }).then(async (result) => {

//         const data = await result.json()
        
//         if (result.status === 200) {
//             Swal.fire({
//                 title: "Predicción cargada con éxito!",
//                 toast: true,
//                 icon: "success",
//                 color: "#716add",
//                 text: data.Msg,

//               }).then((result) => {
//                 if (result.isConfirmed) {
//                   window.location.assign('./index')
//                 }
//               });

          
            
// 		}  else {
     	   
//                   throw await data 
  
//         }
//       })    
//         .catch(error => {

//             Swal.fire({
//                 toast: true,
//                 position: "top-right",
//                 text: "Error: " + error.Msg,
//                 timer: 10000,
//                 showConfirmButton: false
//             });
//         })
//     form.reset();

// })	

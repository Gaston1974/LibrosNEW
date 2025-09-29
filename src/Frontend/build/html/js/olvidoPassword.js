
const form3 = document.getElementById('loginForm');
const button = document.getElementById('botonOlvido');

const domain = getCookie("dominio") 
const urlPassword = 'http://' + domain + '/v1/password'

    button.addEventListener('click', f => {
    f.preventDefault();

    const noEmail = isEmpty('email-input', 0)
    
	if (noEmail) 
		return
 

    const data = new FormData(form3);
    const obj = {};
    data.forEach((value, key) => obj[key] = value);
    
  

  fetch(urlPassword, {
        method: 'POST',
         body: JSON.stringify(obj),
        headers: {
            'Content-Type': 'application/json' ,
            'Authorization': '111' 
        }
    }).then(async (result) => {

        const data = await result.json()
        

        if (result.status === 200) {
            Swal.fire({
                title: "Se enviara una nueva contraseña a su correo registrado",
                toast: true,
                icon: "success",
                color: "#716add",
                text: data.Msg,

              })

		}  else {
      
            throw await data
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
    
})

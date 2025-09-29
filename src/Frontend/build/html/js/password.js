
const form2 = document.getElementById('passwordForm');

const dominio4 = getCookie("dominio") 
const urlPassword = 'http://' + dominio4 + '/v1/password'

form2.addEventListener('submit', f => {
    f.preventDefault();

    const token = getCookie("Token")	

    const noPassword = isEmpty('password-input', 0)
    const noPassword2 = isEmpty('password2-input', 0)

	if (noPassword || noPassword2)
		return

    const data = new FormData(form2);
    const obj = {};
    data.forEach((value, key) => obj[key] = value);
    
        if(obj.Password != obj.Password2) {
        passwordError()
        return
    }
	
  fetch(urlPassword, {
        method: 'POST',
         body: JSON.stringify(obj),
        headers: {
            'Content-Type': 'application/json' ,
            'Authorization': token 			
        }
    }).then(async (result) => {

        const data = await result.json()
        
        if (result.status === 200) {

            Swal.fire({
                title: "listo!",
                toast: true,
                icon: "success",
                color: "#716add",
                customClass: 'swal-wide',
                text: data.Msg,
                showConfirmButton: true
              }).then((result) => {
                if (result.isConfirmed) {
                  window.location.assign('./index')
                }
              });
	
		}  else {
            return data.then(errorData => {
                throw new Error(errorData.error);
            });
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

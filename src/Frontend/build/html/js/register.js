
const form = document.getElementById('registerForm');

const dominio2 = getCookie("dominio") 
const urlUser = 'http://' + dominio2 + '/v1/user'


    form.addEventListener('submit', e => {
    e.preventDefault();

    const noName = isEmpty('nombre-input', 0)
    const noLastName = isEmpty('apellido-input', 0)
    const noPass = isEmpty('password-input', 0)
    const noPass2 = isEmpty('password2-input', 0)

    if (noName || noLastName || noPass || noPass2) return

    const data = new FormData(form);
    const obj = {};
    data.forEach((value, key) => obj[key] = value);

    if(obj.Password != obj.Password2) {
        passwordError()
        return
    }

  
    fetch(urlUser, {
        method: 'POST',
        body: JSON.stringify(obj),
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 27738
        }
    }).then(async (result) => {

        const data = await result.json()
       
        if (result.status === 200) {

            Swal.fire({
                title: "Registro realizado con éxito!",
                icon: "success",
                iconColor: '#30a702',
                confirmButtonColor: '#d33'
              }).then((result) => {
                if (result.isConfirmed) {
                  window.location.assign('./index')
                }
              });
                

        } else {

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
    form.reset();
})


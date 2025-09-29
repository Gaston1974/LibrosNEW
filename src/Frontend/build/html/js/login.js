      // eliminar cookies
       setCookie("token", " ", -1)
     

const dominio3 = getCookie("dominio") 
const urlLogin = 'http://' + dominio3 + '/v1/login'

let res = 200

const form = document.getElementById('loginForm');

form.addEventListener('submit', e => {
    e.preventDefault();

    const noEmail = isEmpty('email-input', 0)
    const noPass = isEmpty('password-input', 0)

    if (noEmail || noPass) 

	return

    const data = new FormData(form);
    const obj = {};
    data.forEach((value, key) => obj[key] = value);


    fetch(urlLogin, {
        method: 'POST',
        body: JSON.stringify(obj),
        headers: {
            'Content-Type': 'application/json' // ,
        //    'Authorization': 27738
                }
    }).then(async (result) => {

        const data = await result.text()
        res = result.status
        let tok = getCookie("token") 
        console.log("tok: " + tok)
        console.log(result.headers.get("Authorization"));

        if (res === 200) {

                  const parsedObject = JSON.parse(data)
                  setCookie("token", parsedObject.Id, 1)
                  setCookie("usuario", parsedObject.Nombre, 1)
          
                 // console.log("json: " + data)
                //   console.log("data : " + result.headers)
                   window.location.assign('./dashboard')
   
               
            } else {
		
	    throw data
        
        }
    })
        .catch(error => {
            if (res === 401)
           err = badValue('password-input', 0)

                Swal.fire({
                toast: true,
                position: "top-right",
                color: "#dc3545",
                text: error ,
                timer: 10000,
                showConfirmButton: true
            }).then((result) => {
                if (result.isConfirmed) {
                    reset('password-input', 0, err)
                  
                };
          
        })
   // form.reset();
  
})

})
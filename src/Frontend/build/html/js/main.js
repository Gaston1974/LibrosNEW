function isEmpty(id, i) {
    const inputDiv = document.getElementById(id)
    const input = inputDiv.getElementsByTagName('input')[i]

    if (input.value.trim() === "") {
        if (inputDiv.querySelector('.error-line')) return true
        const errorLine = document.createElement('span')
        errorLine.innerHTML = 'Información requerida.'
        errorLine.classList.add('error-line', 'inline-flex', 'text-sm', 'text-rose-600')
        inputDiv.appendChild(errorLine)
        input.classList.remove('border-sky-600', 'bg-sky-100')
        input.classList.add('border-rose-600', 'bg-rose-100')
        return true
    }
    const errorLine = inputDiv.querySelector('.error-line')
    if (errorLine) {
        inputDiv.removeChild(errorLine)
        input.classList.remove('border-rose-600', 'bg-rose-100')
        input.classList.add('border-sky-600', 'bg-sky-100')
    }

    return false
}

function badValue(id, i) {
  const inputDiv = document.getElementById(id)
  const input = inputDiv.getElementsByTagName('input')[i]

  
      if (inputDiv.querySelector('.error-line')) return 

      const errorLine = document.createElement('span')
      errorLine.innerHTML = 'contraseña incorrecta.'
      errorLine.classList.add('error-line', 'inline-flex', 'text-sm', 'text-rose-600')
      inputDiv.appendChild(errorLine)
      input.classList.remove('border-sky-600', 'bg-sky-100')
      input.classList.add('border-rose-600', 'bg-rose-100')  

  return errorLine
}

function reset(id, i, err) {
  const inputDiv = document.getElementById(id)
  const input = inputDiv.getElementsByTagName('input')[i]
  
      err.innerHTML = ""
      inputDiv.appendChild(err)
  
}

function getCookie(name) {
  const value = `; ${document.cookie}`;
  console.log("valor: " + value +  " nombre: " + name)
  const parts = value.split(`; ${name}=`);
  
  if (parts.length === 2) 
	return parts.pop().split(';').shift()
	else {
		console.log(parts)
	return value	
	}
}


function setCookie(cname, cvalue, exdays) {
  
  const d = new Date();
  d.setTime(d.getTime() + (exdays*24*60*60*1000));
  let expires = "expires="+ d.toUTCString(); // Session
  document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/v1; SameSite=Strict";

	        
}

function passwordError() {
    const inputDiv = document.getElementById('password2-input')
    const input = inputDiv.getElementsByTagName('input')[0]


        if (inputDiv.querySelector('.error-line')) return
        const errorLine = document.createElement('span')
        errorLine.innerHTML = 'Las contraseñas ingresadas no son iguales'
        errorLine.classList.add('error-line', 'inline-flex', 'text-sm', 'text-rose-600')
        inputDiv.appendChild(errorLine)
        input.classList.remove('border-sky-600', 'bg-sky-100')
        input.classList.add('border-rose-600', 'bg-rose-100')
        return
}


/* 
  setmenu() function is used to set menu items dynamically
  the user to the selected menu option from the frontend.
*/
function setMenu(menuItems) {
  const menuContainer = document.getElementById('menu');
  menuContainer.innerHTML = ''; // Clear existing menu items

  // Loop through menuItems array and create buttons for each item
  menuItems.forEach(item => {
    const button = document.createElement('button');
    button.classList.add('menu');
    button.textContent = item.text;
    button.onclick = () => {
      window.location.href = item.url;
    };
    menuContainer.appendChild(button);
  });
}

/*
  Validate password. Check for at least 8 characters, 
  one number, one uppercase letter, one lowercase letter, 
  and that the password matches the confirm password.
*/
function validatePassword() {
  const password           = document.getElementById('password').value;
  const confirmPassword    = document.getElementById('confirmPassword').value;
  const charLengthCheck    = document.getElementById('charLengthCheck');
  const numberCheck        = document.getElementById('numberCheck');
  const upperCaseCheck     = document.getElementById('upperCaseCheck');
  const lowerCaseCheck     = document.getElementById('lowerCaseCheck');
  const passwordMatchCheck = document.getElementById('passwordMatchCheck');

  // Check password length
  if (password.length >= 8) {
    charLengthCheck.textContent = '✓';
  } else {
    charLengthCheck.textContent = '✗';
  }

  // Check for at least one number
  if (/\d/.test(password)) {
    numberCheck.textContent = '✓';
  } else {
    numberCheck.textContent = '✗';
  }

  // Check for at least one uppercase letter
  if (/[A-Z]/.test(password)) {
    upperCaseCheck.textContent = '✓';
  } else {
    upperCaseCheck.textContent = '✗';
  }

  // Check for at least one lowercase letter
  if (/[a-z]/.test(password)) {
    lowerCaseCheck.textContent = '✓';
  } else {
    lowerCaseCheck.textContent = '✗';
  }

  // Check if password matches confirm password
  if (password === confirmPassword) {
    passwordMatchCheck.textContent = '✓';
  } else {
    passwordMatchCheck.textContent = '✗';
  }
}
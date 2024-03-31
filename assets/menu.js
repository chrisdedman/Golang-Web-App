function setmenu(option) {
  var host             = window.location.origin;
  window.location.href = host + option;
}

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
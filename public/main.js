const passwordManager = (() => {
  const $change_password_form = document.getElementById('password_change');
  const $message = document.getElementById('message');

  let message = "";
  let isError = false;

  $change_password_form.addEventListener("submit", handleSubmit);

  function update() {
    $message.textContent = message;

    if (isError) {
      $message.classList.add('error');
    } else {
      $message.classList.remove('error');
    }
  }

  async function handleSubmit(e) {
    e.preventDefault();

    message = "";

    const user = e.target[0].value;
    const password = e.target[1].value;
    const newPassword = e.target[2].value;
    const newPasswordRepeat = e.target[3].value;

    if (!checkPasswordRepeat(newPassword, newPasswordRepeat)) {
      isError = true;
      message = "Password do not match";
      update();
      return;
    }

    const response = await changePassword({
                      user: user,
                      password: password,
                      newPassword: newPassword
                     });

    isError = response.error;
    message = response.message;

    update();
  }

  function checkPasswordRepeat(password, password_repeat) {
    return password == password_repeat;
  }

  async function changePassword(changePasswordRequest) {
    const option = {
      method: 'POST',
      body: JSON.stringify(changePasswordRequest),
    }
    const res = await fetch("/api/changePassword", option);
    const data = await res.json();

    return data;
  }

})();


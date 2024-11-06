export function initForms() {
  // login form
  const form = document.getElementById("login-form");
  form?.addEventListener("submit", Login);
  // register form
  const registerForm = document.getElementById("register-form");
  registerForm?.addEventListener("submit", Register);
}

export async function Register(e) {
  e.preventDefault();
  const errElem = document.getElementById("error");

  // get the email, username, password, confirm-password
  const email = document.getElementById("register-email")?.value;
  const username = document.getElementById("register-username")?.value;
  const password = document.getElementById("register-password")?.value;
  const confirmPassword = document.getElementById(
    "register-confirm-password"
  )?.value;

  // check if the email, username, password are not empty
  if (!email || !username || !password || !confirmPassword) {
    // style the err
    errElem.style.color = "red";
    errElem.textContent = "Please fill in all fields";
    setTimeout(() => (errElem.textContent = ""), 3000);
    return;
  }
  // check if the password and confirm-password are the same
  if (password !== confirmPassword) {
    // style the err
    errElem.style.color = "red";
    errElem.textContent = "Password not the same";
    setTimeout(() => (errElem.textContent = ""), 3000);
    return;
  }

  let form = new URLSearchParams();
  form.append("email", email);
  form.append("username", username);
  form.append("password", password);
  form.append("confirmPassword", confirmPassword);

  try {
    // send the register request
    const response = await fetch("/api/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded", // Important!
      },
      body: form,
    });
    const data = await response.text();
    if (response.ok) {
      // redirect to /
      window.location.href = "/";
    } else {
      const errElem = document.getElementById("error");
      // style the err
      errElem.style.color = "red";
      errElem.textContent = data;
      setTimeout(() => (errElem.textContent = ""), 3000);
    }
  } catch (err) {
    console.log(err);
  }
}

export async function Login(e) {
  e.preventDefault();

  const username = document.getElementById("Username")?.value;
  const password = document.getElementById("password")?.value;

  let form = new URLSearchParams();
  form.append("username", username);
  form.append("password", password);

  try {
    const response = await fetch("/api/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded", // Important!
      },
      body: form,
    });
    const data = await response.text();
    if (response.ok) {
      // redirect to /
      window.location.href = "/";
    } else {
      const errElem = document.getElementById("error");
      // style the err
      errElem.style.color = "red";
      errElem.textContent = data;
      setTimeout(() => (errElem.textContent = ""), 3000);
    }
  } catch (e) {
    console.log(e);
  }
}

document.addEventListener("DOMContentLoaded", initForms);

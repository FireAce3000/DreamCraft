document.addEventListener("DOMContentLoaded", function () {
  // Code that should be executed when the page is fully loaded.
  console.log("DOM fully loaded and parsed");
});

// Alert for Login.html button
function ShowMyAlert(): void {
  var userName = document.getElementById("name") as HTMLInputElement;
  var userPassword = document.getElementById("password") as HTMLInputElement;

  alert("Login as " + userName.value + " with Password " + userPassword.value);
}

// Send JSON to Server
async function SendDataToGolang() {
  var userName = document.getElementById("name") as HTMLInputElement;
  var userPassword = document.getElementById("password") as HTMLInputElement;

  if (userName.value === "" && userPassword.value === "") {
    alert("FALSE! Name and Password is empty");
  } else if (userName.value === "" || userPassword.value === "") {
    alert("FALSE! Name or Password is empty");
  }
  try {
    // Connent to Server
    await fetch("http://localhost:8080/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        action: "regist",
        userName: userName.value,
        userPassword: userPassword.value,
      }),
    });
  } catch (error) {
    console.error("ERROR (typescript): ", error);
    alert("Error (typescript): ");
  }
}

async function ToDreamCraft() {

  var userName = document.getElementById("name") as HTMLInputElement;
  var userPassword = document.getElementById("password") as HTMLInputElement;

  if (userName.value === "" && userPassword.value === "") {
    alert("FALSE! Name and Password is empty");
  } else if (userName.value === "" || userPassword.value === "") {
    alert("FALSE! Name or Password is empty");
  }
  try {
    // Connent to Server
    await fetch("http://localhost:8080/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        action: "login",
        userName: userName.value,
        userPassword: userPassword.value,
      }),
    });
  } catch (error) {
    console.error("ERROR (typescript): ", error);
    alert("Error (typescript): ");
  }

  alert("Goto DreamCraft")
}
document.addEventListener("DOMContentLoaded", function () {
  // Hier ist der Code, der ausgeführt werden soll, wenn die Seite vollständig geladen ist.
  console.log("DOM fully loaded and parsed");
  // Fügen Sie hier Ihre anderen Funktionen hinzu.
});

// Test alert for Login.html button
function showMyAlert(): void {
  var userName = document.getElementById("name") as HTMLInputElement;
  var userPassword = document.getElementById("password") as HTMLInputElement;

  alert("Login as " + userName.value + " with Password " + userPassword.value);
}

// Check input and send a JSON to Server
function CheckInput(): void {
  var userName = document.getElementById("name") as HTMLInputElement;
  var userPassword = document.getElementById("password") as HTMLInputElement;

  if (userName.value === "" && userPassword.value === "") {
    alert("FALSE! Name and Password is empty");
  } else if (userName.value === "" || userPassword.value === "") {
    alert("FALSE! Name or Password is empty");
  } else {
    sendDataToGolang();
  }
}

// Send JSON to Server
async function sendDataToGolang() {
  try {
    // Create var for the JSON
    var userName = document.getElementById("name") as HTMLInputElement;
    var userPassword = document.getElementById("password") as HTMLInputElement;

    // Connent to Server
    await fetch("http://localhost:8080/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        userName: userName.value,
        userPassword: userPassword.value,
      }),
    });
  } catch (error) {
    console.error("ERROR (typescript): ", error);
    alert("Error (typescript): ");
  }
}

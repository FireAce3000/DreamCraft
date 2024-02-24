document.addEventListener("DOMContentLoaded", function () {
  // Hier ist der Code, der ausgeführt werden soll, wenn die Seite vollständig geladen ist.
  console.log("DOM fully loaded and parsed");
  // Fügen Sie hier Ihre anderen Funktionen hinzu.
});

function showMyAlert(): void {
  var userName = document.getElementById("name") as HTMLInputElement;
  var userPassword = document.getElementById("password") as HTMLInputElement;

  alert("Login as " + userName.value + " with Password " + userPassword.value);
}

async function sendDataToGolang() {
  try {
    var userName = document.getElementById("name") as HTMLInputElement;
    var userPassword = document.getElementById("password") as HTMLInputElement;

    await fetch("http://localhost:8080/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        userName: userName.value,
        userPassword: userPassword.value,
      }),
      // body: JSON.stringify({ userName: "value", userPassword: "value" })
    });
  } catch (error) {
    console.error("ERROR (typescript): ", error);
    alert("Error (typescript): ");
  }
}

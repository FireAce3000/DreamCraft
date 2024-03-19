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
async function RegistToDreamCraft() {
  var userName = document.getElementById("name") as HTMLInputElement;
  var userPassword = document.getElementById("password") as HTMLInputElement;

  if (userName.value === "" && userPassword.value === "") {
    alert("FALSE! Name and Password is empty");
  } else if (userName.value === "" || userPassword.value === "") {
    alert("FALSE! Name or Password is empty");
  } else {
    try {
      // Connect to Server
      const response = await fetch("http://localhost:8080/dreamcraft", {
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
      
      // Ensure that the response is OK
      if (!response.ok) {
        throw new Error('Failed to register');
      }
      
      // Handle the response here if needed
      
    } catch (error) {
      console.error("ERROR (typescript): " + error);
      alert("ERROR (typescript): " + error);
    }
  }
}

async function LoginToDreamCraft() {
  var userName = document.getElementById("name") as HTMLInputElement;
  var userPassword = document.getElementById("password") as HTMLInputElement;

  if (userName.value === "" && userPassword.value === "") {
    alert("FALSE! Name and Password is empty");
  } else if (userName.value === "" || userPassword.value === "") {
    alert("FALSE! Name or Password is empty");
  }
  
  try {
    // Connect to Server
    const response = await fetch("http://localhost:8080/dreamcraft", {
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
    
    // Ensure that the response is OK
    if (!response.ok) {
      throw new Error('Failed to login');
    }
    
    // Handle the response here if needed
    
  } catch (error) {
    console.error("ERROR (typescript): " + error);
    alert("ERROR (typescript): " + error);
  }
}

/* JavaScript for login.html */

document.getElementById("login-submit").addEventListener("click", function () {
    const email = document.querySelector(".user-input").value;
    const password = document.querySelector(".pass-input").value;

    // Perform login process here
    // In a real-world scenario, you would handle user authentication using APIs

    // For demonstration purposes, we'll simply alert the entered credentials
    alert(`Email: ${email}\nPassword: ${password}`);
});

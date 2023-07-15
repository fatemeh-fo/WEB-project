/* JavaScript for signup.html */

document.getElementById("signup-submit").addEventListener("click", function () {
    const email = document.querySelector(".user-input").value;
    const password = document.querySelector(".pass-input").value;

    // Perform sign-up process here
    // In a real-world scenario, you would handle user registration using APIs

    // For demonstration purposes, we'll simply alert the entered credentials
    alert(`Email: ${email}\nPassword: ${password}`);
});

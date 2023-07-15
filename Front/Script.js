document.getElementById("login-btn").addEventListener("click", function () {
    window.location.href = "login.html";
});

document.getElementById("signup-btn").addEventListener("click", function () {
    window.location.href = "signup.html";
});

function handleLogin() {
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
  
    // Perform the validation and authentication here
    // Replace the console.log statements with actual backend validation
  
    if (username === '' || password === '') {
      console.log('Please enter both username and password.');
    } else {
      // Simulate server-side validation
      // In a real-world scenario, make an API call to the backend to verify credentials
      if (username === 'exampleUser' && password === 'examplePassword') {
        console.log('Login successful! Redirecting...');
        // Redirect to the dashboard or user profile page
        window.location.href = 'dashboard.html';
      } else {
        console.log('Invalid username or password.');
      }
    }
  }
  
  document.getElementById("signup-btn").addEventListener("click", function () {
    const username = document.getElementById("username").value;
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    // Check if any of the fields are empty
    if (!username || !email || !password) {
        alert("Please fill all fields.");
    } else {
        // Save the user info (for demonstration purposes, we'll display an alert)
        alert(`Username: ${username}\nEmail: ${email}\nPassword: ${password}`);
    }
});

// For demonstration purposes only; replace this with actual username from the backend
const loggedInUsername = "JohnDoe";

document.addEventListener("DOMContentLoaded", function () {
    const usernamePlaceholder = document.getElementById("username-placeholder");
    usernamePlaceholder.textContent = loggedInUsername;

    document.getElementById("take-quiz-btn").addEventListener("click", function () {
        // Redirect the user to the explore quizzes page (replace "explore.html" with the actual page)
        window.location.href = "explore.html";
    });

    document.getElementById("create-quiz-btn").addEventListener("click", function () {
        // Redirect the user to the create quiz page (replace "create-quiz.html" with the actual page)
        window.location.href = "create-quiz.html";
    });
});



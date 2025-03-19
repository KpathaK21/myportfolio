document.addEventListener("DOMContentLoaded", function () {
    const toggleButton = document.querySelector(".mode-toggle");
    const body = document.body;
    const stars = document.querySelector(".stars");

    // ✅ Check localStorage and apply dark mode if previously enabled
    if (localStorage.getItem("darkMode") === "enabled") {
        body.classList.add("dark-mode");
        if (stars) stars.style.display = "block";
    }

    // ✅ Toggle dark mode on button click
    toggleButton.addEventListener("click", () => {
        body.classList.toggle("dark-mode");

        if (body.classList.contains("dark-mode")) {
            localStorage.setItem("darkMode", "enabled");
            if (stars) stars.style.display = "block";
        } else {
            localStorage.setItem("darkMode", "disabled");
            if (stars) stars.style.display = "none";
        }
    });
});

// Function to load projects dynamically
function loadProjects() {
    fetch("/projects")
        .then(response => response.text())
        .then(html => {
            document.getElementById("content-area").innerHTML = html;
        })
        .catch(error => console.error("Error loading projects:", error));
}

// Function to load contacts dynamically (FIXED SYNTAX)
function loadContacts() {
    document.getElementById("content-area").innerHTML = `
        <div class="contact-info">
            <a href="mailto:kunjpathak74@gmail.com" class="icon-link">
                <i class="fa fa-envelope"></i>
            </a>
            <a href="https://www.linkedin.com/in/kunjpathak/" target="_blank" class="icon-link">
                <i class="fab fa-linkedin"></i>
            </a>
            <a href="tel:+14088632623" class="icon-link">
                <i class="fas fa-phone"></i>
            </a>
        </div>
    `;
}

// Function to load "About Me" dynamically (FIXED SYNTAX)
function loadAbout() {
    document.getElementById("content-area").innerHTML = `
        <div class="about-me">
            <h2>About Me</h2>
            <p>Hello! I'm Kunj Pathak, a software developer with experience in AI, machine learning, and backend systems.</p>
        </div>
    `;
}

&#x20;🚀 Go URL Shortener (Production Style Project)



!\[Go](https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge\&logo=go)

!\[Status](https://img.shields.io/badge/Status-Active-success?style=for-the-badge)

!\[License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)



A fast, lightweight, and simple URL Shortener service built using Go (Golang).  

It converts long URLs into short links and redirects users to the original destination.



\---



&#x20;📌 Overview



This project demonstrates backend development using Go, including:

\- REST API development  

\- HTTP server handling  

\- URL routing \& redirection  

\- File-based persistence (JSON)  

\- Random short code generation  



\---



&#x20;✨ Features



\- 🔗 Convert long URLs into short links  

\- 🔄 Redirect short URLs to original URLs  

\- 💾 Persistent storage using `urls.json`  

\- ⚡ Fast HTTP server using Go `net/http`  

\- 🧠 Random 6-character short code generator  

\- 📦 Lightweight and dependency-free  



\---



🏗️ Project Architecture



Client (Browser / Curl)

↓

Go HTTP Server (net/http)

↓

/shorten API

↓

JSON Storage (urls.json)

↓

Redirect Handler (/abc123 → original URL)



\---



&#x20;📁 Project Structure



go-url-shortener/

├── main.go

├── go.mod

├── urls.json

└── README.md



\---



&#x20;🚀 Getting Started



1️⃣ Clone the repository

git clone https://github.com/Saniya4400/go-url-shortener.git

cd go-url-shortener



&#x20;2️⃣ Run the project

go run main.go



&#x20;3️⃣ Server will start at

http://localhost:8080



\---



&#x20;📡 API Reference



🔹 Create Short URL



POST /shorten



Example:

curl -X POST -d "url=https://github.com" http://localhost:8080/shorten



\---



&#x20;🔹 Response

{

&#x20; "short\_url": "http://localhost:8080/Ab12Cd"

}



\---



&#x20;🔹 Redirect Example



Open in browser:

http://localhost:8080/Ab12Cd



It will redirect to:

https://github.com



\---



&#x20;🔁 Flow



Long URL → Short URL → Redirect → Original Website



\---



&#x20;💡 Future Improvements



\- Database integration (MySQL / MongoDB)  

\- URL expiration system  

\- Click analytics tracking  

\- Frontend UI  

\- Authentication system  



\---



&#x20;👨‍💻 Author



Saniya Pathan  

GitHub: https://github.com/Saniya4400



\---



&#x20;⭐ Support



If you like this project:

⭐ Star the repo  

🍴 Fork it  

🚀 Share it  



\---



Built with ❤️ using Go (Golang)


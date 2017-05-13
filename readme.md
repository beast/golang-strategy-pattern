*A discount model implemented in Golang using strategy design pattern.*
Technology Stack:

* [Golang] (https://golang.org/)  

<img src="https://blog.golang.org/gopher/gopher.png" height="80">

* [Echo] (https://echo.labstack.com/)

<img src="https://echo.labstack.com/images/logo.png" height="50">




##Project structure
```.
├── handler
│   ├── checkout.go
│   └── handler.go
├── main.go
└── readme.md
```

* **handler** business logic and API binding
* **main.go** application entry point and API routes
* **model** contains all DB or model related struct or functions
* **test** are for test coverages
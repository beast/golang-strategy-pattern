*A discount model implemented in Golang using strategy design pattern.*
Technology Stack:

* [Golang] (https://golang.org/)  

<img src="https://blog.golang.org/gopher/gopher.png" height="80">

* [Echo] (https://echo.labstack.com/)

<img src="https://echo.labstack.com/images/logo.png" height="50">

##Project structure
```.
├── handler
│   ├── calculatePrice.go
│   └── handler.go
├── main.go
├── model
│   └── model.go
├── readme.md
├── strategy-pattern
├── test
└── utils
    ├── checkout.go
    └── discount.go
```

* **handler** business logic and API binding
* **main.go** application entry point and API routes
* **model** contains all DB or model related struct or functions
* **strategy-pattern** executable
* **test** are for end-to-end test coverages, unit testing will be added later
* **utils** utility struct or functions
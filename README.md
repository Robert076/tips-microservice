# Tips Microservice 🍀

A multi-container read-only API built with Go and PostgreSQL that communicate via a bridge network. The Go container listens for HTTP GET requests, fetches data from the PostgreSQL database, and returns it to the user.
In order to run the API, do the following:

1. Clone the repository
   
Run the following command to clone the repository to your local machine:

`git clone https://github.com/Robert076/tips-microservice/`

2. Run the containers
   
Use Docker Compose to start the containers:

`docker-compose up`

3. Accesing the API

Once the containers are running, you can either:

 - Open your browser and visit: http://localhost:8080/
 - Or use Postman with a GET request to localhost:8080/


```
.
├── 📖 README.md                     # Project documentation
├── 🐳 docker-compose.yml            # Docker Compose configuration to define services
└── 🏗️ src                           
    ├── 🐋 Dockerfile                # Multi-stage Dockerfile for building the Go container
    ├── 🗃️ db                        
    │   └── 🗄️ db.go                 # Database connection and queries (as of right now, only GET is implemented)
    ├── 📦 go.mod                    # Go module dependencies
    ├── 🔑 go.sum                    # Go module checksums for dependency integrity
    ├── 🚀 main.go                   # Main entry point for the Go application
    ├── 💡 tip                       
    │   └── 📝 tip.go                # Tip structure
    └── 🔧 utils                     
        └── 🛠️ utils.go              # Helper functions used throughout the project
```

Make sure the *Docker daemon* is running before attempting to run `docker-compose up`

You are going to also need a *.env* file, so take this one as an example:
```
POSTGRES_USER="admin"
POSTGRES_PASSWORD="admin"
POSTGRES_DB="db"
HOST="database"
PORT="5432"
```

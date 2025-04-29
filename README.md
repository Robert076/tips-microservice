# Tips Microservice 🍀

A multi-container read-only API built with Go and PostgreSQL that communicate via a bridge network. The Go container listens for HTTP GET requests, fetches data from the PostgreSQL database, and returns it to the user.
In order to run the API, do the following:

1. Clone the repository
   
Run the following command to clone the repository to your local machine:

```bash
git clone https://github.com/Robert076/tips-microservice/
```

2. Run the containers
   
Use Docker Compose to start the containers:

```bash
docker-compose up
```

3. Accesing the API

Once the containers are running, you can either:

 - Open your browser and visit: http://localhost:8080/
 - Or use Postman with a GET request to localhost:8080/


```
.
├── 📖 README.md                     # Project documentation
├── 🐳 docker-compose.yml            # Docker Compose configuration to define services
├── 🚢 kubernetes
│   ├── ⚙️ api-deployment.yml           # Deployment for the api, describes how many replicas (pods) and how the pods will look regarding the API service
│   ├── 🌐 api-service.yml
│   ├── ⚙️ db-deployment.yml
│   ├── 💾 db-pvc.yml
│   ├── 🌐 db-service.yml
│   └── ⚙️ tips-config.yml
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

Make sure the *Docker daemon* is running before attempting to run `docker-compose up`.

You are going to also need a *.env* file in the root of the project, so take this one as an example:
```
POSTGRES_USER="admin"
POSTGRES_PASSWORD="admin"
POSTGRES_DB="db"
HOST="database"
PORT="5432"
```

---

### 🚀 Deploying to Kubernetes (Locally with Minikube)

To run your app inside a local Kubernetes cluster, follow these steps:

1. 🛠 **Install Minikube** (macOS):  
   ```bash
   brew install minikube
   ```

2. 🧰 **Install kubectl** (Kubernetes CLI for macOS):  
   ```bash
   brew install kubernetes-cli
   ```

3. ⚙️ **Start your local cluster**:  
   ```bash
   minikube start
   ```

4. 📡 **List running services**:  
   ```bash
   kubectl get svc
   ```

5. 🌐 **Get the service URL**:  
   Replace `<SERVICE_NAME>` with the name of your service from the previous command.  
   ```bash
   minikube service <SERVICE_NAME> --url
   ```

6. 🧭 **Open the URL** provided in your browser to access the app.



***NOTE: This only works on macOS or Linux because of the multi-stage Dockerfile for the API. You have to modify that if you want to run it on Windows.***

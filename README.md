# Book Store Assignment

## Prerequisites
Ensure you have the following installed:
- Docker
- Minikube
- Kubectl

## Running the Application

1. Start Minikube:
   ```sh
   minikube start
   ```

2. Set Docker to use Minikube:
   ```sh
   eval $(minikube docker-env)
   ```

3. Build the Docker image:
   ```sh
   docker build -t my-book-api-image:latest .
   ```

4. Deploy the application:
   ```sh
   kubectl apply -f deployment.yaml
   ```

5. Expose the service:
   ```sh
   minikube service my-book-api-service
   ```

## Stopping the Application
To stop Minikube:
```sh
minikube stop
```


# Book Store Assignment

## Prerequisites
Ensure you have the following installed:
- Docker
- Minikube
- Kubectl

## Running the Application

1. Start Minikube:
 
   minikube start
 

2. Set Docker to use Minikube:
 
   eval $(minikube docker-env)


3. Build the Docker image:
  
   docker build -t my-book-api-image:latest .


4. Deploy the application:
 
   kubectl apply -f deployment.yaml
  

5. Expose the service:
  
   minikube service my-book-api-service




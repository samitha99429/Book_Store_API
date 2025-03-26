# Book-store Assigment

# should be installed 


   docker
   minikube
   kubectl

# Commands to follow run athe application


minikube start
eval $(minikube docker-env)
docker build -t my-book-api-image:latest .
kubectl apply -f deployment.yaml
minikube service my-book-api-service



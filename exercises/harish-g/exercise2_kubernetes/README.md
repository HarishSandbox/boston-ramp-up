# Containers Exercise

## Part 1: Steps to run a container image for a server program written in go
Note: The server program is exposed on port 8080. The docker file also defines 
the container port to be exposed using `EXPOSE <port>`

### Step 1:
Build an image using `podman build <directory-containing-dockerfile> --tag <tagName>`
Example:  `podman build . --tag server`

### Step 2:
Run the image in a new container using podman run
`podman run -p <host_port>:<container_port> [--name <containerName>] <imageName>`
Example: `podman run -p 8080:80 --name mycontainer <server>`
Note: `podman ps` // list all available containers
podman port --all  // lists port mappings for containers

### Step 3:
Verify your container is listening by navigating to "localhost:8080/<pattern>" on your browser

### Step 4:
Push to quay.io :  `podman push <imageName> <path-to-quay-repo>`ls

## Part 2: Spinning up a cluster and deploying single instance of go server program using the container image created above

### Step 1: Create a cluster using cluster-bot or openshift-install binary

### Step 2: Create a deployment using the container image created in Part 1 (Step 1)
`kubctl create deployment <deployment-name> --image=<path to quay.io>`

### Validation:  Verify that your deployment->pods are up and running
`kubectl get deployments`
`kubectl get pods`
`kubectl get events`
`kubectl describe <resource_type> <resource_name>`  - To look at logs in case of failures

## Part 3: Expose deployment as a Service

### Step 1: Create a service :
`kubectl expose deployment <deplyment_name> --type=LoadBalancer --port=8080`

### Step 2: Verify service is created and do a curl on the external-ip
`kubectl get services`
`curl -i <external-ip>`

### Reference links:
https://kubernetes.io/docs/tutorials/hello-minikube/
https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#creating-a-deployment
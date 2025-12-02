Intro to k8s:

What is Kubernetes (k8s)?

k8s architecture :

Main k8s Components:

- Pods : Smallest unit in k8s, it can contain one or more containers. It provides abstraction layer over containers. Usually 1 application per pod. Each pod has its own IP address. 

- Service : Service is a permanent IP address that can be attach to a pod. 
Note : Pods communicate with each other using services.

Functionalities of Service:
1. Permanent IP address to pods
2. Load balancing (Service can distribute traffic to multiple pods)

- Ingress : Ingress is used to expose HTTP and HTTPS routes from outside the cluster to services within the cluster. Basically they acts as DNS for services.
Minikube and kubectl - Local Set

- ConfigMap : ConfigMap is used to store non-confidential data in key-value pairs. It can be used to store configuration settings for applications running in pods.

- Secret : Secret is used to store confidential data, such as passwords, OAuth tokens, and ssh keys. It provides a way to securely store and manage sensitive information. It stores data in an encoded format (base64).

- Volume : Volume is used to store data that needs to be persisted beyond the lifecycle of a pod. It provides a way to share data between containers in a pod and to store data that needs to be retained even if the pod is deleted or recreated.

- deployment : Deployment is a higher-level abstraction that manages a group of replicas of a pod. It provides declarative updates to pods and replica sets, allowing you to easily scale and manage your applications. Acts as blueprint for pods.

- StatefulSet : StatefulSet is used to manage stateful applications that require stable network identities and persistent storage. 

- Worker Service or Node : Worker nodes are the machines where the actual applications run. They host the pods and provide the necessary resources for running containers. Three process must be installed on every Node.
1st Process : Container Runtime (Docker, containerd)
2nd Process : Kubelet interacts with both - the container and node. Kubelet starts the pod with container inside it. Kublet is responsible for taking resources from node to pod.
3rd Process : Kube-proxy : Kube-proxy is responsible for maintaining network rules on nodes. It enables communication between pods and services. Kube-proxy uses intelligent load balancing to distribute traffic to pods.
 

- How do you interact with this k8s cluster?
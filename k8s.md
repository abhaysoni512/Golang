Intro to k8s:

What is Kubernetes (k8s)? k8s is orchestration framework used to manage containerized applications in a clustered environment. It automates the deployment, scaling, and management of containerized applications.

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
You interact with the k8s cluster using kubectl command line tool. Kubectl allows you to deploy and manage applications on k8s cluster, inspect and manage cluster resources, and view logs.

- What does configuation file look like?
k8s configuration files are written in YAML or JSON format. They define the desired state of the k8s resources, such as pods, services, deployments, etc. The configuration file typically includes metadata, specifications, and status information for the resource.

each configuration file has 5 main sections:

1. apiVersion : Specifies the version of the k8s API that the resource belongs to
2. kind : Specifies the type of resource being created (pod, service, deployment, etc)
3. metadata : Provides information about the resource, such as its name, namespace, labels, and annotations.
4. spec : Defines the desired state of the resource, including its configuration and settings. it includes details such as container images, resource limits, environment variables, and other specifications. it depends on the resource type(kind)
5. status : Provides information about the current state of the resource (usually filled by k8s itself) 

Where we get status data?
- Under controller plane -> etcd

if we have nested spec, outer spec is for the resource itself, inner spec is for pod inside that resource.

Note : metadata contains name and labels. Name is unique for each resource in a namespace. Labels are key-value pairs that are used to organize and select resources.

How pods know which deployment they belong to?
Pods have labels that match the selector defined in the deployment. The deployment uses these labels to identify and manage the pods that it creates.

How to get status of a resource?
You can get the status of a resource using the kubectl get command. For example, to get the status of all pods in a namespace, you can use the command:
kubectl get pods -n <namespace>

what is a namespace?
Namespace is a virtual cluster inside a k8s cluster. It provides a way to divide cluster resources between multiple users or teams. Each namespace has its own set of resources, such as pods, services, and deployments. Namespaces are useful for organizing and managing resources in a multi-tenant environment

Purpose of namespace:
- hearbeats of nodes are sent to specific namespace (kube-system)
-each node has associated lease object in namespace kube-node-lease
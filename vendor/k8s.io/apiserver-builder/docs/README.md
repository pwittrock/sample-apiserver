# Apiserver Builder

Apiserver builder is a collection of libraries and tools to
build Kubernetes native extensions into their own
Kubernetes apiservers.

## Building a simple apiserver

1. Copy the example and find all `ACTION REQUIRED` sections.  Follow instructions.
 - main.go: update import statements and uncomment registration
 - apis/doc.go: set your domain name
 - .../yourapigroup: change package name to match your api group
 - .../yourapiversion: change package name to match your version
 - .../yourapiversion/doc.go: update +conversion-gen with your api group go package
 - .../yourapiversion/types.go: update type names, fields, and comment tags


2. Generate code
  - generate registration code
  - generate deepcopy code
  - generate typeconversion code

3. Build the apiserver main.go
  - go build main.go

## Running the apiserver with delegated auth against minikube

- start minikube and make sure the kubectl is talking to it
- copy `~/.kube/config` to` ~/.kube/auth_config` - this will be used by the api server for delegating auth to minkube
- add a `~/.kube/config` entry for your apiserver, using the minikube user
  - `kubectl config set-cluster` // Set a cluster to talk to https://localhost:9443 with cert /var/run/kubernetes/apiserver.crt
  - `kubectl config set-context` // Use the cluster you created and the minikube user
- run the server with ` ./main --authentication-kubeconfig ~/.kube/auth_config --authorization-kubeconfig ~/.kube/auth_config --client-ca-file /var/run/kubernetes/apiserver.crt  --requestheader-client-ca-file /var/run/kubernetes/apiserver.crt --requestheader-username-headers=X-Remote-User --requestheader-group-headers=X-Remote-Group --requestheader-extra-headers-prefix=X-Remote-Extra- --etcd-servers=http://localhost:2379 --secure-port=9443 --tls-ca-file  /var/run/kubernetes/apiserver.crt  --print-bearer-token`

## Generating docs



## Using apiserver-builder libraries directly (without generating code)
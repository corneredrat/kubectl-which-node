# kubectl-which-node

## Contents
- Description
- Usage
- Examples
- Installation
- Build from source

## Description

This command display the kubernetes nodes in which a workload controller is deployed in. Workload controller can be `replicaset`, `daemonset` , `deployment` or any other standard controller.

All shortforms that kubernetes api server support , such as `rs` for `replicaset` , `ds` for `daemonset` is supported by the command.
The command also supports relavent native kubectl flags.

## Usage
Displays node(s) in which the object(s) is deployed on.

Usage:
```
  kubectl which node [kind] [name] [flags]
```
### Examples:

To display which nodes a `daemonset` is deployed in:

![daemonset-example](assets/images/example-daemonset.PNG)


To display which nodes a `deployment` is deployed in:

![daemonset-example](assets/images/example-deployment.PNG)


To display which nodes a `pod` is deployed in:

![daemonset-example](assets/images/example-pod.PNG)
## Installation
```
# Get the binary
wget https://raw.githubusercontent.com/corneredrat/kubectl-which-node/master/bin/kubectl-which-node

# Change permissions
chmod +x kubectl-which-node

# Add to path
sudo mv ./kubectl-which-node /usr/bin/kubectl-which-node
```

## Build from source
```
# Get code
git clone https://github.com/corneredrat/kubectl-which-node.git

# Run build script
./build
```
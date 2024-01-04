# goingressmonitor
Test assignment for Go and Kubernetes

## Description
1. Luua githubis või Atlassian Bitbucketis repo (meile edastada pärast link või repo zipina).
2. Koostada go operaator-sdk  https://github.com/operator- framework/operator-sdk raamistiku abil näidisoperaator, mis loeb Kubernetese ingress tüüpi CRD-st „host“ ning kirjutab selle operaatori logisse.
3. Lisada juurde „pipeline“, mis kompileerib ja pakendab rakenduse  dockeri konteineriks.
4. Täiendada „pipelinet“ nii et rakendus paigaldatakse mõnda vabalt valitud Kubernetese keskkonda kasutades yaml faile.
5. Lisada näidis CRD giti, mida testimisel  kasutati.

## Solution
1. Created repo in github
## Solution
1. Created a repository on GitHub.
2. Initialized an operator-sdk project with the following command:
   ```bash
   operator-sdk init --domain example.com --repo github.com/madismeer/GoIngressMonitor --plugins=go/v4-alpha
    ```
3. Created an API definition for the IngressMonitor CRD with the following command:
    ```bash
    operator-sdk create api --group network --version v1 --kind Ingress --resource --controller
    ```
    - Edited ingress_types.go to add the Host field to the IngressMonitor CRD
    - Run make manifests to generate the CRD
4. Implemented the reconcile function in ingress_controller.go to log the Host field of the IngressMonitor CRD
    - Edited ingress_controller.go to edit the reconcile function
5. Implemented the pipeline to build and deploy the operator


# Operator SDK Readme

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/goingressmonitor:tag
```

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/goingressmonitor:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller from the cluster:

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

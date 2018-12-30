# k8s-to-application-insights
Push kubernetes cluster info (deployments, restart counters, etc.) to azure application insights


## Setup

```
make docker
helm upgrade --install --debug test1 --set aikey=XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX,monitorNamespace=default,,monitorCluster=minikube,image.tag=git-0fac99d helm/k8s-to-ai
```

## Dev

```
code .
make dev
```

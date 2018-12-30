# k8s-to-application-insights
Push kubernetes cluster info (deployments, restart counters, etc.) to azure application insights


## Setup

```
make docker

cd helm
helm upgrade --install --debug test1 --set aikey=XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX,monitorNamespace=default,monitorCluster=mycluster,image.tag=git-7f1d5cb k8s-to-ai/
```

## Dev

```
code .
make dev
```

# Istio Examples

## Create new cluster

```bash
k3d cluster create istio-cluster -p "8000:30000@loadbalancer" --agents 2
```

To remove:

```bash
k3d cluster delete istio-cluster
```

## Install istioctl

<https://istio.io/latest/docs/setup/getting-started/>

## Install Istio

```bash
istioctl install --set profile=demo
```

Apply labels to inject proxies

```bash
kubectl label namespace default istio-injection=enabled
```

## Dashboards

### Install

```bash
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.13/samples/addons/grafana.yaml ;
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.13/samples/addons/kiali.yaml ;
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.13/samples/addons/prometheus.yaml ;
```

### View

```bash
istioctl dashboard kiali
istioctl dashboard grafana
istioctl dashboard prometheus
```

## Install Fortio pod

```bash
kubectl apply -f ~/apps/istio/istio-1.13.1/samples/httpbin/sample-client/fortio-deploy.yaml
```

## Exec Fortio load tests

```bash
export FORTIO_POD=$(k get po -lapp=fortio -o 'jsonpath={.items[0].metadata.name}') ;
kubectl exec -it $FORTIO_POD -c fortio -- fortio load --timeout 12s -allow-initial-errors -c 2 -qps 0 -n 200 -loglevel Warning http://echo-service:80 ;
```

# Deployments

Concentrate all k8s configurations regarding the API deploy and resources

## Vault

The API secrets are stored and injected by Vault sidecar.

You must install and configure the vault in your k8s cluster, you can do it following the commands or the ref links

Commands:
```shell
kubectl create namespace type-coffee
kubectl config set-context --current --namespace=type-coffee
helm repo add hashicorp https://helm.releases.hashicorp.com
helm install vault hashicorp/vault --set "server.dev.enabled=true"
```

Ref:
https://developer.hashicorp.com/vault/tutorials/kubernetes/kubernetes-sidecar?product_intent=vault#install-the-vault-helm-chart
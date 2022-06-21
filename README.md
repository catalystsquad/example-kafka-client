# example-kafka-client

Examples of kafka clients (Producers and Consumers) for various tech stacks, for use with Kafka tutorials

Spec of what each client codebase should be capable of coming when that's been settled.

## Setup Strimzi

If you don't already have a kubernete cluster, I suggest creating a local one using [KinD](https://kind.sigs.k8s.io/docs/user/quick-start/) or [K3S](https://rancher.com/docs/k3s/latest/en/installation/) and we have an example [kind_cluster.yaml](kind_cluster.yaml) and you can use it like thus: `kind create cluster --config https://raw.githubusercontent.com/catalystsquad/example-kafka-client/main/kind_cluster_config.yaml`

Once you have a cluster, install the Strimzi Operator to it using the following:

```bash
helm repo add strimzi https://strimzi.io/charts/
helm install strimzi-ops strimzi/strimzi-kafka-operator --create-namespace --namespace mykafka -f https://raw.githubusercontent.com/catalystsquad/example-kafka-client/main/strimzi_operator_values.yaml
```

You can of course use any values file you wish, or copy that as a starting point and adjust it.

Everything you need to create a kafka cluster is now accessible via `kubectl apply -f myKafka_cluster.yaml`

This will create a production-like 3 broker and 3 zookeeper node kafka cluster. It will require about 12GB of RAM in your K8s cluster, so configure appropriately. If you are low on resources for this on your local machine, you can attempt to adjust resources lower. Keep in mind that Confluents suggestions for [minimum requirements](https://docs.confluent.io/platform/current/installation/system-requirements.html) are 64GB RAM per broker and 24 CPUs, so this is a heavy system to operate.

You can also reduce the number of replicas, but this will be less useful for tutorial use since you're missing all the distributed pieces to experiment with.

## Download Certs (external)

In order to do anything with this external to the cluster, we'll need to get the user cert, key, and CA cert from the kubernetes cluster. If you are going to use a pod inside the kubernetes cluster or something like Telepresence, you can just mount these secrets as files in the pod.

To download, these are commands that should work if you are in the roo of the repo of this Readme:

```bash
kubectl get secret tutuser -n mykafka -o json | jq -r '.data."ca.crt"' | base64 -d > kafkaca.crt
kubectl get secret tutuser -n mykafka -o json | jq -r '.data."user.crt"' | base64 -d > tutuser.crt
kubectl get secret tutuser -n mykafka -o json | jq -r '.data."user.key"' | base64 -d > tutuser.key
```

Move them around appropriately and make sure you update your environment vars according to where they have moved.
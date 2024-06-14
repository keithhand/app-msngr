# app-msngr

**This project is a WIP and does not work yet.**

Keep the talking with your applications short. Reads an application spec for
accessing common Kubernetes tasks without needing to know deployment specifics.

## Example

You can access controller logs for an `nginx-ingress` deployment with use the
command:

```sh
kubectl logs --context dev-cluster --namespace default \
  -l app.kubernetes.io/component: controller
```

Or, after defining an application with the following spec:

```yaml
...
kubernetes:
  context: dev-cluster
  namespace: default
  logContainers:
    labels:
      - app.kubernetes.io/component: controller
...
```

you can use the `logs` command.

```sh
app logs
```

## Commands

- `app create`: Install an application.
- `app deploy`: Update an application.
- `app logs`: Retrieve logs for the defined container set.
- `app re-init`: Uninstall and reinstall an application.
- `app destroy`: Uninstall an application.

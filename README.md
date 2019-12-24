# Drone plugin for Helm 3

This plugin provides an interface between [Drone](https://drone.io/) and [Helm 3](https://github.com/kubernetes/helm):

* Lint your charts
* Deploy your service
* Delete your service

The plugin is inpsired by [drone-helm](https://github.com/ipedrazas/drone-helm), which fills the same role for Helm 2. It provides a comparable feature-set and the configuration settings are backwards-compatible.

## Example configuration

The examples below give a minimal and sufficient configuration for each use-case. For a full description of each command's settings, see [docs/parameter_reference.md](docs/parameter_reference.md).

### Linting

```yaml
steps:
  - name: lint
    image: pelotech/drone-helm3
    settings:
      helm_command: lint
      chart: ./
```

### Deployment

```yaml
steps:
  - name: deploy
    image: pelotech/drone-helm3
    settings:
      helm_command: upgrade
      chart: ./
      release: my-project
    environment:
      API_SERVER: https://my.kubernetes.installation/clusters/a-1234
      KUBERNETES_TOKEN:
        from_secret: kubernetes_token
```

### Deletion

```yaml
steps:
  - name: uninstall
    image: pelotech/drone-helm3
    settings:
      helm_command: uninstall
      release: my-project
    environment:
      API_SERVER: https://my.kubernetes.installation/clusters/a-1234
      KUBERNETES_TOKEN:
        from_secret: kubernetes_token
```

## Upgrading from drone-helm

The setting names for drone-helm3 are backwards-compatible with those for drone-helm, so the only mandatory step is to update the `image` clause so that drone uses the new plugin.

There are some recommended changes, though:

* If your `service_account` is currently `tiller`, change it to a service account with more restricted permissions.
    * If possible, remove the tiller account entirely, since its superuser status presents a security risk.
* Remove outdated settings that have no effect in drone-helm3:

* `purge` -- this is the default behavior in Helm 3
* `recreate_pods`
* `tiller_ns`
* `upgrade`
* `canary_image`
* `client_only`
* `stable_repo_url`

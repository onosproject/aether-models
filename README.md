<!--
SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>

SPDX-License-Identifier: Apache-2.0
-->

# aether-models
Repository to hold the onos-config models related to Aether

In order to load these models into `onos-config`:

- Build the models Docker images:
```shell
make docker-build
```

- [Optionally] Load the Docker images in your `kind` cluster:
```shell
KIND_CLUSTER_NAME=kind make kind-load
```

- Deploy `onos-config` with these extra values:
```yaml
modelPlugins:
  - name: aether-2x
    image: onosproject/aether-2.x:2.0.0
    port: 5152
  - name: aether-4x
    image: onosproject/aether-4.x:4.0.18
    port: 5153
```
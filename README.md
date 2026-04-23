# Provider Upjet-for-CANCOM

> **Community Project** — This is an unofficial, community-maintained Crossplane provider. It is not affiliated with or endorsed by CANCOM GmbH.

`provider-upjet-for-cancom` is a [Crossplane](https://crossplane.io/) provider built using [Upjet](https://github.com/crossplane/upjet) code generation tools. It exposes XRM-conformant managed resources for the [CANCOM API](https://registry.terraform.io/providers/cancom/cancom/latest/docs), enabling you to manage CANCOM Managed Services Cloud infrastructure declaratively via Kubernetes.

This provider is generated from the official [CANCOM Terraform provider](https://github.com/cancom/terraform-provider-cancom) (MPL-2.0), which is the upstream source of truth for all resource schemas.

## Supported Resources

| Group                                                  | Kind            | Description                                                                                        |
| ------------------------------------------------------ | --------------- | -------------------------------------------------------------------------------------------------- |
| `objectstorage.upjet-for-cancom.crossplane.nvst.cloud` | `StorageBucket` | S3-compatible object storage bucket with configurable availability class (`singleDc` or `multiDc`) |
| `objectstorage.upjet-for-cancom.crossplane.nvst.cloud` | `StorageUser`   | Object storage IAM user with policy-based permissions                                              |

> **Note:** This provider currently covers **Object Storage** resources. Contributions to expose additional CANCOM services (DNS, etc.) as Crossplane managed resources are welcome.

## Getting Started

### Prerequisites

- A running [Crossplane](https://docs.crossplane.io/latest/software/install/) installation (v1.14+)
- A CANCOM API token — obtain one from the [CANCOM portal](https://portal.cancom.io)

### Install the Provider

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-upjet-for-cancom
spec:
  package: ghcr.io/investifytech/provider-upjet-for-cancom:v0.1.0
```

### Configure Credentials

Create a Kubernetes secret with your CANCOM token:

```console
kubectl create secret generic cancom-creds \
  --from-literal=credentials='{"token":"<your-cancom-token>"}' \
  -n crossplane-system
```

The credentials JSON supports the following fields:

| Field              | Required | Description                                                        |
| ------------------ | -------- | ------------------------------------------------------------------ |
| `token`            | Yes      | CANCOM API token                                                   |
| `role`             | No       | CRN of a role to assume (e.g. `crn:123456789012::iam:role:MyRole`) |
| `service_registry` | No       | Service Registry URL for endpoint discovery                        |

Then create a `ProviderConfig`:

```yaml
apiVersion: upjet-for-cancom.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: cancom
spec:
  credentials:
    source: Secret
    secretRef:
      name: cancom-creds
      namespace: crossplane-system
      key: credentials
```

### Example: Create an Object Storage Bucket

```yaml
apiVersion: objectstorage.upjet-for-cancom.crossplane.nvst.cloud/v1alpha1
kind: StorageBucket
metadata:
  name: my-bucket
spec:
  forProvider:
    bucketName: my-unique-bucket-name   # must be globally unique
    availabilityClass: multiDc          # singleDc or multiDc
  providerConfigRef:
    name: cancom
```

### Example: Create an Object Storage User

```yaml
apiVersion: objectstorage.upjet-for-cancom.crossplane.nvst.cloud/v1alpha1
kind: StorageUser
metadata:
  name: my-storage-user
spec:
  forProvider:
    username: svc-myuser
    description: Service account for app X
    permissions: |
      {
        "Statement": [
          {
            "Effect": "Allow",
            "Action": "*",
            "Resource": "*"
          }
        ]
      }
  providerConfigRef:
    name: cancom
```

## Developing

This provider is generated from the [CANCOM Terraform provider](https://github.com/cancom/terraform-provider-cancom) using [Upjet](https://github.com/crossplane/upjet). The upstream Terraform provider is the source of truth for resource schemas.

### Prerequisites

- Go 1.24+
- A running Kubernetes cluster (for `make run`)

### Run the code-generation pipeline

Regenerates all `zz_*` files from the upstream Terraform provider schema:

```console
go run cmd/generator/main.go "$PWD"
```

### Adding a new resource

1. Add or update the resource configuration under `config/`
2. Re-run the generator: `go run cmd/generator/main.go "$PWD"`
3. Verify the generated types under `apis/`

For a detailed walkthrough, see the [Upjet provider generation guide](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md).

### Run against a Kubernetes cluster

```console
make run
```

### Build, push, and install

```console
make all
```

### Build binary only

```console
make build
```

## License

This project is licensed under the **Apache License 2.0** — see the [LICENSE](LICENSE) file for details.

It incorporates resource schemas derived from the [CANCOM Terraform provider](https://github.com/cancom/terraform-provider-cancom), which is licensed under the **Mozilla Public License 2.0 (MPL-2.0)**. These two licenses are compatible: MPL-2.0 is a file-level copyleft license that does not propagate to the Larger Work. Attribution is provided in the [NOTICE](NOTICE) file.

This is a **community project** and is not officially affiliated with or endorsed by CANCOM GmbH.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please open an [issue](https://github.com/InvestifyTECH/provider-upjet-for-cancom/issues).

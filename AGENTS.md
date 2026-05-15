# AGENTS.md

This file provides guidance to coding agents (e.g. Claude Code, claude.ai/code) when working with code in this repository.

## Repository purpose

Go module `kubestash.dev/installer` — Helm charts and supporting tooling for installing [KubeStash](https://kubestash.com/). It also exposes a Go API package (`apis/installer/v1alpha1/`) describing chart values so other KubeStash components can consume strongly typed installation parameters.

Charts shipped:
- `charts/kubestash` — meta chart (installs operator + metrics + catalog).
- `charts/kubestash-operator` — the KubeStash operator (the controller + webhooks).
- `charts/kubestash-catalog` — bundled addon catalog (the `Addon` and `Function` CRs that map engines to addon images).
- `charts/kubestash-metrics` — metrics exporter.
- `charts/kubestash-certified` — Red Hat-certified packaging of the operator.
- `charts/kubestash-certified-crds` — CRDs only (Red Hat-certified packaging).

## Architecture

- `charts/` — one subdirectory per Helm chart. Each has `Chart.yaml`, `values.yaml`, `templates/`, plus generated artifacts `doc.yaml`, `README.md`, and (where applicable) `values.openapiv3_schema.yaml` and `crds/`.
- `apis/installer/v1alpha1/` — Go types backing chart values. Used for OpenAPI/values-schema generation and as a typed surface for downstream programs. Single API group: `installer:v1alpha1`.
  - `register.go`, `install/`, `fuzzer/` — standard k8s scheme registration and round-trip fuzz helpers.
- `crds/` — top-level CRDs (chart-specific CRDs live under each chart's `crds/`).
- `catalog/` — image catalog driven by `kmodules.xyz/image-packer`. `imagelist.yaml` is the source of truth; `copy-images.sh`, `export-images.sh`, `import-images.sh`, `import-into-k3s.sh` are generated. `catalog/raw/` holds the source manifests. `catalog/README.md` is the auto-generated CVE report.
- `hack/scripts/` — codegen/release helpers (`update-catalog.sh`, `update-chart-dependencies.sh`, `import-crds.sh`, `ct.sh`, `open-pr.sh`, `update-release-tracker.sh`).
- `tests/` — chart-testing (`ct`) configuration.
- `DEVELOPMENT.md` — developer guide.
- `vendor/` — checked-in deps.

## Common commands

All Make targets run inside `ghcr.io/appscode/golang-dev` — Docker must be running.

- `make gen` — full regen: clientset + manifests + values schemas + chart docs. Run after any change to `apis/installer/v1alpha1/*_types.go`.
- `make clientset` — regenerate Go client only.
- `make manifests` — `gen-crds gen-values-schema gen-chart-doc`.
- `make gen-values-schema` — regenerate `values.openapiv3_schema.yaml` from the Go types.
- `make gen-chart-doc` — regenerate per-chart `README.md` (one target per chart subdir under `charts/`).
- `make update-charts` — refresh chart-level metadata.
- `make fmt`, `make lint`, `make unit-tests` / `make test` — standard.
- `make ct` — chart-testing lint+test.
- `make verify` — `verify-modules`; `go mod tidy && go mod vendor` must leave the tree clean.
- `make add-license` / `make check-license` — manage license headers.

Auxiliary helpers (invoked outside Make):

- `./hack/scripts/update-catalog.sh` — regenerate `catalog/` from `imagelist.yaml`.
- `./hack/scripts/import-crds.sh` — pull CRDs from sibling KubeStash repos into each chart's `crds/`.
- `./hack/scripts/update-chart-dependencies.sh` — refresh `Chart.lock` / subchart pins.

Run a single Go test (requires a local Go toolchain):

```
go test ./apis/installer/v1alpha1/... -run TestName -v
```

## Conventions

- Module path is `kubestash.dev/installer` (vanity URL). Imports must use that.
- Edit `apis/installer/v1alpha1/*_types.go` to change a chart's values surface, then run `make gen` so the generated clientset, `values.openapiv3_schema.yaml`, and per-chart `README.md` stay in sync. Do not hand-edit `zz_generated.*.go`, generated chart `README.md` files, `values.openapiv3_schema.yaml`, or anything under `catalog/` except `imagelist.yaml` and `catalog/raw/`.
- License: `LICENSE.md` (AppsCode); use `make add-license` to apply headers to new Go files.
- Sign off commits (`git commit -s`); contributions follow the DCO (`DCO`, `CONTRIBUTING.md`).
- Vendor directory is checked in — `go mod tidy && go mod vendor` must leave the tree clean (enforced by `verify-modules`).
- CRDs for charts are **imported** from the upstream `kubestash.dev/apimachinery` repo via `import-crds.sh`; do not hand-author `charts/*/crds/*.yaml`.
- The `-certified` charts mirror the standard charts for Red Hat certification — bumps must apply to both pairs (`kubestash-operator` ↔ `kubestash-certified`, top-level CRDs ↔ `kubestash-certified-crds`).

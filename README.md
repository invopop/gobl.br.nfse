# GOBL ➡️ Brazil NFS-e

Brazil NFS-e (service invoice) addon for [GOBL](https://github.com/invopop/gobl).

Released under the Apache 2.0 [LICENSE](https://github.com/invopop/gobl.br.nfse/blob/main/LICENSE), Copyright 2026 [Invopop S.L.](https://invopop.com).

[![Lint](https://github.com/invopop/gobl.br.nfse/actions/workflows/lint.yaml/badge.svg)](https://github.com/invopop/gobl.br.nfse/actions/workflows/lint.yaml)
[![Test Go](https://github.com/invopop/gobl.br.nfse/actions/workflows/test.yaml/badge.svg)](https://github.com/invopop/gobl.br.nfse/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/invopop/gobl.br.nfse)](https://goreportcard.com/report/github.com/invopop/gobl.br.nfse)
[![codecov](https://codecov.io/gh/invopop/gobl.br.nfse/graph/badge.svg)](https://codecov.io/gh/invopop/gobl.br.nfse)
[![GoDoc](https://godoc.org/github.com/invopop/gobl.br.nfse?status.svg)](https://godoc.org/github.com/invopop/gobl.br.nfse)
![Latest Tag](https://img.shields.io/github/v/tag/invopop/gobl.br.nfse)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/invopop/gobl.br.nfse)

This module implements Brazil's **NFS-e** service invoicing documents as a GOBL
tax addon:

- **NFS-e** (`br-nfse-v1`) — Brazil's NFS-e 1.X service invoices.

Brazil's goods/consumer invoices (NF-e / NFC-e) live in a separate module,
[`gobl.br.nfe`](https://github.com/invopop/gobl.br.nfe).

Unlike the format converters in the GOBL ecosystem, this is a true **addon**: it
registers extensions, normalizers, and validation rules into GOBL's global
registry. It lives in its own module so that only projects handling Brazil NFS-e
documents take on its weight.

## Layout

- `addon/` — the GOBL addon: extensions, normalizers, and validation rules that
  register into GOBL on import. This package is kept dependency-light so
  importing it never pulls in conversion tooling.
- the module root (and future subpackages) is reserved for converters and other
  Brazil NFS-e tooling that build on the addon.

## Usage

Add a blank import of the **addon** so it registers itself, then use GOBL as
normal:

```go
import (
	"github.com/invopop/gobl"
	_ "github.com/invopop/gobl.br.nfse/addon"
)
```

Declare the addon on a document (or let the regime/scenario add it) and
`Calculate` + `Validate` will run the full Brazil NFS-e normalization and rules.

> **Note**: the `br-nfse-v1` key is listed in GOBL core's approved external-addon
> registry, so it is recognised as a valid `$addons` value in the JSON Schema.
> The runtime check stays strict, however: a document declaring this addon will
> fail validation with `add-on must be registered` unless this module is
> imported. Any service that processes Brazil NFS-e documents must import it.

## Development

The addon builds on core GOBL features (the `regimes/br` regime and the approved
external-addon registry) that are not yet in a tagged release. The `go.mod`
therefore pins `github.com/invopop/gobl` to a commit on the core
`extract-br-addons` branch (a pseudo-version); bump it to the release tag once
core is published.

```sh
go test ./...
```

### Examples

`examples/` holds sample documents with their expected JSON envelopes under
`examples/out/`. They are verified via GOBL's shared `pkg/examples` helpers.
Regenerate the golden output after intentional changes with:

```sh
go test . -run TestExamples -update
```

## License

Apache 2.0 — see [LICENSE](./LICENSE).

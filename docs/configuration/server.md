---
layout: default
title: Server
parent: Configuration
nav_order: 9
---

The server section configures and tunes the http server module Authelia uses.

## Configuration

```yaml
server:
  read_buffer_size: 4096
  write_buffer_size: 4096
  path: ""
  enable_pprof: false
  enable_expvars: false
```

## Options

### read_buffer_size

<div markdown="1">
type: integer
{: .label .label-config .label-purple }
default: 4096
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

Configures the maximum request size. The default of 4096 is generally sufficient for most use cases.

### write_buffer_size

<div markdown="1">
type: integer
{: .label .label-config .label-purple }
default: 4096
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

Configures the maximum response size. The default of 4096 is generally sufficient for most use cases.

### path

<div markdown="1">
type: string
{: .label .label-config .label-purple }
default: ""
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

Authelia by default is served from the root `/` location, either via its own domain or subdomain.

Modifying this setting will allow you to serve Authelia out from a specified base path. Please note
that currently only a single level path is supported meaning slashes are not allowed, and only
alphanumeric characters are supported.

Example: <https://auth.example.com/>, <https://example.com/>

```yaml
server:
  path: ""
```

Example: <https://auth.example.com/authelia/>, <https://example.com/authelia/>

```yaml
server:
  path: authelia
```

### enable_pprof

<div markdown="1">
type: boolean
{: .label .label-config .label-purple }
default: false
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

Enables the go pprof endpoints.

### enable_expvars

<div markdown="1">
type: boolean
{: .label .label-config .label-purple }
default: false
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

Enables the go expvars endpoints.

## Additional Notes

### Buffer Sizes

The read and write buffer sizes generally should be the same. This is because when Authelia verifies
if the user is authorized to visit a URL, it also sends back nearly the same size response as the request. However
you're able to tune these individually depending on your needs.

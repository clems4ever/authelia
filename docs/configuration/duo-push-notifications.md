---
layout: default
title: Duo Push Notifications
parent: Configuration
nav_order: 2
---

Authelia supports mobile push notifications relying on [Duo].

Follow the instructions in the dedicated [documentation](../features/2fa/push-notifications.md)
to know how to set up push notifications in Authelia.

**Note:** The configuration options in the following sections are noted as required. They are however only required when
you have this section defined. i.e. if you don't wish to use the [Duo] push notifications you can just not define this
section of the configuration.

## Configuration

The configuration is as follows:

```yaml
duo_api:
  hostname: api-123456789.example.com
  integration_key: ABCDEF
  secret_key: 1234567890abcdefghifjkl
```

The secret key is shown as an example, you also have the option to set it using an environment
variable as described [here](./secrets.md).

## Options

### hostname

<div markdown="1">
type: string
{: .label .label-config .label-purple }
default: ""
{: .label .label-config .label-blue }
required: yes
{: .label .label-config .label-red }
</div>

The [Duo] API hostname supplied by [Duo].

### integration_key

<div markdown="1">
type: string
{: .label .label-config .label-purple }
default: ""
{: .label .label-config .label-blue }
required: yes
{: .label .label-config .label-red }
</div>

The non-secret [Duo] integration key. Similar to a client identifier.

### secret_key

<div markdown="1">
type: string
{: .label .label-config .label-purple }
default: ""
{: .label .label-config .label-blue }
required: yes
{: .label .label-config .label-red }
</div>

The secret [Duo] key used to verify your application is valid.

[Duo]: https://duo.com/

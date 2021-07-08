---
layout: default
title: Configuration
nav_order: 4
has_children: true
---

Authelia uses a YAML file as configuration file. A template with all possible options can be
found [here](https://github.com/authelia/authelia/blob/master/config.template.yml), at the root of the repository.

When running **Authelia**, you can specify your configuration by passing the file path as shown below.

```shell
authelia --config config.custom.yml
```

## Documentation

We document the configuration in two ways:

1.  The configuration yaml default has comments documenting it. All documentation lines start with `##`. Lines starting
   with a single `#` are yaml configuration options which are commented to disable them or as examples.

2.  This documentation site. Generally each section of the configuration is in its own section of the documentation
   site. Each configuration option is listed in its relevant section as a heading, under that heading generally are two
   or three colored labels.

-   The `type` label is purple and indicates the yaml value type of the variable. It optionally includes some
     additional information in parentheses.
-   The `default` label is blue and indicates the default value if you don't define the option at all. This is not the
     same value as you will see in the examples in all instances, it is the value set when blank or undefined.
-   The `required` label changes color. When required it will be red, when not required it will be green, when the
     required state depends on another configuration value it is yellow.  

## Validation

Authelia validates the configuration when it starts. This process checks multiple factors including configuration keys
that don't exist, configuration keys that have changed, the values of the keys are valid, and that a configuration
key isn't supplied at the same time as a secret for the same configuration option.

You may also optionally validate your configuration against this validation process manually by using the validate-config
option with the Authelia binary as shown below. Keep in mind if you're using [secrets](./secrets.md) you will have to
manually provide these if you don't want to get certain validation errors (specifically requesting you provide one of
the secret values). You can choose to ignore them if you know what you're doing. This command is useful prior to
upgrading to prevent configuration changes from impacting downtime in an upgrade. This process does not validate
integrations, it only checks that your configuration syntax is valid.

```shell
authelia validate-config configuration.yml
```

## Duration Notation Format

We have implemented a string based notation for configuration options that take a duration. This section describes its
usage. You can use this implementation in: session for expiration, inactivity, and remember_me_duration; and regulation
for ban_time, and find_time. This notation also supports just providing the number of seconds instead.

The notation is comprised of a number which must be positive and not have leading zeros, followed by a letter
denoting the unit of time measurement. The table below describes the units of time and the associated letter.

|Unit   |Associated Letter|
|:-----:|:---------------:|
|Years  |y                |
|Months |M                |
|Weeks  |w                |
|Days   |d                |
|Hours  |h                |
|Minutes|m                |
|Seconds|s                |

Examples:

- 1 hour and 30 minutes: 90m
- 1 day: 1d
- 10 hours: 10h

## TLS Configuration

Various sections of the configuration use a uniform configuration section called TLS. Notably LDAP and SMTP.
This section documents the usage.

### Server Name

<div markdown="1">
type: string
{: .label .label-config .label-purple }
default: ""
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

The key `server_name` overrides the name checked against the certificate in the verification process. Useful if you
require to use a direct IP address for the address of the backend service but want to verify a specific SNI.

### Skip Verify

<div markdown="1">
type: boolean
{: .label .label-config .label-purple }
default: false
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

The key `skip_verify` completely negates validating the certificate of the backend service. This is not recommended,
instead you should tweak the `server_name` option, and the global option [certificates_directory](./miscellaneous.md#certificates_directory).

### Minimum Version

<div markdown="1">
type: string
{: .label .label-config .label-purple }
default: TLS1.2
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

The key `minimum_version` controls the minimum TLS version Authelia will use when opening TLS connections.
The possible values are `TLS1.3`, `TLS1.2`, `TLS1.1`, `TLS1.0`. Anything other than `TLS1.3` or `TLS1.2`
are very old and deprecated. You should avoid using these and upgrade your backend service instead of decreasing
this value.

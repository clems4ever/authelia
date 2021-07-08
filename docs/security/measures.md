---
layout: default
title: Security Measures
parent: Security
nav_order: 1
---

The Authelia team takes security seriously. As such there are several measures we take to ensure the security of the
users who utilize Authelia are protected.

## Protection against cookie theft

Authelia sets several key cookie attributes to prevent cookie theft:

1. `HttpOnly` is set forbidding client-side code like javascript from access to the cookie.
2. `Secure` is set forbidding the browser from sending the cookie to sites which do not use the https scheme.
3. `SameSite` is by default set to `Lax` which prevents it being sent over cross-origin requests.

Read about these attributes in detail on the
[MDN](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie).

## Protection against multi-domain cookie attacks

Since Authelia uses multi-domain cookies to perform single sign-on, an attacker who poisoned a user's DNS cache can
easily retrieve the user's cookies by making the user send a request to one of the attacker's IPs.

This is technically mitigated by the `Secure` attribute set in cookies by Authelia, however it's still advisable to
only use HTTPS connections with valid certificates and enforce it with HTTP Strict Transport Security ([HSTS]) which
will prevent domains from serving over HTTP at all as long as the user has visited the domain before. This means even
if the attacker poisons DNS they are unable to get modern browsers to connect to a compromised host unless they can also
obtain the certificate.

Note that using [HSTS] has consequences. That's why you should read the blog post nginx has written on [HSTS].

## Protection against username enumeration

Authelia adaptively delays authentication attempts based on the mean (average) of the previous 10 successful attempts
in addition to a small random interval of time. The result of this delay is that it makes it incredibly difficult to
determine if the unsuccessful login was the result of a bad password, a bad username, or both. The random interval of
time is anything between 0 milliseconds and 85 milliseconds.

When Authelia first starts it assumes the last 10 attempts took 1000 milliseconds each. As users login successfully it
quickly adjusts to the actual time the login attempts take. This process is independent of the login backend you have
configured.

The cost of this is low since in the instance of a user not existing it just stops processing the request to delay the
login. Lastly the absolute minimum time authentication can take is 250 milliseconds. Both of these measures also have
the added effect of creating an additional delay for all authentication attempts increasing the time that a brute-force
attack will take, this combined with regulation greatly delays brute-force attacks and the effectiveness of them in
general.

## Protections against password cracking (File authentication provider)

Authelia implements a variety of measures to prevent an attacker cracking passwords if they somehow obtain the file used
by the file authentication provider, this is unrelated to LDAP auth.

First and foremost Authelia only uses very secure hashing algorithms with sane and secure defaults. The first and
default hashing algorithm we use is Argon2id which is currently considered the most secure hashing algorithm. We also
support SHA512, which previously was the default.

Secondly Authelia uses salting with all hashing algorithms. These salts are generated with a random string generator,
which is seeded every time it's used by a cryptographically secure 1024bit prime number. This ensures that even if an
attacker obtains the file, each password has to be brute forced individually.

Lastly Authelia's implementation of Argon2id is highly tunable. You can tune the key length, salt used, iterations
(time), parallelism, and memory usage. To read more about this please read how to
[configure](../configuration/authentication/file.md) file authentication.

## User profile and group membership always kept up-to-date (LDAP authentication provider)

This measure is unrelated to the File authentication provider.

Authelia by default refreshes the user's profile and membership every 5 minutes. This ensures that if you alter a users
groups in LDAP that their new groups are obtained relatively quickly in order to adjust their access level for
applications secured by Authelia.

Additionally, it will invalidate any session where the user could not be retrieved from LDAP based on the user filter,
for example if they were deleted or disabled provided the user filter is set correctly. These updates occur when a user
accesses a resource protected by Authelia. This means you should ensure disabled users or users with expired passwords
are not obtainable using the LDAP filter, the default filter for Active Directory implements this behaviour.
LDAP implementations vary, so please ask if you need some assistance in configuring this.

These protections can be [tuned](../configuration/authentication/ldap.md#refresh-interval) according to your security
policy by changing refresh_interval, however we believe that 5 minutes is a fairly safe interval.

## Notifier security measures (SMTP)

The SMTP Notifier implementation does not allow connections that are not secure without changing default configuration
values.

As such all SMTP connections require the following:

1.  TLS Connection (STARTTLS or SMTPS) has been negotiated before authentication or sending emails (unauthenticated
connections require it as well)
2.  Valid X509 Certificate presented to the client during the TLS handshake

There is an option to disable both of these security measures however they are **not recommended**.

The following configuration options exist to configure the security level in order of most preferable to least
preferable:

### Configuration Option: certificates_directory

You can [configure a directory](../configuration/miscellaneous.md#certificates_directory) of certificates for Authelia
to trust. These certificates can either be CA's or individual public certificates that should be trusted. These
are added in addition to the environments PKI trusted certificates if available. This is useful for trusting a
certificate that is self-signed without drastically reducing security. This is the most recommended workaround to not
having a valid PKI trusted certificate as it gives you complete control over which ones are trusted without disabling
critically needed validation of the identity of the target service.

Read more in the [documentation](../configuration/miscellaneous.md#certificates_directory) for this option.

### Configuration Option: tls.skip_verify

The [tls.skip_verify](../configuration/notifier/smtp.md#tls) option allows you to skip verifying the certificate
entirely which is why [certificates_directory](#configuration-option-certificates_directory) is preferred over this.
This will effectively mean you cannot be sure the certificate is valid which means an attacker via DNS poisoning or MITM
attacks could intercept emails from Authelia compromising a user's security without their knowledge.

### Configuration Option: disable_require_tls

Authelia by default ensures that the SMTP server connection is secured via STARTTLS or SMTPS prior to sending sensitive
information. The [disable_require_tls](../configuration/notifier/smtp.md#disable_require_tls) disables this requirement
which means the emails are sent in plain text. This is the least secure option as it effectively removes the validation
of SMTP certificates and removes the encryption offered by the STARTTLS/SMTPS connection all together.

This means not only can the vulnerabilities of the [skip_verify](#configuration-option-tlsskip_verify) option be
exploited, but any router or switch along the route of the email which receives the packets could be used to silently
exploit the plain text nature of the email. This is only usable currently with authentication disabled (comment out the
password) and as such is only an option for SMTP servers that allow unauthenticated relay (bad practice).

### SMTPS vs STARTTLS

All connections start as plain text and are upgraded via STARTTLS. SMTPS is an exception to this rule where the
connection is over TLS. As SMTPS is deprecated, the only way to configure this is to set the SMTP
[port](../configuration/notifier/smtp.md#port) to the officially recognized SMTPS port of 465 which will cause Authelia
to automatically consider it to be a SMTPS connection. As such your SMTP server, if not offering SMTPS, should not be
listening on port 465 which is bad practice anyway.

## Additional security

### Reset Password

It's possible to disable the reset password functionality and is an optional adjustment to consider for anyone wanting
to increase security. See the [configuration](../configuration/authentication/index.md#disable_reset_password) for more
information.

### Session security

We have a few options to configure the security of a session. The main and most important one is the session secret.
This is used to encrypt the session data when it is stored in the [Redis](../configuration/session/redis.md) key value
database. The value of this option should be long and as random as possible. See more in the
[documentation](../configuration/session/index.md#secret) for this option.

The validity period of session is highly configurable. For example in a highly security conscious domain you could
set the session [remember_me_duration](../configuration/session/index.md#remember_me_duration) to 0 to disable this
feature, and set the [expiration](../configuration/session/index.md#expiration) to 2 hours and the
[inactivity](../configuration/session/index.md#inactivity) of 10 minutes. Configuring the session security in this
manner would mean if the cookie age was more than 2 hours or if the user was inactive for more than 10 minutes the
session would be destroyed.

### Additional proxy protection measures

You can also apply the following headers to your proxy configuration for improving security. Please read the
relevant documentation for these headers before applying them blindly.

#### nginx

```text
# We don't want any credentials / TOTP secret key / QR code to be cached by
# the client
add_header Cache-Control "no-store";
add_header Pragma "no-cache";

# Clickjacking / XSS protection

# We don't want Authelia's login page to be rendered within a <frame>,
# <iframe> or <object> from an external website.
add_header X-Frame-Options "SAMEORIGIN";

# Block pages from loading when they detect reflected XSS attacks.
add_header X-XSS-Protection "1; mode=block";
```

#### Traefik 2.x - Kubernetes CRD

```yaml
---
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
  name: headers-authelia
spec:
  headers:
    browserXssFilter: true
    customFrameOptionsValue: "SAMEORIGIN"
    customResponseHeaders:
      Cache-Control: "no-store"
      Pragma: "no-cache"
---
apiVersion: traefik.containo.us/v1alpha1
kind: IngressRoute
metadata:
  name: authelia
spec:
  entryPoints:
    - http
  routes:
    - match: Host(`auth.example.com`) && PathPrefix(`/`)
      kind: Rule
      priority: 1
      middlewares:
        - name: headers-authelia
          namespace: authelia
      services:
        - name: authelia
          port: 80
```

#### Traefik 2.x - docker-compose

```yaml
services:
  authelia:
    labels:
      - "traefik.http.routers.authelia.middlewares=authelia-headers"
      - "traefik.http.middlewares.authelia-headers.headers.browserXssFilter=true"
      - "traefik.http.middlewares.authelia-headers.headers.customFrameOptionsValue=SAMEORIGIN"
      - "traefik.http.middlewares.authelia-headers.headers.customResponseHeaders.Cache-Control=no-store"
      - "traefik.http.middlewares.authelia-headers.headers.customResponseHeaders.Pragma=no-cache"
```

### More protections measures with fail2ban

If you are running fail2ban, adding a filter and jail for Authelia can reduce load on the application / web server.
Fail2ban will ban IPs exceeding a threshold of repeated failed logins at the firewall level of your host. If you are
using Docker, the Authelia log file location has to be mounted from the host system to the container for
fail2ban to access it.

Create a configuration file in the `filter.d` folder with the content below. In Debian-based systems the folder is
typically located at `/etc/fail2ban/filter.d`.

```text
# Fail2Ban filter for Authelia

# Make sure that the HTTP header "X-Forwarded-For" received by Authelia's backend
# only contains a single IP address (the one from the end-user), and not the proxy chain
# (it is misleading: usually, this is the purpose of this header).

# the failregex rule counts every failed login (wrong username or password) and failed TOTP entry as a failure
# the ignoreregex rule ignores debug, info and warning messages as all authentication failures are flagged as errors

[Definition]
failregex = ^.*Error while checking password for.*remote_ip=<HOST> stack.*
            ^.*Credentials are wrong for user .*remote_ip=<HOST> stack.*
            ^.*Wrong passcode during TOTP validation.*remote_ip=<HOST> stack.*

ignoreregex = ^.*level=debug.*
              ^.*level=info.*
              ^.*level=warning.*
```

Modify the `jail.local` file. In Debian-based systems the folder is typically located at `/etc/fail2ban/`. If the file
does not exist, create it by copying the jail.conf `cp /etc/fail2ban/jail.conf /etc/fail2ban/jail.local`. Add an
Authelia entry to the "Jails" section of the file:

```text
[authelia]
enabled = true
port = http,https,9091
filter = authelia
logpath = /path-to-your-authelia.log
maxretry = 3
bantime = 1d
findtime = 1d
chain = DOCKER-USER
```

If you are not using Docker remove the line "chain = DOCKER-USER". You will need to restart the fail2ban service for the
changes to take effect.

## Container privilege de-escalation

Authelia will run as the root user and group by default, there are two options available to run as a non-root user and
group.

It is recommended which ever approach you take that to secure the sensitive files Authelia requires access to that you
make sure the chmod of the files does not inadvertently allow read access to the files by users who do not need access
to them.

Examples:

If you wanted to run Authelia as UID 8000, and wanted the GID of 9000 to also have read access to the files
you might do the following assuming the files were in the relative path `.data/authelia`:

```shell
chown -r 8000:9000 .data/authelia
find .data/authelia/ -type d -exec chmod 750 {} \;
find .data/authelia/ -type f -exec chmod 640 {} \;
```

If you wanted to run Authelia as UID 8000, and wanted the GID of 9000 to also have write access to the files
you might do the following assuming the files were in the relative path `.data/authelia`:

```shell
chown -r 8000:9000 .data/authelia
find .data/authelia/ -type d -exec chmod 770 {} \;
find .data/authelia/ -type f -exec chmod 660 {} \;
```

### Docker user directive

The docker user directive allows you to configure the user the entrypoint runs as. This is generally the most secure
option for containers as no process accessible to the container ever runs as root which prevents a compromised container
from exploiting unnecessary privileges.

The directive can either be applied in your `docker run` command using the `--user` argument or by
the docker-compose `user:` key. The examples below assume you'd like to run the container as UID 8000 and GID 9000.

Example for the docker CLI:

```shell
docker run --user 8000:9000 -v /authelia:/config authelia/authelia:latest
```

Example for docker-compose:

```yaml
version: '3.8'
services:
  authelia:
    image: authelia/authelia
    container_name: authelia
    user: 8000:9000
    volumes:
      - ./authelia:/config
```

Running the container in this way requires that you manually adjust the file owner at the very least as described above.
If you do not do so it will likely cause Authelia to exit immediately. This option takes precedence over the PUID and
PGID environment variables below, so if you use it then changing the PUID and PGID have zero effect.

### PUID/PGID environment variables using the entrypoint

The second option is to use the `PUID` and `PGID` environment variables. When the container entrypoint is executed
as root, the entrypoint automatically runs the Authelia process as this user. An added benefit of using the environment
variables is the mounted volumes ownership will automatically be changed for you. It is still recommended that
you run the find chmod examples above in order to secure the files even further especially on servers multiple people
have access to.

The examples below assume you'd like to run the container as UID 8000 and GID 9000.

Example for the docker CLI:

```shell
docker run -e PUID=1000 -e PGID=1000 -v /authelia:/config authelia/authelia:latest
```

Example for docker-compose:

```yaml
version: '3.8'
services:
  authelia:
    image: authelia/authelia
    container_name: authelia
    environment:
      PUID: 1000
      PGID: 1000
    volumes:
      - ./authelia:/config
```

[HSTS]: https://www.nginx.com/blog/http-strict-transport-security-hsts-and-nginx/

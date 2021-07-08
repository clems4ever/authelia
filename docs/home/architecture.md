---
layout: default
title: Architecture
parent: Home
nav_order: 1
---

**Authelia** is a companion of reverse proxies like Nginx, Traefik and HAProxy. It can be seen as an extension of those
proxies providing authentication functions and a login portal.

As shown in the following architecture diagram, Authelia is directly connected to the reverse proxy but never directly
connected to application backends and therefore the payloads sent by the clients of the protected API will never reach
Authelia, only the authentication pieces will, like the Authorization header for instance. Consequently, the protected
APIs can be a REST, a GraphQL or any other kind of API on top of HTTP.

<p align="center" style="margin:50px">
  <img alt="Architecture Diagram" src="../images/archi.png"/>
</p>

## Workflow

Reverse proxies are configured so that every incoming request generates an authentication request sent to Authelia.
Authelia responds and will instruct the reverse proxy to either allow the incoming request to pass through, or block it
because the user is not authenticated or is not sufficiently authorized.

### Step by step

When the first request of an unauthenticated user hits the reverse proxy, Authelia determines the user is not
authenticated because no session cookie has been sent along with the request. Consequently, Authelia redirects the user
to the authentication portal provided by Authelia itself. The user can then execute the authentication workflow using
that portal to obtain a session cookie valid for all subdomains of the domain protected by Authelia.

When the user visits the initial website again, the query is sent along with the session cookie which is forwarded in
the authentication request to Authelia. This time, Authelia can verify the user is authenticated and order the reverse
proxy to let the query pass through.

### Sequence Diagram

Here is a description of the complete workflow:

<p align="center">
  <img alt="Sequence Diagram" src="../images/sequence-diagram.png"/>
</p>

## HTTP/HTTPS

Authelia only works for websites served over HTTPS because the session cookie can only be transmitted over secure
connections. Please note that it has been decided that we won't support websites served over HTTP in order to avoid any
risk due to misconfiguration. (see [#590](https://github.com/authelia/authelia/issues/590)).

If a self-signed certificate is required, the following command can be used to generate one for example.com in the
/tmp/certs directory (substitute your desired values for these):

```shell
docker run authelia/authelia authelia certificates generate --host example.com --dir /tmp/certs/
```

# Eso-sops-server

is a backend for External Secret Operator using the [Webhook Provider](https://external-secrets.io/latest/provider/webhook/).



## Overview

![Architecture](architecture.drawio.png)

Secrets Files in Json Format can be stored save and encrypted in Git. 
For the encyrption [age](https://github.com/FiloSottile/age) in common with [SOps](https://github.com/getsops/sops) is used.

The server will clone the repo. Decryption of the secretfile on access.
The whole file file will be afterwards provided to ESO.

## Request Interface
<pre>
http://hostname:8080/reponame/filename
</pre>

## TODOS / Roadmap

* Enable TLS
* Let ESO Authenticate via User
* Map a User it a age private key
* Support more then one Repository
* Support other Storage Backend ( s3 , provided PVC ) 
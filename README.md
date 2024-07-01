This is a custom implementation on top of Benthos to register a new function called `encrypt_rsa` which receives 2 parameters:

* key: the publis RSA key to be used for encryption
* text: the plain text to be encrypted

In order to run this project, you need to run the following commands:

* `go build -o encrypt`
* `KEY="<your public key value here>" ./encrypt -c config.yaml`

The public key should look like something like this: 

```
KEY="-----BEGIN RSA PUBLIC KEY-----\nMqMVpFa ... more here ... L9Vb3XwIDAQAB\n-----END RSA PUBLIC KEY-----\n"
```

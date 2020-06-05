### Create keys

it is created in the tokens folder

**private key**

`openssl genrsa -out private.rsa 1024`

**public key**

`openssl rsa -in private.rsa -pubout > public.rsa.pub`

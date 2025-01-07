generate-rsa-private-key-access:
	openssl genrsa -out access-private.pem 2048
generate-rsa-public-key-access:
	openssl rsa -in access-private.pem -outform PEM -pubout -out access-public.pem

generate-new-rsa-keys:
	make generate-rsa-private-key-access
	make generate-rsa-public-key-access
import hmac
import hashlib
import base64


canonical_string_from_browser = "/auth/login_eyJ1c2VybmFtZSI6ImRxc3NkcWRzcUBnbWFpbC5jb20iLCJwYXNzd29yZCI6IjNFUkZzZGZnZmQhIn0=_2025-10-28T00:11:00.000Z" 

secret_key = "wQxTBLXYMXMbVjc7wnXkcsr2phZCRVnJNAa4CRJfUCqsQhc6qJ"


secret_key_bytes = secret_key.encode('utf-8')

canonical_string_bytes = canonical_string_from_browser.encode('utf-8')
print(canonical_string_bytes)
hasher = hmac.new(secret_key_bytes, canonical_string_bytes, hashlib.sha256)

binary_signature = hasher.digest()

python_signature = base64.b64encode(binary_signature).decode('utf-8')


print(f"[+] Python Script Signature: {python_signature}")
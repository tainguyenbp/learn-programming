import base64

mystr_encoded = 'tainguyenbp@passw0rd'
mystr_decoded = 'dGFpbm5AMTIzNDU2Cg=='

# Encode
mystr_encoded = base64.b64encode(mystr_encoded.encode('utf-8'))
print(mystr_encoded)

# Decode
mystr_decoded = base64.b64decode(mystr_encoded).decode('utf-8')
print(mystr_decoded)

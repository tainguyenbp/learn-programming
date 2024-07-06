# You sould install pycryptodome before runing this code
# pip install pycryptodome
import base64
from Crypto.Cipher import DES
from Crypto.Util.Padding import pad, unpad

def encrypt_data(data, key):
    cipher = DES.new(key.encode(), DES.MODE_ECB)
    padded_data = pad(data.encode(), DES.block_size)
    encrypted_data = cipher.encrypt(padded_data)
    return base64.b64encode(encrypted_data).decode('utf-8')

def decrypt_data(encrypted_data, key):
    cipher = DES.new(key.encode(), DES.MODE_ECB)
    decrypted_data = cipher.decrypt(base64.b64decode(encrypted_data.encode()))
    return unpad(decrypted_data, DES.block_size).decode('utf-8')

if __name__ == '__main__':
    key = 'abcdefgh'                 # 8 bytes key for DES
    data = 'Hello, World'            # Data to be encrypted

    encrypted_data = encrypt_data(data, key)
    print('Encrypted data:', encrypted_data)

    decrypted_data = decrypt_data(encrypted_data, key)
    print('Decrypted data:', decrypted_data)

# Importing required libraries
import base64, os

# Importing required libraries from cryptography
from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes
from cryptography.hazmat.backends import default_backend

# Function to encrypt data using AES-GCM algorithm and return the encrypted data in string format
def encrypt(data:str, key:str) -> str:
    iv = os.urandom(12)
    encryptor = Cipher(algorithms.AES(key.encode('utf-8')), modes.GCM(iv), backend=default_backend()).encryptor()
    encrypted_data = encryptor.update(data.encode('utf-8')) + encryptor.finalize()
    return base64.urlsafe_b64encode(iv + encryptor.tag + encrypted_data).decode('utf-8')

# Function to decrypt data using AES-GCM algorithm and return the decrypted data in string format
def decrypt(encrypted_data, key) -> str:
    decoded_data = base64.urlsafe_b64decode(encrypted_data)
    iv = decoded_data[:12]
    tag = decoded_data[12:28]
    encrypted_data = decoded_data[28:]
    decryptor = Cipher(algorithms.AES(key.encode('utf-8')), modes.GCM(iv, tag), backend=default_backend()).decryptor()
    return (decryptor.update(encrypted_data) + decryptor.finalize()).decode('utf-8')

# Main function to test the above functions
if __name__ == '__main__':
    key = '689ef728d55342d9af07ed4194cf1d4C' # 32 bytes key for AES-256
    data = 'Hello, World'                    # Data to be encrypted

    # Encrypting and decrypting the data
    encrypted_data = encrypt(data, key)
    print('Encrypted data:', encrypted_data)

    decrypted_data = decrypt(encrypted_data, key)
    print('Decrypted data:', decrypted_data)

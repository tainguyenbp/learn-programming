# import os
# from werkzeug.security import check_password_hash, generate_password_hash
# from dotenv import load_dotenv

# load_dotenv()

# def login(username, password):
#     admin_username = os.getenv('ADMIN_USERNAME')
#     admin_password_hash = os.getenv('ADMIN_PASSWORD_HASH')

#     if username == admin_username and check_password_hash(admin_password_hash, password):
#         # Login successful
#         return True
#     else:
#         # Login failed
#         return False

# # Example usage
# if __name__ == "__main__":
#     # For demonstration, let's generate a hash for a password
#     password_to_hash = "password123"
#     hashed_password = generate_password_hash(password_to_hash)
#     print(f"Hashed password for '{password_to_hash}': {hashed_password}")

#     # Simulate a login attempt
#     username_input = "admin"
#     password_input = "password123"

#     if login(username_input, password_input):
#         print("Login successful!")
#     else:
#         print("Login failed!")

# import os
# from werkzeug.security import check_password_hash, generate_password_hash
# from dotenv import load_dotenv

# load_dotenv()

# def login(username, password):
#     admin_username = os.getenv('ADMIN_USERNAME')
#     admin_password_hash = os.getenv('ADMIN_PASSWORD_HASH')

#     if username == admin_username and check_password_hash(admin_password_hash, password):
#         # Login successful
#         return True
#     else:
#         # Login failed
#         return False

# # Example usage
# if __name__ == "__main__":
#     # For demonstration, let's generate a hash for a password
#     password_to_hash = "password123"
#     hashed_password = generate_password_hash(password_to_hash)
#     print("Hashed password for '{}': {}".format(password_to_hash, hashed_password))

#     # Simulate a login attempt
#     username_input = "admin"
#     password_input = "password123"

#     if login(username_input, password_input):
#         print("Login successful!")
#     else:
#         print("Login failed!")

# import os
# import argparse
# from werkzeug.security import check_password_hash, generate_password_hash
# from dotenv import load_dotenv

# def login(username, password):
#     admin_username = os.getenv('ADMIN_USERNAME')
#     admin_password_hash = os.getenv('ADMIN_PASSWORD_HASH')

#     if username == admin_username and check_password_hash(admin_password_hash, password):
#         # Login successful
#         return True
#     else:
#         # Login failed
#         return False

# def main():
#     # Set up argument parsing
#     parser = argparse.ArgumentParser(description="Run the login script with environment variables.")
#     parser.add_argument('-e', '--env', action='append', help="Environment variables in the form KEY=VALUE")
    
#     args = parser.parse_args()

#     # Set environment variables from command line arguments
#     if args.env:
#         for env_var in args.env:
#             key, value = env_var.split('=', 1)
#             os.environ[key] = value

#     # Load .env file if it exists
#     load_dotenv()

#     # Simulate a login attempt
#     username_input = "admin"
#     password_input = "password123"

#     if login(username_input, password_input):
#         print("Login successful!")
#     else:
#         print("Login failed!")

# if __name__ == "__main__":
#     main()

import bcrypt

# Example function to hash a password (this would be done during user registration)
def hash_password(password):
    salt = bcrypt.gensalt()
    hashed = bcrypt.hashpw(password.encode(), salt)
    return hashed

# Example hashed password storage (in a real application, this would be stored in a database)
stored_password_hash = hash_password('password123')

# Function to verify login credentials
def login(username, password):
    # In a real application, retrieve the stored password hash from the database
    if username == 'admin':
        return bcrypt.checkpw(password.encode(), stored_password_hash)
    else:
        # Login failed
        return False

# Example usage
print(login('admin', 'password123'))  # True
print(login('admin', 'wrongpassword'))  # False
print(login('user', 'password123'))  # False

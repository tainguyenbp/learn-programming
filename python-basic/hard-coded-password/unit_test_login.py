import unittest
import bcrypt
from your_module import login, hash_password  # Replace 'your_module' with the actual module name where login and hash_password functions are defined

class TestLogin(unittest.TestCase):
    
    @classmethod
    def setUpClass(cls):
        # Setup any necessary preparations, such as hashing a password for testing
        cls.stored_password_hash = hash_password('password123')

    def test_correct_login(self):
        self.assertTrue(login('admin', 'password123'))

    def test_incorrect_password(self):
        self.assertFalse(login('admin', 'wrongpassword'))

    def test_incorrect_username(self):
        self.assertFalse(login('user', 'password123'))

    # Add more test cases as needed

if __name__ == '__main__':
    unittest.main()

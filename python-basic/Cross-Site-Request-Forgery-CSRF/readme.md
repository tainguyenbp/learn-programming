```
python3 -m venv path/to/venv
source path/to/venv/bin/activate
python3 -m pip install requests
pip install Flask Flask-WTF Flask-Session

# Step 1: Fetch the CSRF token and store the session cookie
csrf_token=$(curl -s -c cookiefile http://127.0.0.1:5000/get_csrf_token | jq -r '.csrf_token')

# Step 2: Use the CSRF token and the session cookie in the transfer request
curl -X POST -b cookiefile http://127.0.0.1:5000/dashboard -d "amount=100&destination_account=123456&csrf_token=${csrf_token}"
```

from flask import Flask, request, jsonify

import requests

app = Flask(__name__)

# Fetch data from JSONPlaceholder

API_URL = "https://jsonplaceholder.typicode.com/posts"

@app.route('/posts', methods=['GET'])

def get_posts():
    user_id = request.args.get('userId')
    response = requests.get(API_URL)
    posts = response.json()

    # Filter posts by userId if provided
    if user_id:

        posts = [post for post in posts if post['userId'] == int(user_id)]

    return jsonify(posts)

if __name__ == '__main__':

    app.run(host='0.0.0.0', port=5000)

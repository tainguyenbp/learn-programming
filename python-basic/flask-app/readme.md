# Step 1: Application

```

Open code Environment
..if doesn't work click again in 30sec
First, let's create a simple Flask application [+] that retrieves data from an open-source API such as JSONPlaceholder and allows filtering.

Open your code editor by clicking Open Code Environment above.

Next open Terminal in the Code Editor and Create a new directory for your project. To open Terminal press CTRL+` On Windows or CMD+J on MacOS.

Alternatively you can open terminal from menubar:

image

Once terminal is open type this command and run it. It will create new folder flask-app

mkdir flask-app
Next export directory as an environment variable [+]

export LABROOT=/home/labs/flask-app
Lets navigate into newly created directory by typing cd (change directory) command in the terminal.

cd flask-app
Install python virtual environment venv [+], create it and then a Python virtual environment and activate it:

apt install python3.10-venv -y
python3 -m venv venv
source venv/bin/activate
Install Flask using pip:

pip install Flask requests
Create a file named app.py with the content below. To create file you can type code app.py in the terminal and then save it with Ctrl+S (Win) or CMD+S (MacOS). Alternatively you can create file from the file explorer on the left side of the window.

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
Run the Flask app to test it locally:

python app.py
Now Open second terminal, you can use Image above for second terminal or use split terminal screen for that (Ctrl+\ on Windows or CMD+\ On MacOS) and type command below. This will execute curl [+] command to the Flask application:

curl http://127.0.0.1:5000/posts
You can add a query parameter to filter by user ID, e.g.

curl http://127.0.0.1:5000/posts?userId=1
Curl returns HTTP response [200] - Status OK then all is correct and you can finish this step. Once step is passed you can stop running python app.py by typing Ctrl+C
```

# Step 2: Docker
```
In this stage, we'll containerize the Flask application [+] we created in Stage 1 using Docker. Containerizing the application allows it to run consistently across different environments, making it easier to deploy.

Step 1: Create a Dockerfile
First, let's create a Dockerfile that defines the environment and steps required to run our Flask application inside a Docker container.

Open your Terminal in the Code Editor and navigate to the flask-app directory if you're not already there.

cd flask-app
Create a file named Dockerfile with the following content:

# Use an official Python runtime as a parent image
FROM python:3.9-slim

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any needed packages specified in requirements.txt
RUN pip install --no-cache-dir Flask requests

# Make port 3000 available to the world outside this container
EXPOSE 3000

# Define environment variable
ENV FLASK_APP=app.py

# Run the application
CMD ["flask", "run", "--host=0.0.0.0", "--port=3000"]
Step 2: Build the Docker Image
Now that the Dockerfile is ready, we can build the Docker image. Run the following command in your terminal:

docker build -t flask-app .
This command tells Docker to build an image from the Dockerfile in the current directory and tag it as flask-app.

Step 3: Run the Docker Container
Once the image is built, you can run it in a container using the following command:

docker run -p 3000:3000 flask-app
This command maps port 3000 on your host machine to port 3000 in the container, making the Flask app accessible at http://localhost:3000.

Step 4: Test the Application
To test that the Flask app is running inside the Docker container [+], open a new terminal or use split terminal (Ctrl+\ on Windows or CMD+\ On MacOS) and run:

curl http://127.0.0.1:3000/posts
You can also filter by user ID, e.g., http://127.0.0.1:3000/posts?userId=1.

If everything is set up correctly, you should see the output from your Flask application served from within the Docker container via Curl. Now you can complete the step and once completed you can stop docker container by typing Ctrl+C

```



# Step 3: Container Registry

```
In this stage, we'll upload the Docker [+] image we created in Stage 2 to the GitHub Container Registry [+] (GHCR). This allows you to store and share your Docker images via GitHub.

Step 1: Authenticate with GitHub Container Registry
First, you need to authenticate Docker with the GitHub Container Registry. You can do this using a GitHub Personal Access Token (PAT).

Create a GitHub Personal Access Token (PAT) with the write:packages, read:packages, and delete:packages permissions. You can do this from your GitHub account settings under Developer settings > Personal access tokens. Use this link to navigate directly to the Personal access tokens page https://github.com/settings/tokens. After navigating to the page click New personal access token (classic) and create your token with persmissions indicated above.

Export your GitHub username and access token as an environment variable using export= [+]

export GH_TOKEN=<github_access_token>
export GH_USERNAME=<github_username>
Replace github_username with your actual GitHub username.

Authenticate Docker with GitHub using your PAT:
echo $GH_TOKEN | docker login ghcr.io -u $GH_USERNAME --password-stdin
Export Image Name as an Environment Variable
export GH_IMAGE="ghcr.io/${GH_USERNAME,,}/flask-app:latest"
Step 2: Tag the Docker Image
Before pushing the image to GHCR, you need to tag it with the correct repository path. The format for the tag is ghcr.io/$GH_USERNAME/<repository-name>:<tag>.

Run the following command to tag your Docker image [+]

docker tag flask-app $GH_IMAGE
Step 3: Push the Docker Image to GitHub Container Registry
Now, push the tagged image to the GitHub Container Registry:

docker push $GH_IMAGE
This command uploads the Docker image to your GitHub account under the flask-app repository in the GitHub Container Registry.

Step 4: Verify the Image Upload
To verify that the image has been successfully uploaded, you can do one of the following:

Check via GitHub Website:

Go to https://github.com/your-username?tab=packages.
You should see the flask-app package listed under the Packages section.
Replace your-username with your GitHub username.

Pull the Image Locally:

Run the following command to pull the image from GHCR:
docker pull $GH_IMAGE
Once pulled, you can verify it by listing your local Docker images:
docker images
You should see ghcr.io/your-username/flask-app listed among the images.
If the image is listed, the upload to the GitHub Container Registry was successful.

```

# 
Step 4: Helm Charts
```
Open code Environment
..if doesn't work click again in 30sec
In this stage, we’ll create Helm charts [+] to deploy our Flask application on Kubernetes. Helm allows us to package, configure, and deploy Kubernetes applications more easily and consistently. We’ll create a deployment with different kubernetes deployment replica counts [+] for production and development environments and expose the application using a ClusterIP service.

First, we need to create an empty Helm chart. This chart will contain all the necessary templates and configurations for our application.

Ensure you are in the flask-app directory.
Run the following command to create a new Helm chart:
helm create flask-app-chart
This command creates a directory named flask-app-chart with a predefined structure for your Helm chart.

The helm create command generates some default files that we don’t need. Let’s clean up the chart to keep it minimal and focused on our application.

Delete the unnecessary files:
rm -rf $LABROOT/flask-app-chart/templates/tests
rm $LABROOT/flask-app-chart/templates/NOTES.txt
rm $LABROOT/flask-app-chart/templates/serviceaccount.yaml
rm $LABROOT/flask-app-chart/templates/ingress.yaml
rm $LABROOT/flask-app-chart/templates/hpa.yaml
Open the flask-app-chart/Chart.yaml file and update it with the following content:
apiVersion: v2
name: flask-app
description: A Helm chart for deploying the Flask app
version: 0.1.0
appVersion: "1.0"
We’ll now create a Kubernetes deployment template to run our Flask application with different replica counts for production and development.

Replace the content of flask-app-chart/templates/deployment.yaml with the following:
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "flask-app-chart.fullname" . }}
  labels:
    app: {{ include "flask-app-chart.name" . }}
spec:
  replicas: {{ .Values.replicaCount }}
  selector:
    matchLabels:
      app: {{ include "flask-app-chart.name" . }}
  template:
    metadata:
      labels:
        app: {{ include "flask-app-chart.name" . }}
    spec:
      imagePullSecrets:
      - name: {{ .Values.imagePullSecretName }}
      containers:
      - name: flask-app
        image: {{ .Values.image.repository }}:{{ .Values.image.tag }}
        ports:
        - containerPort: 3000
        resources:
          requests:
            memory: "64Mi"
            cpu: "50m"
          limits:
            memory: "128Mi"
            cpu: "200m"
Next, we’ll expose our application using a Kubernetes ClusterIP service. [+]

Replace the content of flask-app-chart/templates/service.yaml with the following:
apiVersion: v1
kind: Service
metadata:
  name: {{ include "flask-app-chart.fullname" . }}
  labels:
    app: {{ include "flask-app-chart.name" . }}
spec:
  type: ClusterIP
  ports:
    - port: 80
      targetPort: 3000
  selector:
    app: {{ include "flask-app-chart.name" . }}
This service will make the Flask application accessible within the Kubernetes cluster on port 80 and target traffic from port 80 to the port 3000 which was exposed port for our Docker container.

We’ll define different values for the production and development environments using Helm’s values files.

Open the flask-app-chart/values.yaml file and update it with the following common values:
Pay attention! - replace ghcr.io/your-username/flask-app with correct github image name in the code below.

replicaCount: 1
imagePullSecretName: ghcr-secret

image:
  repository: ghcr.io/your-username/flask-app
  tag: latest
Create a file named values-prod.yaml in flask-app-chart directory for production-specific values:
replicaCount: 3
Create a file named values-dev.yaml in flask-app-chart directory for development-specific values:
replicaCount: 1
Next we need to create a namespace in Kubernetes [+] where we’ll deploy the Flask application. This helps isolate resources and makes management easier.

kubectl create namespace flask-app-namespace
Now we need to create Kubernetes Docker pull secret [+]

kubectl create secret docker-registry ghcr-secret --docker-server=ghcr.io --docker-username=$GH_USERNAME --docker-password=$GH_TOKEN -n flask-app-namespace
For purposes of this lab we will use the development environment, make sure you are in flask-app directory and run below:

helm upgrade --install flask-app-dev $LABROOT/flask-app-chart/ -f $LABROOT/flask-app-chart/values-dev.yaml --namespace flask-app-namespace
Optionally to deploy the Helm chart for the production environment we can use command below

helm upgrade --install flask-app-prod $LABROOT/flask-app-chart/ -f $LABROOT/flask-app-chart/values-prod.yaml --namespace flask-app-namespace
Regardless of environment these commands will deploy your Flask application to Kubernetes with the appropriate replica counts and expose it via a ClusterIP service.

You can check the status of all Kubernetes pods [+] in given namespace by running:

kubectl get pods -n flask-app-namespace
The result should show something similar (Pod name ending will be different).

| Pod Name                       | READY | STATUS            | RESTARTS | AGE |
|--------------------------------|-------|-------------------|----------|-----|
| flask-app-prod-d76b9f6b5-p9s9j |  1/1  | Running           |    0     | 15s |
```

# Step 5: Manual Deployment
```
Open code Environment
..if doesn't work click again in 30sec
In this stage, you will access a service running in a Kubernetes cluster from your local machine using port forwarding. Port forwarding is a powerful feature that allows you to establish a connection between your local machine and a Kubernetes service without exposing the service externally via a LoadBalancer or Ingress. This method is particularly useful for debugging and testing applications that are running within a Kubernetes cluster.

Step 1: Port Forward the Service to Access the Application
To access the application externally, use kubectl port-forward. This command allows you to map a port on your local machine to a port on a service running in the Kubernetes cluster. This means that you can interact with the service as if it were running locally, even though it's actually running inside the cluster.

Simple Port Forward
This method involves forwarding a port from your local machine to a port on a service in the Kubernetes cluster. This allows you to access the service using localhost.

Find the name of the service you created:
kubectl get svc -n flask-app-namespace
To stop port-forward just type CTRL + c in terminal

Port forward the service to your local machine:
kubectl port-forward svc/flask-app-prod 8080:80 -n flask-app-namespace
To access the service open a web browser and go to http://localhost:8080 or use curl in your terminal to interact with the service.

Port Forward with --address 0.0.0.0
This method extends the simple port forward by using the --address flag to bind the port to all network interfaces (0.0.0.0). This allows the service to be accessed not just from localhost, but from any device through your public IP address or DNS name.

Find the name of the service you created:
kubectl get svc -n flask-app-namespace
Port forward the service to your local machine:
kubectl port-forward svc/flask-app-prod 3000:80 -n flask-app-namespace --address 0.0.0.0
This command forwards port 3000 on your public IP to port 80 on the service within Kubernetes, allowing you to access the application using your machine's public IP or a connected DNS name.

With the port forwarding in place, you can now access your Flask application from a browser or via curl:

curl http://<your_preparesh_username>.environment.sh:3000/posts
Or, if you have a DNS name connected to your public IP:

curl http://your-dns-name:3000/posts
You can also test with a specific user ID:

curl http://<your_preparesh_username>.environment.sh:3000/posts?userId=1
If everything is set up correctly, you should see the output from your Flask application.

Step 3: Cleanup
Once you're done testing, you can clean up by deleting the Helm releases and the namespace:

Delete the production deployment:
helm uninstall flask-app-prod --namespace flask-app-namespace
Delete the development deployment:
helm uninstall flask-app-dev --namespace flask-app-namespace
Delete the namespace:
kubectl delete namespace flask-app-namespace
This will remove all the resources associated with your Flask application in Kubernetes.
```


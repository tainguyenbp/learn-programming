### setup lib elasticsearch
```
pip install elasticsearch
```
```
from elasticsearch import Elasticsearch
import json

# Configure Elasticsearch client
es = Elasticsearch(
    ["https://localhost:9200"],
    http_auth=("username", "password"),
    scheme="https",
    port=9200,
)

# Fetch all roles
roles = es.security.get_role()

# Save roles to a JSON file
with open("roles_export.json", "w") as f:
    json.dump(roles, f, indent=4)

print("Roles have been exported to roles_export.json")
```

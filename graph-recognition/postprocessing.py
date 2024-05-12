import json
import requests

def postprocess(vertices_list: list):
    json_data = []
    for vertex_id, vertex in enumerate(vertices_list):
        for adj_vertex_id in vertex.adjacency_list:
            json_data.append({
                "Source": vertex_id,
                "Destination": adj_vertex_id,
                "Weight": 1 
            })

    json_str = json.dumps(json_data, indent=2)
    url = 'http://localhost:8080'  
    headers = {'Content-Type': 'application/json'}
    response = requests.post(url, headers=headers, json=json_str)

    if response.status_code == 200:
        print("Data posted successfully")
    else:
        print("Failed to post data. Status code:", response.status_code)


import requests


response = requests.get("https://api.github.com/repos/kubernetes/kubernetes/pulls")

complete_detail = response.json()

# print(complete_detail[0]["id"])
# print(complete_detail[0]["user"]["login"])

for i in range(len(complete_detail)):
    print(complete_detail[i]["user"]["login"])
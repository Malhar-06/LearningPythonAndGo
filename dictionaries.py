student_info = {
    "name": "John",
    "age": 17,
    "courses": ["Math", "CompSci"],
    "class": "12th"
}

print(student_info["name"])
print(student_info["courses"][0])

#--------------------------------------------------------------------------->

ec2_instance_info = [
    {
        "instance_id": "i-1234567890abcdef0",
        "instance_type": "t2.micro",
        "instance_state": "running",
        "instance_launch_time": "2020-02-02T12:00:00Z",
        "instance_private_ip": "172.01.54.00"
    },
    {
        "instance_id": "i-1234567890abcdef1",
        "instance_type": "t2.micro",
        "instance_state": "stopped",
        "instance_launch_time": "2020-02-02T12:00:00Z",
        "instance_private_ip": "172.01.54.01"
    },
    {
        "instance_id": "i-1234567890abcdef2",
        "instance_type": "t2.micro",
        "instance_state": "terminated",
        "instance_launch_time": "2020-02-02T12:00:00Z",
        "instance_private_ip": "172.01.54.02"
    },
    {
        "instance_id": "i-1234567890abcdef3",
        "instance_type": "t2.micro",
        "instance_state": "running",
        "instance_launch_time": "2020-02-02T12:00:00Z",
        "instance_private_ip": "172.01.54.04"
    }
]


print(ec2_instance_info[0]["instance_id"])


#--------------------------------------------------------------------------->


ec2_instance_infor = {
    "instance1_ids": ["i-1234567890abcdef0", "i-1234567890abcdef1", "i-1234567890abcdef2", "i-1234567890abcdef3"],
    "instance1_types": ["t2.micro", "t2.micro", "t2.micro", "t2.micro"],
    "instance1_states": ["running", "stopped", "terminated", "running"],
    "instance1_launch_times": ["2020-02-02T12:00:00Z", "2020-02-02T12:00:00Z", "2020-02-02T12:00:00Z", "2020-02-02T12:00:00Z"],
    "instance1_private_ips": ["172.01.54.00", "172.01.54.01", "172.01.54.02", "172.01.54.04"]
 
}

print(ec2_instance_infor(instance1_ids[0])) 
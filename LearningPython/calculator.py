import sys
import os

def add(x, y):
    return(x + y)

def sub(x, y):
    return(x - y)

def mult(x, y):
    return(x * y)

x = float(sys.argv[1])     #By default arguments are read as strings
operation = sys.argv[2]
y = float(sys.argv[3])

if operation == "add":
    print(add(x, y))

elif operation == "sub":
    print(sub(x, y))

elif operation == "mult":
    print(mult(x, y))

else:
    print("I don't know that one.")

# Revision of environmental variables
    
print(os.getenv("password"))
print(os.getenv("username"))
# Exceptional Handling


import os

folders = input("Enter the list of folders names with spaces in between: ").split()


for folder in folders:

    try:
        files = os.listdir(folder)

    except FileNotFoundError:
        print("Please provide a valid folder name, looks like you have entered a wrong folder name." + folder)
        continue

    except PermissionError:
        print("You don't have permission to access this folder." + folder)
        break
        
    print(" ==== listing files for the folder :- " + folder)
    
    for file in files:
        print(file)
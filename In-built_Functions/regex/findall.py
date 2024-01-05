import re

text = "The quick brown fox"
pattern = r"brown"

search = re.search(pattern, text)
if search:
    print("Pattern found:", search.group())
else:
    print("Pattern not found")


#----------------------------------------------------------------------->
    

import re

words = ["apple", "banana", "cherry", "date"]
pattern = r"an"

for word in words:
    search = re.search(pattern, word)
    if search:
        print(f"Pattern found in {word}: {search.group()}")
    else:
        print(f"Pattern not found in {word}")



#----------------------------------------------------------------------->
        
#Searching for a pattern in a file

import re

with open("lorem-ipsum.txt", "r") as file:
    text = file.read()

pattern = r"brown"

search = re.search(pattern, text)
if search:
    print("Pattern found:", search.group())
else:
    print("Pattern not found")


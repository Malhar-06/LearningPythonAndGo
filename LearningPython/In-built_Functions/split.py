text = "Python is awesome"
words = text.split()
print("Words:", words)



text = "apple,banana,orange"
fruits = text.split(",")
print(fruits)  # Output: ['apple', 'banana', 'orange']



text = "apple,banana,orange,grape"
fruits = text.split(",", 2)
print(fruits)  # Output: ['apple', 'banana', 'orange,grape']

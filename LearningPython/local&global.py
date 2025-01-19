# Local variables.
# Variables that are created inside a function are known as local variables.
# Local variables can only be used inside the function where they are created.

def my_function():
    x = 10
    print("Inside the function, x =", x)

my_function()

# This will raise a NameError because x is not defined outside the function
print("Outside the function, x =", x)






# Global variables.
# Variables that are created outside of a function are known as global variables.
# Global variables can be used by everyone, both inside of functions and outside.

x = 10

def my_function():
    print("Inside the function, x =", x)

my_function()

print("Outside the function, x =", x)

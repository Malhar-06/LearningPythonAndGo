
# def greet(name: int):
#     print("Hello, " + str(name) + ". How are you?")

# greet(8)

###################################################################################################
###################################################################################################

# def add_numbers(x: int, y: int) -> int:
#     return(x + y)

# answer = add_numbers(8, 9)
# answer2 = add_numbers(8, int(9))
# print(answer)
# print(answer2)

# print(type(answer))

# print(isinstance(answer, int))
# print(isinstance(answer2, str))

###################################################################################################
###################################################################################################

def add_numbers(x, y):
    return(x + y)

answer = add_numbers(8, 9)
answer2 = add_numbers("mskdja", "9")
print(answer)
print(answer2)

print(type(answer))

print(isinstance(answer, int))
print(isinstance(answer2, str))
# def pyfunction(n):
#     for i in range(n):
#         for j in range(i+1):
#             print("*", end="")
#         print()
    
# pyfunction(5)


#<=========================================================================>#

# # Printing a square
# def pyfunc(r):
#     for x in range(r):
#         for y in range(r):
#             print('*', end='  ')
#         print()
# pyfunc(5)


# #printing a right angle triangle/ increasing triangle pattern
# def pyfunc(r):
#     for x in range(r):
#         for y in range(x+1):
#             print('*', end='  ')
#         print()
# pyfunc(5)

# #printing a decreasing triangle pattern
# def pyfunc(r):
#     for x in range(r):
#         for y in range(x, r):          #How? # first row/outer loop will run for 1
#             print('*', end='  ')
#         print()
# pyfunc(4)

#printing a right triangle pattern
def pyfunc(r):
    for x in range(r):                  #This will be same for every pattern
        for y in range(x, r):           #printing spaces 
            print(' ', end='  ')
        for y in range(x+1):            #printing stars
            print('*', end='  ')
        print()                         #This will be same for every pattern
pyfunc(0)

#printing a right triangle pattern
def pyfunc(r):
    for x in range(r):                  #This will be same for every pattern
        for y in range(x+1):           #printing spaces 
            print(' ', end='  ')    
        for y in range(x, r):            #printing stars
            print('*', end='  ')
        print()                         #This will be same for every pattern
pyfunc(0)

# Printing a Hill Pattern
def pyfunc(r):
    for x in range(r):
        for y in range(x, r):
            print(' ', end=' ')
        for y in range(x):
            print('*', end=' ')
        for y in range(x+1):
            print('*', end=' ')
        print()

pyfunc(0)


# pyramid star pattern
def pyfunc(n):
    for i in range(1, n+1):
        print(' ' * (n - i) + '*' * (2 * i - 1))
pyfunc(0)

# Reverse pyramid star pattern
def pyfunc(n):
    for i in range(n, 0, -1):
        print(' ' * (n - i) + '*' * (2 * i - 1))

pyfunc(0)

# Reverse pyramid star pattern
def pyfunc(n):
    for i in range(n):
        for j in range(i+1):
            print(' ', end=' ')
        for j in range(i, n):
            print('*', end=' ')
        for j in range(i+1, n):
            print('*', end=' ')
        print()

pyfunc(0)

class Diamond_Pattern:
    def pyfunc(self, r):
        for x in range(r):
            for y in range(x-1, r):
                print(' ', end=' ')
            for y in range(x-1):
                print('*', end=' ')
            for y in range((x+1)-1):
                print('*', end=' ')
            print()

    def inverted_diamond_pattern(self, n):
        for i in range(n):
            for j in range(i+1):
                print(' ', end=' ')
            for j in range(i, n):
                print('*', end=' ')
            for j in range(i+1, n):
                print('*', end=' ')
            print()

dp = Diamond_Pattern()
dp.pyfunc(5)
dp.inverted_diamond_pattern(5)

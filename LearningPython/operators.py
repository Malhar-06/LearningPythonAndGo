def arithmetic_operators(a, b):
    print(a + b)
    print(a - b)
    print(a * b)
    print(a / b)
    print(a % b)  # Modulus
    print(a ** b) # Power of a to b
    print(a // b) # Floor division

def comparison_operators(a, b): # relational operators
    print(a == b) # Equal to
    print(a != b) # Not equal to
    print(a > b)  # Greater than
    print(a < b)  # Less than
    print(a >= b) # Greater than or equal to
    print(a <= b) # Less than or equal to

def logical_operators(a, b):
    print(a and b) # And
    print(a or b)  # Or
    print(not a)   # Not

def assignment_operators(total, a, b):
    print(total)

    total += a # total = total + a
    print(total)
    total -= b # total = total - b
    print(total)
    total *= a # total = total * a
    print(total)
    total /= b # total = total / b
    print(total)


def identity_operators():
    mylist = [1, 2, 3]
    mylist2 = [1, 3, 2]
    print(mylist is mylist2)
    print(mylist is not mylist2)
    print(mylist in mylist2)
    print(mylist not in mylist2)
    
    print(1 in mylist)
    print(1 not in mylist)


identity_operators()
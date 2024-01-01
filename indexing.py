# indexing = accessign elements of a sequence using "[]" which is indexing operator
#                        [start:end:step]
#         starting index is inclusive and ending index is exclusive

full_name = "Malhar Bhaiyyasaheb Chikhale"
phone_number = "1234-5678-9012-2356"

#print(full_name[0:len(full_name):2])
print(full_name[-22])
#print(len(full_name))

# # print(phone_number[0:len(phone_number):2])

#   0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27             
# -28-27-26-25-24-23-22-21-20-19-18-17-16-15-14-13-12-11-10 -9 -8 -7 -6 -5 -4 -3 -2 -1
#   M  a  l  h  a  r     B  h  a  i  y  y  a  s  a  h  e  b     C  h  i  k  h  a  l  e


# def print_minus(phone_number):
#     for x in phone_number:
#         if x == '-':
#             print(x, end=' ')
        
# print_minus(1234-5678-9012-2356)



credit_card_number = "1234-5678-9012-2356"
print(f"XXXX-XXXX-XXXX-{credit_card_number[-4:]}")


# rEVERSE STRING
credit_card_number = "1234-5678-9012-2356"
print(credit_card_number[::-1])


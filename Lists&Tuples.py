#Lists

list_s3_bucket = ['bucket1', 'bucket2', 'bucket3']
print(list_s3_bucket[0])

appended = list_s3_bucket.append('bucket4')
print(appended)

sorted = list_s3_bucket.sort(reverse=False)
print(sorted)

inserted = list_s3_bucket.insert(0, 'bucket0')
print(inserted)

new = list_s3_bucket + ['bucket5', 'bucket6']
print(new)

sliced = list_s3_bucket[0:2]   # slicing
print(sliced)

removed = list_s3_bucket.remove('bucket0')
print(removed)

pop = list_s3_bucket.pop()
print(pop)

list_s3_bucket.pop(0)
print(list_s3_bucket)

list_s3_bucket.pop(1)
print(list_s3_bucket)






#Tuples
print("Tuples are immutable.")
print(             "                                                            ")
print(             "                                                            ")
print(             "                                                            ")

tuple_s3_bucket = ('bucket1', 'bucket2', 'bucket3')
#print(tuple_s3_bucket[0])

for i in tuple_s3_bucket:
    print(i)

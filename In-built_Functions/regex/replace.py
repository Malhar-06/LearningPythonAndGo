# sub() is used for substituting a specific pattern in the string with a replacement string.

import re

text = "The quick brown fox jumps over the lazy brown dog"
pattern = r"brown"

replacement = "red"

new_text = re.sub(pattern, replacement, text)
print("Modified text:", new_text)



#----------------------------------------------------------------------------------------------------->

import re

text = "apple,banana,orange,grape"
pattern = r","

split_result = re.split(pattern, text)
print("Split result:", split_result)
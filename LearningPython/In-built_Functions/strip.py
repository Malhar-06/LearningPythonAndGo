text = "   Some spaces around   "
stripped_text = text.strip()
print("Stripped text:", stripped_text)


text = "  Hello, world!  "
stripped_text = text.strip("!")
print(stripped_text)  # Output: "  Hello, world  "


text = "  Hello, world!  "
stripped_text = text.strip(" H!")
print(stripped_text)  # Output: "ello, world"

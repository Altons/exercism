def is_isogram(string):
    # Alternatively: Counter(string.lower()).most_common(1)[0][1] == 1
    repeated = {}
    for ch in string.lower():
        if ch == "-" or ch == " ":
            continue
        if ch in repeated:
            return False
        repeated[ch] = True
    return True

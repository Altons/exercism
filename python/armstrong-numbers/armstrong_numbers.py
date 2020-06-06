def is_armstrong_number(number):
    s = sum(int(char) ** len(str(number)) for char in str(number))
    return s == number

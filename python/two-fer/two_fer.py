def two_fer(name=None):
    if name == None:
        return "One for you, one for me."
    return "One for {}, one for me.".format(name)

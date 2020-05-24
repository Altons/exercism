
def convert(number):
    logic = {3: 'Pling', 5: 'Plang', 7: 'Plong'}
    r = ''
    for k in logic:
        if number % k == 0:
            r += logic[k]
    return str(number) if r == '' else r

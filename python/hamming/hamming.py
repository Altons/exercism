def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("Error: Strands are not equal in size")
    return sum([1 for x in zip(strand_a, strand_b) if x[0] != x[1]])

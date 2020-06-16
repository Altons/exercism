def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("Error: Strands are not equal in size")
    return sum(list(map(lambda x: x[0] != x[1], zip([c for c in strand_a], [c for c in strand_b]))))

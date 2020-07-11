class Hamming
  def self.compute(strand1, strand2)
    new(strand1, strand2).distance
  end

  def initialize(strand1, strand2)
    return raise ArgumentError, 'Strand lengths must be of identical length' unless strand1.length == strand2.length

    @nucleotides = strand1.chars.zip(strand2.chars)
  end

  def distance
    @nucleotides.count { |c| c.first != c.last }
  end
end

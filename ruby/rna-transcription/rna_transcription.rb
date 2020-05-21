# frozen_string_literal: true

class Complement
  def self.of_dna(seq)
    seq.gsub(/[GCTA]/, 'G' => 'C', 'C' => 'G', 'T' => 'A', 'A' => 'U')
  end
end

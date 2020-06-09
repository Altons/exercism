# frozen_string_literal: true

class Grains
  def self.square(number)
    raise ArgumentError unless (1..64).to_a.include? number

    2**(number - 1)
  end

  def self.total
    (1..64).inject { |total, n| total + 2**(n - 1) }
  end
end

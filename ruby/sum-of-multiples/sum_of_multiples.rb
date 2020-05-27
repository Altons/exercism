# frozen_string_literal: true

class SumOfMultiples
  def initialize(*multiples)
    @multiples = multiples
  end

  def to(input)
    (1...input).map { |i| i if @multiples.map { |v| (i % v).zero? }.any? }.compact.sum
  end
end

# sum_of_multiples = SumOfMultiples.new(3, 5)
# p sum_of_multiples.to(4)

# frozen_string_literal: true

class Series
  def initialize(digits)
    @digits = digits
  end

  def slices(n)
    return raise ArgumentError if @digits.size < n

    @digits.size.times.each_with_object([]) { |i, a| a << @digits[i, n] if i + n <= @digits.size }
  end
end

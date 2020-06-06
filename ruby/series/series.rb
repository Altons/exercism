# frozen_string_literal: true

class Series
  attr_reader :digits
  def initialize(digits)
    @digits = digits
  end

  def slices(n)
    return raise ArgumentError if @digits.size < n

    a = []
    digits.chars.each_cons(n) { |e| a << e.join }
    a
  end
end

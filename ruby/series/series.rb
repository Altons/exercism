# frozen_string_literal: true

class Series
  def initialize(digits)
    @digits = digits
  end

  def slices(n)
    return raise ArgumentError if @digits.size < n

    @digits.chars.each_cons(n).map(&:join)
  end
end

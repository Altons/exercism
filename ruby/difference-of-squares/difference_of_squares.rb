# frozen_string_literal: true

class Squares
  def initialize(n)
    @n = n
  end

  def square_of_sum
    sum = (1 + @n) * @n / 2
    sum * sum
  end

  def sum_of_squares
    @n * (@n + 1) * (2 * @n + 1) / 6
  end

  def difference
    square_of_sum - sum_of_squares
  end
end
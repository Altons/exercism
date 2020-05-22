# frozen_string_literal: true

class Triangle
  def initialize(sides)
    @sides = sides.sort!
  end

  def equilateral?
    return false if @sides.sum.zero?

    @sides.uniq.length == 1
  end

  def isosceles?
    return true if equilateral?

    violate_inequality? && @sides.uniq.length == 2
  end

  def scalene?
    violate_inequality? && @sides.uniq.length == 3
  end

  private

  def violate_inequality?
    a, b, c = @sides
    return false if a + b < c || a + c < b || b + c < a

    true
  end
end

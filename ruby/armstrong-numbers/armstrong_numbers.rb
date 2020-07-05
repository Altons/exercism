# frozen_string_literal: true

class ArmstrongNumbers
  def self.include?(number)
    number == number.to_s.chars.map { |v| v.to_i**number.to_s.length }.sum
  end
end

# frozen_string_literal: true

# Write your code for the 'Resistor Color' exercise in this file. Make the tests in
# `resistor_color_test.rb` pass.
#
# To get started with TDD, see the `README.md` file in your
# `ruby/resistor-color` directory.

class ResistorColor
  COLORS = %w[black brown red orange yellow green blue violet grey white].freeze
  def self.color_code(color)
    COLORS.each_with_index do |v, k|
      return k if color == v
    end
  end
end

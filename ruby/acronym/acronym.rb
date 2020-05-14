# frozen_string_literal: true

# Write your code for the 'Acronym' exercise in this file. Make the tests in
# `acronym_test.rb` pass.
#
# To get started with TDD, see the `README.md` file in your
# `ruby/acronym` directory.
class Acronym
  def self.abbreviate(str)
    str.sub(/-/, ' ').split.map { |c| c[0].upcase }.join
  end
end

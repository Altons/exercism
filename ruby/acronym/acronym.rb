# frozen_string_literal: true

class Acronym
  def self.abbreviate(str)
    str.sub(/-/, ' ').split.map { |c| c[0].upcase }.join
  end
end

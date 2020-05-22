# frozen_string_literal: true

class Pangram
  def self.pangram?(sentence)
    sentence.downcase.chars.uniq.select { |item| item.match(/[a-z]/) }.sort == ('a'..'z').to_a
  end
end

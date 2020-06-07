# frozen_string_literal: true

class Anagram
  def initialize(word)
    @word = word
  end

  def match(list)
    list.select do |v|
      v if (v.downcase == @word.downcase ? 1 : v.downcase.chars.sort <=> @word.downcase.chars.sort).zero?
    end
  end
end

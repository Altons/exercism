# frozen_string_literal: true

class Phrase
  def initialize(str)
    @str = str
  end

  def word_count
    @str.downcase.gsub(/\n|'\B|[^a-zA-Z0-9'? ]|\s'/, ' ').split(' ').tally
  end
end

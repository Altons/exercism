# frozen_string_literal: true

# TwoFer class
class TwoFer
  def self.two_fer(name = nil)
    "One for #{(name ||= '').empty? ? 'you' : name}, one for me."
  end
end

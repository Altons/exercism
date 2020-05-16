# frozen_string_literal: true

class Gigasecond
  GIGASECOND = 1_000_000_000
  def self.from(from_time)
    from_time + GIGASECOND
  end
end

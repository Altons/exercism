# frozen_string_literal: true

class Transpose
  def self.transpose(str)
    return '' if str == ''

    l = str.lines.map(&:size).max
    str.lines.map { |row| row.chomp.ljust(l).chars }.transpose.map(&:join).join("\n").strip
  end
end

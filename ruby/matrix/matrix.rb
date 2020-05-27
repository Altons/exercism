# frozen_string_literal: true

class Matrix
  def initialize(str)
    @str = str
    @rows ||= rows
  end

  def rows
    build_rows
  end

  def columns
    @rows.transpose
  end

  private

  def build_rows
    @str.split("\n").map { |value| value.split(' ').map(&:to_i) }
  end
end

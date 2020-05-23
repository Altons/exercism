# frozen_string_literal: true

class Matrix
  def initialize(str)
    @str = str
  end

  def rows
    build_rows
  end

  def columns
    build_rows.transpose
  end

  private

  def build_rows
    @str.split("\n").map { |value| value.split(' ').map(&:to_i) }
  end
end

# frozen_string_literal: true

class ResistorColorDuo
  COLORS = %w[black brown red orange yellow green blue violet
              grey white].freeze
  def self.value(lst)
    lst.first(2).map { |v| COLORS.index(v) }.join.to_i
  end
end

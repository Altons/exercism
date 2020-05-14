# frozen_string_literal: true

class ResistorColorDuo
  COLORS = { 'black' => '0', 'brown' => '1', 'red' => '0', 'orange' => '3',
             'yellow' => '4', 'green' => '5', 'blue' => '6', 'violet' => '7',
             'grey' => '8', 'white' => '9' }.freeze
  def self.value(lst)
    lst.map { |v| COLORS[v] }[0..1].join.to_i
  end
end

# frozen_string_literal: true

class ResistorColorDuo
  def self.value(lst)
    colors = { 'black' => '0', 'brown' => '1', 'red' => '0', 'orange' => '3', 'yellow' => '4', 'green' => '5',
               'blue' => '6', 'violet' => '7', 'grey' => '8', 'white' => '9' }
    res = ''.dup
    lst.each do |value|
      res << colors[value] if colors.key?(value) && (res.length <= 1)
    end
    res.to_i
  end
end

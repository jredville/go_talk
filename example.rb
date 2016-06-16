require 'logger'

# Main is a wrapper
module Main
  # Foo is an example class
  class Foo
    attr_accessor :foo
    def initialize(foo)
      @foo = foo
      @baz = 12
    end

    def bar
      "I haz #{foo}"
    end

    private

    def baz
      @baz + 3
    end
  end
end

def i_take_a_block
  yield 'world'
end

def i_call_with_a_block
  puts i_take_a_block { |f| 'hello' + f }
end

def enumeration(arr)
  arr.each_with_index do |v, i|
    puts "#{v} is at #{i}"
  end
end

f = Main::Foo.new('abc')

puts f.bar

i_call_with_a_block
enumeration [1, 2, 3, 4]
enumeration %w(1 2 3 4)

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

  def i_take_a_block
    yield 'world'
  end

  def i_call_with_a_block
    puts i_take_a_block do |f|
      'hello' + f
    end
  end

  def emumeration(arr)
    arr.each_with_index do |v, i|
      puts "#{v} is at #{i}"
    end
  end
end

# Module depending on a duck type
module Booable
  def call_boo
    puts boo
  end
end

# Works for Booable
class Booer
  def boo
    'boo'
  end
end

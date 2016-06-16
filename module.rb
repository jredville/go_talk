# Module depending on a duck type
module Booable
  def call_boo
    puts boo
  end
end

# Works for Booable
class Booer
  include Booable
  def boo
    'boo'
  end
end

b = Booer.new
b.call_boo

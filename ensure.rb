class EnsureExample
  def initialize
    @var = 1
  end

  def show_it(&blk)
    change_var
    blk.call @var
  ensure
    reset_var
  end

  def change_var
    @var = 2
  end

  def reset_var
    @var = 1
  end

  def print_var
    puts @var
  end
end

e = EnsureExample.new
e.print_var

e.show_it do |v|
  puts v
  e.print_var
end
e.print_var

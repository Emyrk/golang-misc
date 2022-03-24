levels = ["site", "org", "user", "*", "org:mem", "org:non-mem"]
resources = ["resource", "*", "other"]
ids = ["rid", "other", "*"]
actions = ["action", "other", "*"]

# Build the set P
p =
  for l <- levels, r <- resources, i <- ids, a <- actions do
    %{level: l, resource: r, id: i, action: a}
  end

IO.puts("size of P = #{length(p)}")

# Pull out the set A
a =
  Enum.filter(p, fn x ->
    x.level == "org:non-mem" ||
      x.resource == "other" ||
      x.id == "other" ||
      x.action == "other"
  end)

IO.puts("size of A = #{length(a)}")

# Set j is p - a
j = p -- a
IO.puts("size of J = #{length(j)}")

defmodule Groups do
  @groups [
    %{name: :site, levels: ["site"]},
    %{name: :org, levels: ["org", "org:mem", "org:non-mem"]},
    %{name: :user, levels: ["user"]},
    %{name: :wild, levels: ["*"]}
  ]

  def level_group(level) do
    Enum.find(@groups, fn g -> level in g.levels end)
  end

  def get_groups do
    @groups
  end

  # This function will return the permission set split into the appropriate
  # groups.
  def group(set) do
    Enum.reduce(set, Map.new(@groups, &{&1.name, []}), fn x, acc ->
      g = level_group(x.level)
      acc = Map.put(acc, g.name, [x | acc[g.name]])
      acc
    end)
  end

  def print_grouped(grps) do
    sum =
      Enum.reduce(grps, 0, fn {k, v}, acc ->
        l = length(v)
        IO.puts("\tSize #{Atom.to_string(k)} = #{l}")
        acc + l
      end)

    IO.puts("Sum = #{sum}")
  end
end

IO.puts("------")
IO.puts("Set P")
Groups.group(p) |> Groups.print_grouped()

IO.puts("Set J")
Groups.group(j) |> Groups.print_grouped()

IO.puts("Set A")
Groups.group(a) |> Groups.print_grouped()

# Calculate the size of J{site}, J{org}, J{user}
# Calculate the size of A{site}, A{org}, A{user}

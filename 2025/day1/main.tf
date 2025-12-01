terraform {
  required_version = "~>1.13"
}

locals {
  movements_file    = var.use_small_input ? file("./in_small.txt") : file("./in.txt")
  # movements_file    = file("./in_d.txt")
  movements_split   = split("\n", local.movements_file)
  movements_trunc   = slice(local.movements_split, 0, length(local.movements_split) - 1)
  movements_numbers = concat([50], [for mov in local.movements_trunc : tonumber(replace(replace(mov, "L", "-"), "R", "+"))])
  cumulative_sum = [
    for i, num in local.movements_numbers :
    sum(slice(local.movements_numbers, 0, i + 1))
  ]
  debug1 = [for res in local.cumulative_sum : res % 100 == 0 ? 1 : 0]
  out1   = sum(local.debug1)

  floors = [for res in local.cumulative_sum : floor(res / 100)]

  # Compare subsequent numbers
  valid = [
    for i, res in slice(local.cumulative_sum, 1, length(local.cumulative_sum)) :
    res > local.cumulative_sum[i] ?
    # Case 1: increases
    floor(res / 100) - floor((local.cumulative_sum[i]) / 100) :
    # Case 2: decreases
    abs(floor((res - 1) / 100) - floor((local.cumulative_sum[i] - 1) / 100))
  ]
  out2   = sum(local.valid)
}

// How to detect when, during one step, the dial passes through '0'?
//
// 0 + 101 => 1
// 0 + 99 => 0 (already counting the 0)
// 0 + 230 => 2 (0, 2)
// -1 + 230 => 3 (-1, 2)

output "result1" {
  value = local.out1
}

# output "debug1" {
#   value = local.debug1
# }

output "aaa" {
  value = local.floors
}

output "bbb" {
  value = local.cumulative_sum
}

output "debug2" {
  value = local.valid
}


output "result2" {
  value = local.out2
}

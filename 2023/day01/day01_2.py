nums = [
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
]

num_rev = [s[::-1] for s in nums]

ints = list(range(10))

tot = 0
with open("./in.txt", "r") as f:
    for ln in f:
        firstdigit = -1
        lastdigit = -1
        i = 0
        while i < len(ln):
            cs = ln[:i]

            for n in ints:
                if str(n) in cs:
                    firstdigit = n
                    i = len(ln)

            for j in range(len(nums)):
                if nums[j] in cs:
                    firstdigit = j + 1
                    j = len(nums)
                    i = len(ln)
            i += 1

        i = 0
        while i < len(ln):
            cs2 = ln[len(ln) - i - 1 :]

            for n in ints:
                if str(n) in cs2:
                    lastdigit = n
                    i = len(ln)

            for j in range(len(nums)):
                if nums[j] in cs2:
                    lastdigit = j + 1
                    j = len(nums)
                    i = len(ln)
            i += 1

        curr_add = int("".join([str(firstdigit), str(lastdigit)]))
        tot += curr_add

print(tot)

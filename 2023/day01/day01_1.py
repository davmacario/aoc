tot = 0
with open("./in.txt", "r") as f:
    for l in f:
        firstdigit = -1
        lastdigit = -1

        l_as_lst = list(l)
        ind = 0

        while firstdigit == -1:
            if l_as_lst[ind].isdigit():
                firstdigit = l_as_lst[ind]
            ind += 1

        ind = len(l_as_lst) - 1
        while lastdigit == -1:
            if l_as_lst[ind].isdigit():
                lastdigit = l_as_lst[ind]
            ind -= 1

        tot += int("".join([str(firstdigit), str(lastdigit)]))

print(tot)

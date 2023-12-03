import numpy as np

tot = 0
part_numbers = []
part_numbers_coord = []

with open("./in.txt") as f:
    lines = np.array([list(line.rstrip()) for line in f])
    print(lines)
    n_lines = len(lines)

    for i in range(n_lines):
        j = 0
        while j < lines.shape[1]:  # assume all rows have same len
            # Get char
            c_curr = lines[i, j]
            if c_curr.isdigit():
                # Isolate whole number
                others = []
                ind = j + 1
                while ind < lines.shape[1] and lines[i, ind].isdigit():
                    others.append(lines[i, ind])
                    ind += 1

                # NOTE: the columns of the current number go from j to ind
                # (included)

                n_lst = [c_curr] + others

                # Full "isolated" number:
                n_curr = int("".join(n_lst))

                # Indices of positions around current one:
                around = [
                    (max(0, i - 1), max(0, j - 1)),  # TL
                    (i, max(0, j - 1)),  # CL
                    (min(n_lines - 1, i + 1), max(0, j - 1)),  # BL
                    (max(0, i - 1), min(lines.shape[1] - 1, ind)),  # TR
                    (i, min(lines.shape[1] - 1, ind)),  # CR
                    (
                        min(n_lines - 1, i + 1),
                        min(lines.shape[1] - 1, ind),
                    ),  # BR
                ]
                # Add the ones above and below digits
                for index in range(j, ind):
                    around.append((max(0, i - 1), index))
                    around.append((min(n_lines - 1, i + 1), index))

                to_be_added = False
                for k in around:
                    x = lines[k]
                    if not x.isalpha() and not x.isdigit() and x != ".":
                        to_be_added = True

                if to_be_added:
                    # If here, we have a part number - store its coordinates
                    part_numbers_coord.append(
                        [(i, j_2) for j_2 in range(j, ind)]
                    )
                    part_numbers.append(n_curr)

                j = ind

            else:
                j += 1

    for i in range(n_lines):
        for j in range(lines.shape[1]):
            if lines[i, j] == "*":
                around = [
                    (max(0, i - 1), max(0, j - 1)),  # TL
                    (i, max(0, j - 1)),  # CL
                    (min(n_lines - 1, i + 1), max(0, j - 1)),  # BL
                    (max(0, i - 1), j),  # TC
                    (min(n_lines - 1, i + 1), j),  # BC
                    (max(0, i - 1), min(lines.shape[1] - 1, j + 1)),  # TR
                    (i, min(lines.shape[1] - 1, j + 1)),  # CR
                    (
                        min(n_lines - 1, i + 1),
                        min(lines.shape[1] - 1, j + 1),
                    ),  # BR
                ]

                count_adj = []
                for ind_pn in range(len(part_numbers_coord)):
                    for adj in around:
                        if (
                            adj in part_numbers_coord[ind_pn]
                            and part_numbers[ind_pn] not in count_adj
                        ):
                            count_adj.append(part_numbers[ind_pn])

                if len(count_adj) == 2:
                    tot += count_adj[0] * count_adj[1]

    f.close()

print(tot)

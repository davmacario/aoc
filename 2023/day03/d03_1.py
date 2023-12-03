import numpy as np

tot = 0


def get_around(y: int, x1: int, x2: int):
    around = [
        (max(0, y - 1), max(0, x1 - 1)),  # TL
        (y, max(0, x1 - 1)),  # CL
        (min(n_lines - 1, y + 1), max(0, x1 - 1)),  # BL
        (max(0, y - 1), min(lines.shape[1] - 1, x2)),  # TR
        (y, min(lines.shape[1] - 1, x2)),  # CR
        (
            min(n_lines - 1, y + 1),
            min(lines.shape[1] - 1, x2),
        ),  # BR
    ]
    # Add the ones above and below digits
    for index in range(x1, x2):
        around.append((max(0, y - 1), index))
        around.append((min(n_lines - 1, y + 1), index))

    return around


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

                # ind -= 1

                # NOTE: the columns of the current number go from j to ind
                # (included)
                n_lst = [c_curr] + others

                # Full "isolated" number:
                n_curr = int("".join(n_lst))

                # Indices of positions around current one:
                around = get_around(i, j, ind)
                to_be_added = False

                for k in around:
                    x = lines[k]
                    if not x.isalpha() and not x.isdigit() and x != ".":
                        to_be_added = True

                if to_be_added:
                    tot += n_curr

                j = ind

            else:
                j += 1
    f.close()

print(tot)

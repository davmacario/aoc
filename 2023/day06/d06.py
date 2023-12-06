def optimize(tm, dst):
    """
    Find number of ways record can be beaten

    Args:
        tm: max. race time
        dst: record
    """
    possib = 0
    stop = False
    prev = False  # True if previous `i` yields a W
    i = 0
    while i < tm and not stop:
        if i * (tm - i) > dst:
            prev = True
            possib += 1
        else:
            if prev:
                stop = True
            prev = False
        i += 1

    return possib


if __name__ == "__main__":
    in_file = "in.txt"

    out_1 = 1

    with open(in_file) as f:
        lines = [line.rstrip() for line in f]

        times = []
        dist = []
        for line in lines:
            if line.split(":")[0] == "Time":
                # Get times
                times = [
                    int(x) for x in line.split(":")[1].split(" ") if x.isalnum()
                ]

            elif line.split(":")[0] == "Distance":
                # Get distance
                dist = [
                    int(x) for x in line.split(":")[1].split(" ") if x.isalnum()
                ]

        print("File read!")
        print("Times: ", times)
        print("Dist: ", dist)

        assert len(times) == len(dist)

        for k in range(len(times)):
            curr_t = times[k]  # Max time
            curr_d = dist[k]  # To beat

            out_1 *= optimize(curr_t, curr_d)

        time = int("".join([str(x) for x in times]))
        distance = int("".join([str(x) for x in dist]))

        print("Q2: ", optimize(time, distance))

    print("Q1: ", out_1)

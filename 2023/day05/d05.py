#!/usr/bin/python

import threading

import numpy as np

DEBUG = True
Q1 = False
Q2 = True


def do_mapping(in_val, li, lo, ranges):
    """
    Map the values to the new ones

    ### Args:
        in_val: input values (to be mapped)
        li: source range start
        lo: destination range start
        ranges: list of range lengths
    """
    if DEBUG:
        print("Mapping ...")
    out_val = []
    i = 0
    while i < len(in_val):
        n = in_val[i]
        found_range = False
        for j in range(len(li)):
            if n in range(li[j], li[j] + ranges[j]) and not found_range:
                # Found input range
                found_range = True
                diff_in_start = n - li[j]
                out_val.append(lo[j] + diff_in_start)
        if not found_range:
            # Map 1-to-1
            out_val.append(n)
        i += 1

    return out_val


def heart_function(lines, curr_values, sol_lst, sol_ind):
    """
    Center of the loop - make it thread for each range of seeds

    Args:
        lines: file lines
        curr_values: input values for this instance
        sol_list: ptr to list with minima (solution)
        sol_ind: index of current run into sol_lst
    """

    last_key = "seed"

    values = curr_values

    i = 2  # Line index
    while i < len(lines):
        # Check where we're at - seeds, map (or values)
        if DEBUG:
            print(f"Line {i}")
        ln = lines[i]

        if ln.endswith("map:") and ln.split(" ")[0].split("-")[0] == last_key:
            assert values is not None

            if DEBUG:
                print(f"New map: {ln.split(' ')[0]}")
            # Map beginning!
            list_out = []  # Will contain the input mapping starting values
            list_in = []  # Will contain the output mapping starting values
            ranges = []  # Will contain the ranges
            # -> get the values
            i += 1
            while i < len(lines) and lines[i] != "":
                list_out.append(int(lines[i].split(" ")[0]))
                list_in.append(int(lines[i].split(" ")[1]))
                ranges.append(int(lines[i].split(" ")[2]))
                i += 1

            # Values stored - invoke function
            values = do_mapping(values, list_in, list_out, ranges)
            last_key = ln.split(" ")[0].split("-")[2]
        else:
            raise RuntimeError(f"Shouldn't be here - Line [{i}]:\n{ln}")

        i += 1

    sol_lst[sol_ind] = min(values)


if __name__ == "__main__":
    in_file = "in.txt"

    if Q1:
        with open(in_file) as f:
            lines = [line.rstrip() for line in f]

            i = 0  # Line index
            last_key = None
            values: list = []
            while i < len(lines):
                # Check where we're at - seeds, map (or values)
                if DEBUG:
                    print(f"Line {i}")
                ln = lines[i]
                if ln.split(" ")[0] == "seeds:":
                    # 1st line - get seeds
                    values = [
                        int(x)
                        for x in ln.split(":")[1].split(" ")
                        if x.isalnum()
                    ]
                    last_key = "seed"
                    i += 1

                elif (
                    ln.endswith("map:")
                    and ln.split(" ")[0].split("-")[0] == last_key
                ):
                    assert values != []
                    if DEBUG:
                        print(f"New map: {ln.split(' ')[0]}")
                    # Map beginning!
                    list_out = (
                        []
                    )  # Will contain the input mapping starting values
                    list_in = (
                        []
                    )  # Will contain the output mapping starting values
                    ranges = []  # Will contain the ranges
                    # -> get the values
                    i += 1
                    while i < len(lines) and lines[i] != "":
                        list_out.append(int(lines[i].split(" ")[0]))
                        list_in.append(int(lines[i].split(" ")[1]))
                        ranges.append(int(lines[i].split(" ")[2]))
                        i += 1

                    if DEBUG:
                        print("> Map obtained!")

                    # Values stored - invoke function
                    values = do_mapping(values, list_in, list_out, ranges)
                    last_key = ln.split(" ")[0].split("-")[2]
                else:
                    raise RuntimeError(f"Shouldn't be here - Line [{i}]:\n{ln}")

                i += 1

            f.close()

        print(values)

        print("Sol. Q1:", min(values))

    if Q2:
        with open(in_file) as f:
            lines = [line.rstrip() for line in f]
            values: list = []

            # Extract seed ranges here
            seed_line = lines[0]

            seeds_list = np.array(
                [
                    int(x)
                    for x in seed_line.split(":")[1].split(" ")
                    if x.isalnum()
                ]
            )

            seed_start = seeds_list[0::2]
            seed_range = seeds_list[1::2]

            assert len(seed_start) == len(seed_range)

            mins = [0] * len(seed_start)

            thr_lst = []

            for k in range(len(seed_start)):
                vals_k = list(
                    range(seed_start[k], seed_start[k] + seed_range[k])
                )

                t = threading.Thread(
                    target=heart_function, args=[lines, vals_k, mins, k]
                )
                t.start()
                thr_lst.append(t)

            # Wait for threads to finish
            for th in thr_lst:
                th.join()

        print(f"Q2 sol: {min(mins)}")

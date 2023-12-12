#!/usr/bin/env python3

from typing import List

import numpy as np
from numpy._typing import NDArray

DEBUG = True
QUESTION = 2
SPRINGS_VAL = [".", "#"]


class BinEnum:
    val = 0

    def __init__(self, n_digits: int):
        self.n_digits = n_digits
        self.upd_bin()

    def upd_bin(self):
        # NOTE: the binary value will be a string
        self.val_bin = bin(self.val)[2:].zfill(self.n_digits)

    def increase(self, a: int = 1):
        """Increase the current counter by 'a'"""
        self.val += a
        self.upd_bin()


def checkValidSpring(spr: str, cond: List) -> bool:
    """Check if the current spring is valid"""
    damaged_regions = [x for x in spr.split(".") if x != ""]

    if len(damaged_regions) != len(cond):
        return False

    for i in range(len(damaged_regions)):
        if len(damaged_regions[i]) != int(cond[i]):
            return False

    return True


def count_arrangements(spr, cond):
    """
    Count the possible arrangements in the current line

    BRUTE FORCE APPROACH!
    """
    # Number of possible combinations (to enumerate): 2^(# '?')
    n_qmark = spr.count("?")
    spr_arr = np.array(spr)
    n_possibilities = 2**n_qmark
    pos_qmark = np.argwhere(spr_arr == "?")
    bin_count = BinEnum(n_qmark)

    if DEBUG:
        print(f"N. unknowns: {n_qmark}")
        print(bin_count.val_bin)
        print(spr_arr)

    n_valid = 0
    while bin_count.val < n_possibilities:
        spr_copy = spr_arr.copy()
        curr_bin = [int(x) for x in list(bin_count.val_bin)]

        assert len(curr_bin) == len(pos_qmark)

        for i in range(len(curr_bin)):
            assert spr_copy[pos_qmark[i]] == "?"
            spr_copy[pos_qmark[i]] = SPRINGS_VAL[curr_bin[i]]

        # Check validity:
        if checkValidSpring("".join([x for x in spr_copy]), cond):
            n_valid += 1
            if DEBUG:
                print(n_valid, end="\r")

        bin_count.increase()

    if DEBUG:
        print(n_valid)

    return n_valid


if __name__ == "__main__":
    in_file = "in_small.txt"

    with open(in_file) as f:
        lines = [line.rstrip() for line in f]

        springs = [list(line.split(" ")[0]) for line in lines]
        conditions = [line.split(" ")[1].split(",") for line in lines]

        tot_arr = 0
        for i in range(len(springs)):
            if QUESTION == 2:
                tmp = []
                tmp_c = []
                for j in range(5):
                    tmp += springs[i]
                    tmp_c += conditions[i]
                    if j != 4:
                        tmp.append("?")
                springs[i] = tmp
                conditions[i] = tmp_c
            if DEBUG:
                print(f"Line {i}")
            tot_arr += count_arrangements(springs[i], conditions[i])

        print(f"Q1 - tot arrangements: {tot_arr}")

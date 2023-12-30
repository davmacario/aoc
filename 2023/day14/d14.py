#!/usr/bin/env python3

from typing import List

import numpy as np
from numpy._typing import NDArray

DEBUG = True
QUESTION = 2


def in_range(ind_mov: List[int], dim: List[int]):
    return (dim[0] > ind_mov[0] >= 0) and (dim[1] > ind_mov[1] >= 0)


def shiftRocksNorth(platform: NDArray) -> NDArray:
    """Shift the rocks on the platform towards the "north" direction"""
    dim = list(platform.shape)
    print(dim)
    direction = [-1, 0]

    for i in range(dim[0]):
        for j in range(dim[1]):
            if DEBUG:
                print(platform[i, j], end="")
            if platform[i, j] == "O":
                # Can move the rock
                ind_mov = [i, j]
                ind_mov[0] += direction[0]
                ind_mov[1] += direction[1]
                while (
                    in_range(ind_mov, dim)
                    and platform[ind_mov[0], ind_mov[1]] == "."
                ):
                    ind_mov[0] += direction[0]
                    ind_mov[1] += direction[1]

                ind_mov[0] -= direction[0]
                ind_mov[1] -= direction[1]

                assert (
                    platform[ind_mov[0], ind_mov[1]] in [".", "O"]
                    or ind_mov[0] == 0
                ), f"Pos {ind_mov}: {platform[ind_mov[0], ind_mov[1]]}"

                # Here, ind_mov indicates the final position of the rock
                # Swap the rock and the "."
                platform[i, j] = "."
                platform[ind_mov[0], ind_mov[1]] = "O"
        if DEBUG:
            print("\n")

    return platform


def shiftRocksSouth(platform: NDArray) -> NDArray:
    """Shift the rocks on the platform towards the "south" direction"""
    dim = list(platform.shape)
    print(dim)
    direction = [1, 0]

    for i in range(dim[0], 0, -1):
        for j in range(dim[1]):
            if DEBUG:
                print(platform[i, j], end="")
            if platform[i, j] == "O":
                # Can move the rock
                ind_mov = [i, j]
                ind_mov[0] += direction[0]
                ind_mov[1] += direction[1]
                while (
                    in_range(ind_mov, dim)
                    and platform[ind_mov[0], ind_mov[1]] == "."
                ):
                    ind_mov[0] += direction[0]
                    ind_mov[1] += direction[1]

                ind_mov[0] -= direction[0]
                ind_mov[1] -= direction[1]

                assert (
                    platform[ind_mov[0], ind_mov[1]] in [".", "O"]
                    or ind_mov[0] == 0
                ), f"Pos {ind_mov}: {platform[ind_mov[0], ind_mov[1]]}"

                # Here, ind_mov indicates the final position of the rock
                # Swap the rock and the "."
                platform[i, j] = "."
                platform[ind_mov[0], ind_mov[1]] = "O"
        if DEBUG:
            print("\n")

    return platform


def shiftRocksEast(platform: NDArray) -> NDArray:
    """Shift the rocks on the platform towards the "east" direction"""
    dim = list(platform.shape)
    print(dim)
    direction = [0, 1]

    for i in range(dim[1], 0, -1):
        for j in range(dim[0]):
            if DEBUG:
                print(platform[j, i], end="")
            if platform[j, i] == "O":
                # Can move the rock
                ind_mov = [j, i]
                ind_mov[0] += direction[0]
                ind_mov[1] += direction[1]
                while (
                    in_range(ind_mov, dim)
                    and platform[ind_mov[0], ind_mov[1]] == "."
                ):
                    ind_mov[0] += direction[0]
                    ind_mov[1] += direction[1]

                ind_mov[0] -= direction[0]
                ind_mov[1] -= direction[1]

                assert (
                    platform[ind_mov[0], ind_mov[1]] in [".", "O"]
                    or ind_mov[0] == 0
                ), f"Pos {ind_mov}: {platform[ind_mov[0], ind_mov[1]]}"

                # Here, ind_mov indicates the final position of the rock
                # Swap the rock and the "."
                platform[j, i] = "."
                platform[ind_mov[0], ind_mov[1]] = "O"
        if DEBUG:
            print("\n")

    return platform


def shiftRocksWest(platform: NDArray) -> NDArray:
    """Shift the rocks on the platform towards the "west" direction"""
    dim = list(platform.shape)
    print(dim)
    direction = [0, -1]

    for i in range(dim[1]):
        for j in range(dim[0]):
            if DEBUG:
                print(platform[j, i], end="")
            if platform[j, i] == "O":
                # Can move the rock
                ind_mov = [j, i]
                ind_mov[0] += direction[0]
                ind_mov[1] += direction[1]
                while (
                    in_range(ind_mov, dim)
                    and platform[ind_mov[0], ind_mov[1]] == "."
                ):
                    ind_mov[0] += direction[0]
                    ind_mov[1] += direction[1]

                ind_mov[0] -= direction[0]
                ind_mov[1] -= direction[1]

                assert (
                    platform[ind_mov[0], ind_mov[1]] in [".", "O"]
                    or ind_mov[0] == 0
                ), f"Pos {ind_mov}: {platform[ind_mov[0], ind_mov[1]]}"

                # Here, ind_mov indicates the final position of the rock
                # Swap the rock and the "."
                platform[j, i] = "."
                platform[ind_mov[0], ind_mov[1]] = "O"
        if DEBUG:
            print("\n")

    return platform


def shiftNWSE(pattern: NDArray) -> NDArray:
    """
    Shift the pattern North, West, South and East.
    Returns updated pattern.
    """
    shifted_n = shiftRocksNorth(pattern)
    shifted_w = shiftRocksWest(shifted_n)
    shifted_s = shiftRocksSouth(shifted_w)
    shifted_e = shiftRocksEast(shifted_s)
    return shifted_e


if __name__ == "__main__":
    in_file = "in.txt"

    with open(in_file) as f:
        # Isolate the pattern
        pattern = []
        for line in f:
            ln = line.rstrip()
            pattern.append([str(x) for x in ln])

        f.close()

    rocks_init = np.array(pattern)

    if QUESTION == 1:
        # Call function to move all rocks to the north
        shifted_pattern = shiftRocksNorth(rocks_init)

        if DEBUG:
            print(shifted_pattern)

        q1_total = 0
        for i in range(shifted_pattern.shape[0]):
            # Count the number of "O" in the line and multiply the value
            n_rocks = np.sum(shifted_pattern[i, :] == "O")
            if DEBUG:
                print(n_rocks)
            q1_total += n_rocks * (shifted_pattern.shape[0] - i)

        print(q1_total)
    elif QUESTION == 2:
        for ind in range(1000000000):
            if ind == 0:
                final_pattern = shiftNWSE(rocks_init)
                if DEBUG:
                    
            else:
                final_pattern = shiftNWSE()

        # Evaluate load

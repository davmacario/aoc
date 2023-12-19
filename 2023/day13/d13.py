#!/usr/bin/env python3

from typing import List

import numpy as np
from numpy._typing import NDArray

DEBUG = True
QUESTION = 2


def findReflection(pattern: NDArray, direction: int = 0) -> int:
    """
    Find a reflectionm in the pattern, along the specified direction.

    Args:
        pattern: NDArray containing the pattern to be analyzed
        direction: integer specifying the direction along which to look for the
        reflection: 0: look for horizontal reflections along y; 1: look for
        vertical reflections along x

    Returns:
        The integer specifying the index of the element *after* which we have a
        reflection.
        0 if no reflection was found
    """
    # Find reflection
    i = 0
    if DEBUG:
        print(pattern.shape)
    while i < pattern.shape[direction]:
        if DEBUG:
            if not direction:
                print(f"> Row {i}")
            elif direction:
                print(f"> Column {i}")
        symm = True
        j = 0
        while j <= min(i, pattern.shape[direction] - i - 2):
            # for j in range(min(i, pattern.shape[direction] - i)):
            if not direction:
                # if DEBUG:
                #     print(f"horizontal - j = {j}; print rows {i - j} and {i + 1 + j}")
                #     print(pattern[i - j, :])
                #     print(pattern[i + 1 + j, :])
                if not all(pattern[i - j, :] == pattern[i + 1 + j, :]):
                    symm = False
            else:
                # if DEBUG:
                #     print(f"vertical - j = {j}; print columns {i - j} and {i + 1 + j}")
                #     print(pattern[:, i - j])
                #     print(pattern[:, i + 1 + j])
                if not all(pattern[:, i - j] == pattern[:, i + 1 + j]):
                    symm = False
            j += 1
        if symm and i < pattern.shape[direction] - 1:
            return i + 1
        i += 1
    return 0


def findReflectionSmudges(pattern: NDArray, direction: int = 0) -> int:
    """
    Find a reflectionm in the pattern, along the specified direction.
    The reflection will be considered valid even if 1 smudge is present (1 char
    that is not symmetric)

    Args:
        pattern: NDArray containing the pattern to be analyzed
        direction: integer specifying the direction along which to look for the
        reflection: 0: look for horizontal reflections along y; 1: look for
        vertical reflections along x

    Returns:
        The integer specifying the index of the element *after* which we have a
        reflection.
        0 if no reflection was found
    """
    # Find reflection
    i = 0
    if DEBUG:
        print(pattern.shape)
    while i < pattern.shape[direction]:
        smudges = 1
        if DEBUG:
            if not direction:
                print(f"> Row {i}")
            elif direction:
                print(f"> Column {i}")
        symm = True
        j = 0
        while j <= min(i, pattern.shape[direction] - i - 2) and smudges >= 0:
            if not direction:
                if (
                    np.sum(pattern[i - j, :] != pattern[i + 1 + j, :])
                    == smudges
                    and smudges > 0
                ):
                    if DEBUG:
                        print("Found smudge")
                    smudges -= 1
                elif not all(pattern[i - j, :] == pattern[i + 1 + j, :]):
                    symm = False
            else:
                if (
                    np.sum(pattern[:, i - j] != pattern[:, i + 1 + j])
                    == smudges
                    and smudges > 0
                ):
                    if DEBUG:
                        print("Found smudge")
                    smudges -= 1
                elif not all(pattern[:, i - j] == pattern[:, i + 1 + j]):
                    symm = False
            j += 1
        if symm and i < pattern.shape[direction] - 1 and smudges == 0:
            return i + 1
        i += 1
    return 0


if __name__ == "__main__":
    in_file = "in.txt"

    with open(in_file) as f:
        # Isolate the individual patterns
        patterns = []
        new_pattern_rows = []
        for line in f:
            ln = line.rstrip()
            # lines = [line.rstrip() for line in f]
            if list(ln) == []:
                patterns.append(np.array(new_pattern_rows))
                new_pattern_rows = []
            else:
                # Append the new row as a list
                new_pattern_rows.append([x for x in ln])
        patterns.append(np.array(new_pattern_rows))
        f.close()

    ans = 0
    if DEBUG:
        print(f"N. patterns: {len(patterns)}")
    for pattern in patterns:
        # if DEBUG:
        #     print(pattern)

        if QUESTION == 1:
            sh = findReflection(pattern)
        elif QUESTION == 2:
            sh = findReflectionSmudges(pattern)

        if DEBUG:
            print(f"Horizontal symmetry: {sh}")

        if sh > 0:
            ans += 100 * sh
        else:
            if QUESTION == 1:
                sv = findReflection(pattern, direction=1)
            elif QUESTION == 2:
                sv = findReflectionSmudges(pattern, direction=1)
            if DEBUG:
                print(f"Vertical symmetry: {sv}")
            if sv > 0:
                ans += sv
            else:
                raise ValueError("No reflection was detected!")

    print(f"Q1 - answer: {ans}")

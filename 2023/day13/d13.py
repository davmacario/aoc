#!/usr/bin/env python3

from typing import List

import numpy as np
from numpy._typing import NDArray

DEBUG = True


def findReflection(pattern: NDArray, direction: int = 0) -> int:
    """
    Find a reflectionm in the pattern, along the specified direction.

    Args:
        pattern: NDArray containing the pattern to be analyzed
        direction: integer specifying the direction along which to look for the
        reflection (0: look for horizontal reflections along y, 1: look for
        vertical reflections alongx)

    Returns:
        The integer specifying the index of the element *after* which we have a
        reflection.
        0 if no reflection was found
    """
    # Find reflection
    i = 1
    if DEBUG:
        print(pattern.shape)
    while i < pattern.shape[direction]:
        if DEBUG:
            print(i)
        symm = True
        for j in range(min(i, pattern.shape[direction] - i)):
            if not direction:
                if DEBUG:
                    print(pattern[i - j, :])
                    print(pattern[i + 1 + j, :])
                if pattern[i - j, :] != pattern[i + 1 + j, :]:
                    symm = False
            else:
                if DEBUG:
                    print(pattern[:, i - j])
                    print(pattern[:, i + 1 + j])
                if pattern[:, i - j] != pattern[:, i + 1 + j]:
                    symm = False
        if symm:
            return i
        i += 1
    return 0


if __name__ == "__main__":
    in_file = "in_small.txt"

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

    q1_ans = 0
    if DEBUG:
        print(f"N. patterns: {len(patterns)}")
    for pattern in patterns:
        if DEBUG:
            print(pattern)
        sh = findReflection(pattern)
        if DEBUG:
            print(f"Horizontal symmetry: {sh}")
        if sh > 0:
            q1_ans += 100 * sh
        else:
            sv = findReflection(pattern, direction=1)
            if DEBUG:
                print(f"Vertical symmetry: {sv}")
            if sv > 0:
                q1_ans += sv
            else:
                raise ValueError("No reflection was detected!")

    print(f"Q1 - answer: {q1_ans}")

#!/usr/bin/env python3

from typing import List, Tuple

import numpy as np
from numpy._typing import NDArray

DEBUG = True
QUESTION = 2


def expandSpace(mat: NDArray, char_empty: str = ".") -> NDArray:
    """
    Expand rows and columns containing all `char_empty` (default: ".")

    Args:
        mat: initial matrix to be expanded (NDArray)
        char_empty: char to be looked for indicating "empty" lines/columns

    Returns:
        NDArray containing the expanded matrix
    """
    out_map = mat.copy()
    # Expand rows
    i = 0
    while i < out_map.shape[0]:
        if all(out_map[i, :] == char_empty):
            # Add new row
            out_map = np.insert(out_map, i, out_map[i, :], axis=0)
            i += 1
        i += 1

    i = 0
    while i < out_map.shape[1]:
        if all(out_map[:, i] == char_empty):
            # Add new col
            out_map = np.insert(out_map, i, out_map[:, i], axis=1)
            i += 1
        i += 1

    return out_map


def countGalaxies(mat: NDArray, galaxy_char: str = "#") -> Tuple[NDArray, int]:
    """
    Count the number of galaxies and replace the character with the
    corresponding value

    Args:
        mat: input matrix (space)
        galaxy_char (default "#"): char indicating a galaxy

    Returns:
        Matrix with galaxies replaced by their number
        Number of total galaxies
    """
    count = 0
    out_mat = mat.copy()
    for i in range(out_mat.shape[0]):
        for j in range(out_mat.shape[1]):
            if out_mat[i, j] == galaxy_char:
                count += 1
                # out_mat[i, j] = str(count)
                out_mat[i, j] = count

    return out_mat, count


def countEmptyRC(
    mat: NDArray, start: NDArray, end: NDArray, char_empty: str = "."
) -> Tuple[int, int]:
    """
    Count the number of empty rows and columns between starting point and
    destination in a given matrix

    Args:
        mat: matrix
        start: starting point
        end: destination point
        char_empty: character indicating "empty" elements in the matrix

    Returns:
        number of empty rows and empty columns between start and end
    """
    count_e_r = 0
    count_e_c = 0

    for i in range(min(start[0], end[0]) + 1, max(start[0], end[0])):
        if all(mat[i, :] == char_empty):
            count_e_r += 1

    for i in range(min(start[1], end[1]) + 1, max(start[1], end[1])):
        if all(mat[:, i] == char_empty):
            count_e_c += 1

    return count_e_r, count_e_c


def eval_dist(
    mat: NDArray, start: NDArray, end: NDArray, mul_factor_empty: int = 2
) -> int:
    """
    Evaluate the distance between the start and the end in the matrix

    Args:
        mat: matrix
        start: starting point
        end: destination point
        mul_factor_empty: number of empty lines instead of each

    Returns:
        Distance between the start and end points
    """
    n_rows_empty, n_cols_empty = countEmptyRC(mat, start, end)

    dist = np.sum(np.abs(end - start))
    dist += n_rows_empty * (mul_factor_empty - 1)
    dist += n_cols_empty * (mul_factor_empty - 1)

    return dist


if __name__ == "__main__":
    in_file = "in.txt"

    with open(in_file) as f:
        lines = [line.rstrip() for line in f]
        space = np.array([list(line) for line in lines], dtype="<U9")
        if DEBUG:
            print(space)

        # Count galaxies:
        space_count, n_galaxies = countGalaxies(space)

        if DEBUG:
            print(space_count)

        out_space_count = "./space_count.txt"
        with open(out_space_count, "w") as f_out:
            for i in range(space_count.shape[0]):
                for j in range(space_count.shape[1]):
                    f_out.write(str(space_count[i, j]))
                f_out.write("\n")
            f_out.close()

        # Expand the rows and columns
        space_expanded = expandSpace(space_count)

        if DEBUG:
            print(space_expanded)

        dist_mat = np.zeros((n_galaxies, n_galaxies), dtype=np.uint64)

        tot_1 = 0
        for i in range(n_galaxies):
            for j in range(i + 1, n_galaxies):
                # Element i, j of dist_mat contains the distance between
                # galaxies i+1 and j+1 (same for element j, i)
                pos_i = np.argwhere(space_count == str(i + 1))[0]
                pos_j = np.argwhere(space_count == str(j + 1))[0]

                if QUESTION == 1:
                    dist = eval_dist(space_count, pos_i, pos_j)
                elif QUESTION == 2:
                    dist = eval_dist(
                        space_count, pos_i, pos_j, mul_factor_empty=1000000
                    )
                else:
                    raise ValueError("Wrong question number!")

                if DEBUG:
                    print(f"Dist. between galaxy {i+1} and {j+1}: {dist}")

                dist_mat[i, j] = dist_mat[j, i] = dist
                tot_1 += dist

        print(f"Q1 - sum of dist: {tot_1}")

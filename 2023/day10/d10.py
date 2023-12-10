#!/usr/bin/env python3

import os
import sys
from typing import List

import cv2
import matplotlib.pyplot as plt
import numpy as np
from numpy._typing import NDArray

DEBUG = True

sys.setrecursionlimit(1000000)

orientations = ["N", "S", "W", "E"]
wasd = []
pipes = {
    "|": "NS",
    "-": "WE",
    "L": "NE",
    "J": "NW",
    "7": "SW",
    "F": "SE",
    ".": "",
    "S": "SOL",
}
pipes_dir = {
    "|": ["N", "S"],
    "-": ["W", "E"],
    "L": ["N", "E"],
    "J": ["N", "W"],
    "7": ["S", "W"],
    "F": ["S", "E"],
}
rel_pos = {"N": [-1, 0], "S": [1, 0], "W": [0, -1], "E": [0, 1]}


def travelLoop(mat: NDArray, start: List[int], start_dir: List[int]) -> int:
    """
    Travel along the pipe loop, and return the full path length.

    Args:
        mat: NDArray containing the map
        start: coordinates of the starting point ("S")
        start_dir: coordinates of the first point after the origin

    Returns:
        Length of the full loop
    """
    path_len = 1
    # This stores the char in the current position:
    new_char = mat[start_dir[0], start_dir[1]]
    prev_pos = start
    curr_pos = start_dir

    assert mat[prev_pos[0], prev_pos[1]] == "S"

    while new_char != "S":
        # Move along the direction defined by the current char using pipes_dir
        # 1. Find out where we came from (compare with prev. position)
        arr_dir = [curr_pos[0] - prev_pos[0], curr_pos[1] - prev_pos[1]]

        if arr_dir == [1, 0]:
            arr_card = "N"
        elif arr_dir == [0, 1]:
            arr_card = "W"
        elif arr_dir == [-1, 0]:
            arr_card = "S"
        elif arr_dir == [0, -1]:
            arr_card = "E"
        else:
            raise ValueError(
                "The previous position and the current one are not adjacent"
            )

        if DEBUG:
            print(f"Arrival direction: {arr_card}")

        next_directions_lst = pipes_dir[new_char]

        assert arr_card in next_directions_lst

        # This list should always contain only 1 element
        next_card = [x for x in next_directions_lst if x != arr_card][0]

        if DEBUG:
            print(f"Next direction: {next_card}")

        prev_pos = curr_pos
        curr_pos = [
            curr_pos[0] + rel_pos[next_card][0],
            curr_pos[1] + rel_pos[next_card][1],
        ]
        new_char = mat[curr_pos[0], curr_pos[1]]
        path_len += 1

        if DEBUG:
            print(f"Moved to {curr_pos}")
            print()

    return path_len


def getFullPath(
    mat: NDArray, start: List[int], start_dir: List[int]
) -> List[List[int]]:
    """
    Travel along the pipe loop, and return the full path (coordinates).

    Args:
        mat: NDArray containing the map
        start: coordinates of the starting point ("S")
        start_dir: coordinates of the first point after the origin

    Returns:
        Ordered list of the traversed points
    """
    path_points = [start_dir]  # The starting point "S" is added last
    # This stores the char in the current position:
    new_char = mat[start_dir[0], start_dir[1]]
    prev_pos = start
    curr_pos = start_dir

    assert mat[prev_pos[0], prev_pos[1]] == "S"

    while new_char != "S":
        # Move along the direction defined by the current char using pipes_dir
        # 1. Find out where we came from (compare with prev. position)
        arr_dir = [curr_pos[0] - prev_pos[0], curr_pos[1] - prev_pos[1]]

        if arr_dir == [1, 0]:
            arr_card = "N"
        elif arr_dir == [0, 1]:
            arr_card = "W"
        elif arr_dir == [-1, 0]:
            arr_card = "S"
        elif arr_dir == [0, -1]:
            arr_card = "E"
        else:
            raise ValueError(
                "The previous position and the current one are not adjacent"
            )

        if DEBUG:
            print(f"Arrival direction: {arr_card}")

        next_directions_lst = pipes_dir[new_char]

        assert arr_card in next_directions_lst

        # This list should always contain only 1 element
        next_card = [x for x in next_directions_lst if x != arr_card][0]

        if DEBUG:
            print(f"Next direction: {next_card}")

        prev_pos = curr_pos
        curr_pos = [
            curr_pos[0] + rel_pos[next_card][0],
            curr_pos[1] + rel_pos[next_card][1],
        ]
        new_char = mat[curr_pos[0], curr_pos[1]]
        path_points.append(curr_pos)

        if DEBUG:
            print(f"Moved to {curr_pos}")
            print()

    return path_points


def areaFill(
    matrix: NDArray, start: List[int], subst: int = 0, subst_with: int = 2
) -> NDArray:
    """
    Given a matrix, fill the values equal to `subst` with `subst_with`, starting
    from `start`.

    This function is recursive

    Args:
        matrix: NDArray to be filled
        start: starting position (x and y indices of matrix)
        subst: values to be substituted
        subst_with: values that are placed

    Returns:
        The matrix with substituted values (filled)
    """

    if matrix[start[0], start[1]] != subst:
        return matrix

    matrix[start[0], start[1]] = subst_with

    if start[0] - 1 >= 0:
        matrix = areaFill(matrix, [start[0] - 1, start[1]])

    if start[1] - 1 >= 0:
        matrix = areaFill(matrix, [start[0], start[1] - 1])

    if start[0] + 1 <= matrix.shape[0] - 1:
        matrix = areaFill(matrix, [start[0] + 1, start[1]])

    if start[1] + 1 <= matrix.shape[1] - 1:
        matrix = areaFill(matrix, [start[0], start[1] + 1])

    return matrix


if __name__ == "__main__":
    in_file = "in.txt"

    with open(in_file) as f:
        lines = [line.rstrip() for line in f]
        chars = np.array([list(line) for line in lines])

        # Find source
        pos_start = np.argwhere(chars == "S")[0]

        if DEBUG:
            print(pos_start)

        # Find connected pipes (exactly 2, since loop)
        conn_pipes_pos = []
        if pos_start[0] - 1 >= 0 and pipes[
            chars[pos_start[0] - 1, pos_start[1]]
        ] in ["NS", "SW", "SE"]:
            # Connection at the top
            conn_pipes_pos.append([pos_start[0] - 1, pos_start[1]])

        if pos_start[1] - 1 >= 0 and pipes[
            chars[pos_start[0], pos_start[1] - 1]
        ] in ["WE", "NE", "SE"]:
            # Connection at left
            conn_pipes_pos.append([pos_start[0], pos_start[1] - 1])

        if pos_start[0] + 1 < chars.shape[0] and pipes[
            chars[pos_start[0] + 1, pos_start[1]]
        ] in ["NS", "NE", "NW"]:
            # Connection at the bottom
            conn_pipes_pos.append([pos_start[0] + 1, pos_start[1]])

        if pos_start[1] + 1 < chars.shape[1] and pipes[
            chars[pos_start[0], pos_start[1] + 1]
        ] in ["WE", "NW", "SW"]:
            if DEBUG:
                print("HERE")
            conn_pipes_pos.append([pos_start[0], pos_start[1] + 1])

        if DEBUG:
            print(
                chars[
                    pos_start[0] - 1 : pos_start[0] + 2,
                    pos_start[1] - 1 : pos_start[1] + 2,
                ]
            )

        assert len(conn_pipes_pos) == 2, f"{conn_pipes_pos}"

        # Evaluate path length from both directions - should be path length from
        # one direction // 2
        len_path = travelLoop(chars, pos_start, conn_pipes_pos[0])
        print(f"Q1 - max len: {len_path // 2}")

        # QUESTION 2: evaluate the number of tiles enclosed in the loop

        # Idea: get the coordinates of the path
        path = getFullPath(chars, pos_start, conn_pipes_pos[0])
        path_arr = np.array(path)

        # Find clever way to isolate the points *inside* the loop...
        # Use np.ufunc.accumulate to "fill" the contour with "I"

        # First, create a matrix of "0" with ones on the path
        mat_mask = np.zeros(chars.shape, dtype=np.uint8)
        mat_mask[path_arr[:, 0], path_arr[:, 1]] = 1

        # Idea: upsample by 2 the matrix with the mask - need to upsample the
        # path as well:
        double_path = [[2 * p[0], 2 * p[1]] for p in path]
        ups_path = []
        for i in range(1, len(double_path)):
            ups_path.append(double_path[i - 1])
            ups_path.append(
                [
                    (double_path[i - 1][0] + double_path[i][0]) // 2,
                    (double_path[i - 1][1] + double_path[i][1]) // 2,
                ]
            )

        # Add the one completing the path (avg. last and first of double_path)
        ups_path.append(
            [
                (double_path[-1][0] + double_path[0][0]) // 2,
                (double_path[-1][1] + double_path[0][1]) // 2,
            ]
        )
        ups_path.append([double_path[-1][0], double_path[-1][1]])

        ups_path_arr = np.array(ups_path)

        mask_ups = np.zeros(
            (2 * chars.shape[0], 2 * chars.shape[1]), dtype=np.uint8
        )
        mask_ups[ups_path_arr[:, 0], ups_path_arr[:, 1]] = 1

        # Zero padding at the boundary
        mask_ups = np.pad(
            mask_ups,
            ((1, 1), (1, 1)),
            "constant",
            constant_values=((0, 0), (0, 0)),
        )

        plt.figure()
        plt.imshow(mask_ups, cmap="gray")
        plt.show()

        out_mask = "out_mask.txt"
        with open(out_mask, "w") as f_out:
            for i in range(mask_ups.shape[0]):
                for j in range(mask_ups.shape[1]):
                    f_out.write(str(mask_ups[i, j]))
                f_out.write("\n")
            f_out.close()

        first_zero = [0, 0]

        # Fill the matrix starting from (0,0) - recursive function
        # filled_ups_mat = areaFill(mask_ups, start=(first_zero))

        # Alternative: use floodfill - opencv
        im_fill = 255 * mask_ups  # Convert to 0-255
        mask = np.zeros(
            (im_fill.shape[0] + 2, im_fill.shape[1] + 2), dtype=np.uint8
        )
        cv2.floodFill(im_fill, mask, first_zero, 100)

        plt.figure()
        plt.imshow(im_fill, cmap="gray")
        plt.show()

        filled_ups_mat = im_fill

        out_map = "out_map.txt"
        with open(out_map, "w") as f_out:
            for i in range(filled_ups_mat.shape[0]):
                for j in range(filled_ups_mat.shape[1]):
                    f_out.write(str(filled_ups_mat[i, j]))
                f_out.write("\n")
            f_out.close()

        # Count the values that have not been filled - inside loop, still "0"
        # and if they are completely surrounded by 0 (even diagonally), they
        # were inside before

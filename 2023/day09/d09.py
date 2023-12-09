#!/usr/bin/env python3

from typing import List

DEBUG = False


def find_next(sequence: List[int]) -> int:
    """Find the next element in the sequence, using the differences rule"""
    if not all(x == 0 for x in sequence):
        if DEBUG:
            print(sequence)
        diff_vec = [
            sequence[i] - sequence[i - 1] for i in range(1, len(sequence))
        ]
        return sequence[-1] + find_next(diff_vec)
    else:
        if DEBUG:
            print("> ", sequence)
        return 0


def find_prev(sequence: List[int]) -> int:
    """Find the next element in the sequence, using the differences rule"""
    if not all(x == 0 for x in sequence):
        diff_vec = [
            sequence[i] - sequence[i - 1] for i in range(1, len(sequence))
        ]
        return sequence[0] - find_prev(diff_vec)
    else:
        return 0


total = 0
total_prev = 0

if __name__ == "__main__":
    in_file = "in.txt"

    with open(in_file) as f:
        lines = [line.rstrip() for line in f]

        for line in lines:
            line_lst = [int(x) for x in line.split(" ") if x != " "]
            next_elem = find_next(line_lst)
            if DEBUG:
                # print(f"Next element for {line_lst}: {next_elem}")
                print(f"Next element: {next_elem}")
            total += next_elem

            # Question 2
            prev_elem = find_prev(line_lst)
            if DEBUG:
                print(f"Previous element: {prev_elem}")
            total_prev += prev_elem

        print(f"Q1: {total}")
        print(f"Q2: {total_prev}")

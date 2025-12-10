import argparse
import heapq
import math
from dataclasses import dataclass
from pathlib import Path
from typing import List, Tuple

def get_area(p1, p2) -> int:
    return (abs(p1[0] - p2[0]) + 1) * (abs(p1[1] - p2[1]) + 1)

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--small", action="store_true", help="if set, use small input")
    args = parser.parse_args()

    if args.small:
        file = Path("./in_small.txt")
    else:
        file = Path("./in.txt")

    max_y = max_x = 0
    points: List[Tuple[int, int]] = []
    with open(file) as f:
        for line in f:
            coords = [int(n) for n in line.rstrip().split(",")]
            points.append((coords[0], coords[1]))

            if coords[0] > max_y:
                max_y = coords[0]
            if coords[1] > max_x:
                max_x = coords[1]

    n_points = len(points)
    # h = max_y
    # w = max_x

    result1 = 0
    for i in range(n_points):
        for j in range(i + 1, n_points):
            curr_a = get_area(points[i], points[j])
            if curr_a > result1:
                result1 = curr_a

    print(f"Result 1: {result1}")

if __name__ == "__main__":
    main()

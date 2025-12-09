import argparse
import heapq
import math
from dataclasses import dataclass
from pathlib import Path
from typing import List, Tuple


@dataclass
class Point3D:
    x: float
    y: float
    z: float


def euclidean_distance(p1: Point3D, p2: Point3D) -> float:
    return math.sqrt((p1.x - p2.x) ** 2 + (p1.y - p2.y) ** 2 + (p1.z - p2.z) ** 2)

def get_dist_heap(points: List[Point3D]) -> List[Tuple[float, int, int]]:
    distances: List[Tuple[float, int, int]] = []
    for i in range(len(points)):
        for j in range(i+1, len(points)):
            # Only compute distance (point i) -> (point j) where j > i
            heapq.heappush(distances, (euclidean_distance(points[i], points[j]), i, j))
    return distances

def find(n: int, parents: List[int]) -> int:
    """Returns root of node `n` given `parents` array"""
    curr_n = n
    # Root <=> parent of itself
    while curr_n != parents[curr_n]:
        curr_n = parents[curr_n]
    return curr_n

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--small", action="store_true", help="if set, use small input")
    args = parser.parse_args()

    if args.small:
        file = Path("./in_small.txt")
        n_iter = 10
    else:
        file = Path("./in.txt")
        n_iter = 1000

    points: List[Point3D] = []
    with open(file) as f:
        for line in f:
            coords = [int(n) for n in line.rstrip().split(",")]
            points.append(Point3D(x=coords[0], y=coords[1], z=coords[2]))

    distances = get_dist_heap(points)
    setSizes = [1] * len(points)
    parents = [i for i in range(len(points))]
    iter = 0
    while sum([1 if n > 0 else 0 for n in setSizes]) > 1:
        # At each pass, pop from heap
        dist, i, j = heapq.heappop(distances)
        print(f"{i} -> {j}")

        root_i = find(i, parents)
        root_j = find(j, parents)

        if root_i != root_j:
            parents[root_j] = root_i
            setSizes[root_i] += setSizes[root_j]
            setSizes[root_j] = 0

        if iter == n_iter - 1:
            setSizesSort = sorted(setSizes)
            result1 = setSizesSort[-1] * setSizesSort[-2] * setSizesSort[-3]

        iter += 1

    result2 = points[i].x * points[j].x

    setSizes.sort()

    print(f"Part 1: {result1}")
    print(f"Part 2: {result2}")

if __name__ == "__main__":
    main()

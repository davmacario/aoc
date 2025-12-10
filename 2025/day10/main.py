import argparse
import re
from collections import deque
from dataclasses import dataclass
from pathlib import Path
from typing import List, Tuple, Union


def lights_to_int(lights: str) -> int:
    lights_rep = lights.replace(".", "0").replace("#", "1")
    return int(lights_rep, 2)


def str_to_buttons(
    buttons: str, n_bits: int
) -> Tuple[List[int], List[Tuple[int, ...]]]:
    buttons_value: List[int] = []
    buttons_list: List[Tuple[int, ...]] = []
    for b in buttons.split(" "):
        in_parentheses = b[1:-1]
        flip_bit_pos = [int(n) for n in in_parentheses.split(",")]
        buttons_list.append(tuple(flip_bit_pos))
        new_button_list = ["0"] * n_bits

        for fp in flip_bit_pos:
            new_button_list[fp] = "1"

        new_button_bin = int("".join(new_button_list), 2)
        buttons_value.append(new_button_bin)

    return buttons_value, buttons_list


def str_to_joltages(jolt_str: str) -> Tuple[int, ...]:
    return tuple(int(x) for x in jolt_str.split(","))


def press_button_jolt(
    curr_j: Tuple[int, ...], button: Tuple[int, ...]
) -> Tuple[int, ...]:
    out = list(curr_j)

    for b_i in button:
        out[b_i] += 1

    return tuple(out)


@dataclass
class Machine:
    n_lights: int
    target_str: str
    target: int
    buttons: List[int]
    buttons_list: List[Tuple[int, ...]]
    joltages: Tuple[int, ...]

    @classmethod
    def from_line(cls, line: str):
        matches = re.match(r"^\[([.#]+)\](.+)\{(.+)\}", line)
        assert matches is not None
        groups = matches.groups()

        target_str = groups[0]
        buttons_str = groups[1].strip()
        joltage_str = groups[2]

        target = lights_to_int(target_str)
        buttons, buttons_list = str_to_buttons(buttons_str, len(target_str))
        joltages = str_to_joltages(joltage_str)

        return cls(
            n_lights=len(target_str),
            target_str=target_str,
            target=target,
            buttons=buttons,
            buttons_list=buttons_list,
            joltages=joltages,
        )

    def solve_lights(self, value: int) -> int | float:
        """
        Returns the min number of button presses to achieve the desired target str

        Approach: bfs
        - If current value matches target,
        - At each call, try to press each button (index i)
          - Pressing button <-> XORing with button value

        optimization: cache seen. If already seen, return -1, i.e., deadlock (not sol)
        """
        cache = set()

        # Queue items:
        #  - value
        #  - index
        lifo = deque([(0, 0)])

        while lifo:
            value, ind = lifo.popleft()
            if value == self.target:
                return ind
            for b in self.buttons:
                curr = value ^ b
                if curr in cache:
                    continue
                cache.add(curr)
                lifo.append((curr, ind + 1))

        return float("inf")

    def solve_joltages(self):
        """
        Similar to lights


        """
        cache = set()

        start = tuple([0] * len(self.joltages))
        lifo = deque([(start, 0)])

        while lifo:
            curr_j, ind = lifo.popleft()

            if curr_j == self.joltages:
                return ind

            for b in self.buttons_list:
                next = press_button_jolt(curr_j, b)
                print(next)

                if next in cache:
                    continue

                cache.add(next)
                # Optimize: if any value of joltage is > corresp. value in target, return
                if any([next[i] > self.joltages[i] for i in range(self.n_lights)]):
                    continue

                lifo.append((next, ind + 1))

        return float("inf")



def min_number_presses(line: str) -> Tuple[int|float, int|float]:
    machine = Machine.from_line(line)
    print(machine)
    min_presses_l = machine.solve_lights(0)
    if min_presses_l == float("inf"):
        print("ERROR")
        min_presses_l = 10**9

    min_presses_j = machine.solve_joltages()
    if min_presses_j == float("inf"):
        print("ERROR")
        min_presses_j = 10**9

    return min_presses_l, min_presses_j


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--small", action="store_true", help="if set, use small input")
    args = parser.parse_args()

    if args.small:
        file = Path("./in_small.txt")
    else:
        file = Path("./in.txt")

    result1 = result2 = 0
    with open(file) as f:
        for line in f:
            line_clean = line.rstrip()
            curr_sol_lights, curr_sol_joltage = min_number_presses(line_clean)
            print(curr_sol_joltage)
            print()
            result1 += curr_sol_lights
            result2 += curr_sol_joltage

    print(f"Result 1: {result1}")
    print(f"Result 2: {result2}")


if __name__ == "__main__":
    main()

#!/usr/bin/env python

QUESTION = 2
directions = {"R": 1, "L": 0}


def gcd(a, b):
    if a == 0:
        return b
    return gcd(b % a, a)


def lcm(a, b):
    return (a / gcd(a, b)) * b


if __name__ == "__main__":
    in_file = "in.txt"

    with open(in_file) as f:
        lines = [line.rstrip() for line in f]

        instructions = list(lines[0])
        print(instructions)

        elements_map = {}
        # Create map with all instructions
        for i in range(2, len(lines)):
            k = lines[i].split(" ")[0]

            v1 = lines[i].split("(")[1].split(",")[0]
            v2 = lines[i].split("(")[1].split(" ")[1].replace(")", "")

            elements_map[k] = (v1, v2)

        if QUESTION == 1:
            current_elem = "AAA"
            i = 0
            steps = 0
            while current_elem != "ZZZ":
                current_elem = elements_map[current_elem][
                    directions[instructions[i]]
                ]
                i = (i + 1) % len(instructions)
                steps += 1

            print("Q1 - steps: ", steps)
        elif QUESTION == 2:
            curr_elem = [k for k in elements_map.keys() if k.endswith("A")]
            print(curr_elem)

            i = [0] * len(curr_elem)
            steps = [0] * len(curr_elem)

            for el in range(len(curr_elem)):
                while not curr_elem[el].endswith("Z"):
                    curr_elem[el] = elements_map[curr_elem[el]][
                        directions[instructions[i[el]]]
                    ]
                    i[el] = (i[el] + 1) % len(instructions)
                    steps[el] += 1
            tot_steps = lcm(steps[0], steps[1])
            for i in range(2, len(steps)):
                tot_steps = lcm(tot_steps, steps[i])
            print(f"Q2 - steps: {tot_steps}")                            

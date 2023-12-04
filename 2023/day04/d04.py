#!/usr/bin/env

import numpy as np


def get_win_mine(line: str):
    """
    Get the winning numbers and the numbers that appear.

    The line structure is:
        Card #: <w1> ... <wN> | <n1> ... <nM>
    """
    splt = line.split(": ")[1].split("|")

    # Need to remove empty strings (if 2 consecutive spaces)
    winning_lst = splt[0].split(" ")
    winning_lst_correct = [int(c) for c in winning_lst if c.isalnum()]
    my_lst = splt[1].split(" ")
    my_lst_correct = [int(c) for c in my_lst if c.isalnum()]

    return winning_lst_correct, my_lst_correct


in_file = "./in.txt"

tot_pts_1 = 0

if __name__ == "__main__":
    with open(in_file) as f:
        lines = [line.rstrip() for line in f]
        tot_lines = len(lines)

        # This list contains the number of times each card is considered
        count_cards = np.ones((tot_lines,), dtype=np.int64)

        for i in range(tot_lines):
            line = lines[i]
            win, mine = get_win_mine(line)

            # The times the current card is counted
            curr_card_count = count_cards[i]

            curr_pts_exp = -1  # + 1 is the count of matching numbers
            for n in mine:
                if n in win:
                    curr_pts_exp += 1

            if curr_pts_exp >= 0:
                pts_current = 2**curr_pts_exp
                tot_pts_1 += pts_current

                n_matching = curr_pts_exp + 1

                lower = min(i + 1, tot_lines - 1)
                upper = min(i + n_matching + 1, tot_lines)
                count_cards[lower:upper] += curr_card_count
                print(
                    f"Card {i+1}: add {curr_card_count} to cards {lower + 1} to\
 {upper} - {n_matching} winning"
                )

    for j in range(len(count_cards)):
        print(f"{count_cards[j]} copies of card {j+1}")

    print(f"Question 1: {tot_pts_1}")
    print(f"question 2: {np.sum(count_cards)}")

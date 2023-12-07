#!/usr/bin/env python

# QUESTION = 1
QUESTION = 2
DEBUG = True

if QUESTION == 1:
    # Ordered according to increasing strength
    cards_types = [
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
        "T",
        "J",
        "Q",
        "K",
        "A",
    ]
elif QUESTION == 2:
    cards_types = [
        "J",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
        "T",
        "Q",
        "K",
        "A",
    ]

hands_types = [
    "High card",
    "One pair",
    "Two pairs",
    "Three of a kind",
    "Full house",
    "Four of a kind",
    "Five of a kind",
]


def evalHandStrength(hand: list):
    """Evaluate the strength of the hand"""
    # Count occurrence of each card:
    occ = {x: 0 for x in cards_types}

    for card in hand:
        occ[card] += 1

    vals = list(occ.values())

    if 5 in vals:
        return 6
    elif 4 in vals:
        return 5
    elif 3 in vals and 2 in vals:
        return 4
    elif 3 in vals:
        return 3
    elif vals.count(2) == 2:
        return 2
    elif 2 in vals:
        return 1
    else:
        return 0


def evalHandStrength_2(hand: list):
    """Evaluate the strength of the hand"""
    # Count occurrence of each card:
    occ = {x: 0 for x in cards_types}

    for card in hand:
        occ[card] += 1

    occ_no_J = {}
    max_tp = "2"
    for tp in cards_types:
        if tp != "J":
            occ_no_J[tp] = occ[tp]
            if occ[tp] >= occ[max_tp]:
                max_tp = tp

    occ_no_J[max_tp] += occ["J"]

    vals = list(occ_no_J.values())

    if 5 in vals:
        return 6
    elif 4 in vals:
        return 5
    elif 3 in vals and 2 in vals:
        return 4
    elif 3 in vals:
        return 3
    elif vals.count(2) == 2:
        return 2
    elif 2 in vals:
        return 1
    else:
        return 0


def orderStrength(hands, pos):
    """
    Given a list of hands (supposedly of the same type), order them by
    strength.

    Recursive!

    Args:
        hands: list of lists consisting of the hands
        pos: index of the current element in the list to be compared

    Output:
        sortedList: list of hands sorted (strongest to weakest)
    """

    # Should ensure all lengths are the same

    sortedList = []  # Output: sorted list (from strongest to weakest)
    for tp in cards_types[::-1]:
        # Isolate hands starting with it
        curr_h = [h for h in hands if h[pos] == tp]

        if len(curr_h) == 1:
            sortedList.append(curr_h[0])
        elif len(curr_h) > 1:
            sub = orderStrength(curr_h, pos + 1)
            for el in sub:
                sortedList.append(el)
        else:
            # Do nothing
            pass

    return sortedList


def rankHands(hands_lst):
    """Rank all hands in the list

    Output:
        final_rank: it contains the indices of the elements in hands_lst sorted
        by decreasing strength
    """
    by_type = {}
    for h in hands_types:
        by_type[h] = []

    for hand in hands_lst:
        if QUESTION == 1:
            by_type[hands_types[evalHandStrength(hand)]].append(hand)
        elif QUESTION == 2:
            by_type[hands_types[evalHandStrength_2(hand)]].append(hand)

    final_rank = []
    # Then, rank all hands of the same type
    for h in hands_types[::-1]:
        for el in orderStrength(by_type[h], 0):
            final_rank.append(hands_lst.index(el))

    return final_rank


if __name__ == "__main__":
    in_file = "in.txt"

    with open(in_file) as f:
        lines = [line.rstrip() for line in f]

        hands = [list(ln.split(" ")[0]) for ln in lines]
        for i in range(len(hands)):
            hands[i] = [str(x) for x in hands[i]]
        bids = [int(ln.split(" ")[1]) for ln in lines]

        if QUESTION == 1:
            final_ranking = rankHands(hands)

            tot_points = 0
            for i in range(len(final_ranking)):
                # Rank = len(
                if DEBUG:
                    print(
                        f"{(len(final_ranking) - i)} times {bids[final_ranking[i]]}"
                    )
                tot_points += (len(final_ranking) - i) * bids[final_ranking[i]]

            print(f"Q1: {tot_points}")
        elif QUESTION == 2:
            final_ranking = rankHands(hands)

            tot_points = 0
            for i in range(len(final_ranking)):
                # Rank = len(
                if DEBUG:
                    print(
                        f"{(len(final_ranking) - i)} times {bids[final_ranking[i]]}"
                    )
                tot_points += (len(final_ranking) - i) * bids[final_ranking[i]]

            print(f"Q2: {tot_points}")

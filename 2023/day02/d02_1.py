tot = {
    "red": 12,
    "green": 13,
    "blue": 14,
}

sol = 0

game = 1
with open("./in_1.txt") as f:
    # Do stuff
    for ln in f:
        games_str = ln.split(": ")[1]
        possible = True
        for ex in games_str.split("; "):
            for col in ex.split(", "):
                n = int(col.split(" ")[0])
                c = col.split(" ")[1].split("\n")[0]

                if n > tot[c]:
                    possible = False

        if possible:
            sol += game

        game += 1

print(sol)

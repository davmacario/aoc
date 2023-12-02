sol = 0

game = 1
with open("./in_1.txt") as f:
    # Do stuff
    for ln in f:
        tot = {
            "red": 0,
            "green": 0,
            "blue": 0,
        }
        games_str = ln.split(": ")[1]
        possible = True
        for ex in games_str.split("; "):
            for col in ex.split(", "):
                n = int(col.split(" ")[0])
                c = col.split(" ")[1].split("\n")[0]

                if n > tot[c]:
                    tot[c] = n

        # Eval power:
        sol += tot["red"] * tot["green"] * tot["blue"]

        game += 1

print(sol)

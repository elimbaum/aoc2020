# Notes for Conway Cubes (AOC 2020 17)

Rules:
> If a cube is active and exactly 2 or 3 of its neighbors are also active, the cube remains active. Otherwise, the cube becomes inactive.
> If a cube is inactive but exactly 3 of its neighbors are active, the cube becomes active. Otherwise, the cube remains inactive.

There are 26 neighbors; 27 including the current cell (3x3x3).

- DEAD and == 3 -> alive (total neighbor count: 3)
- ALIVE and == 2 -> alive (total neighbor count: 3)
- ALIVE and == 3 -> alive (total neighbor count: 4)
- else, dead

Count in full cube. If 3, alive. If 4, keep same. Else, dead.

ALIVE: die if not == 3 or == 4
DEAD:  alive in 3

- This is going to be symmetric in the z direction! Since we start with a flat pattern, there won't be any oppourtunity for asymmetry. Therefore, just compute in one direction and implictly mirror.

Storage: hash table by coordinates? slow-ish lookup, because I need to run 27 lookups.

Maybe I need an "alive list" and everything else is assumed dead. Or both (list + hash). However, I'm not sure how to handle birth case...?

What about have each alive cell check its own dead neighbors? Can keep track of who has been checked. *Any dead cell that is going to become alive must have an alive neighbor*, duh.

Basic algorithm:
```

for cell in alive list:
    count self using hash
    if not (3 or 4):
        dead
        remove from next state list/hash

    for neigh in dead neighbors:
        count neigh
        if 3:
            alive
            add to next state list/hash
```

Starting with 2D version to verify results.
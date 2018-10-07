""" I was interested in modeling a particular pattern of combinations for work.
    For an arbitrary list [a, b, c, d, e, ...], the pattern would progress as
    follows:

    a, a-b, a-b-c, a-b-c-d, a-b-c-d-e, ...
       a-c, a-c-d, a-c-d-e, ...
       a-d, a-d-e, ...
       a-e, ...
       ...

    b, b-c, b-c-d, b-c-d-e, ...
       b-d, b-d-e, ...
       b-e, ...
       ...

    c, c-d, c-d-e, ...
       c-e, ...
       ...

    d, d-e, ...
       ...

    e, ...
       ...

    ...

    I wrote the Python script below to generate this pattern for an arbitrary
    list. But when I went to test it for a 5-element list, I was surprised to
    find that those five elements could generated 25 different combinations
    along the pattern.

    Cursory double-checking revealed that the 25 combinations were composed of
    smaller groups made up of 11, 7, 4, 2, and 1 combination(s).

    As it turns out, 1, 2, 4, 7, 11, ... is the beginning of the quadratic
    mathematic sequence n(n-1)/2 + 1. Put simply, this is a sequence starting
    at 1 with a step size of n. Each round (starting at round 1), the step size
    is added to the starting element to calculate the starting element for the
    next round. In this case, the step size is simply the round number, meaning
    that 1+1 = 2, 2+2 = 4, 4+3 = 7, 7+4 = 11, and so on.

    Or, more generally: elem = round_num + elem where elem is initialized to 1
    and round_num is one-indexed rather than zero-indexed.
"""

combo_list = []
log_types = ["legit", "copy", "fake", "modified", "deleted"]

for current_index in range(len(log_types)):
    copy_list = log_types[current_index + 1:]
    combo = log_types[current_index]
    combo_list.append(combo)
    while len(copy_list) > 0:
        for elem in copy_list:
            combo = "-".join([combo, elem])
            combo_list.append(combo)
        del copy_list[0]
        combo = log_types[current_index]

for combo in combo_list:
    print(combo)

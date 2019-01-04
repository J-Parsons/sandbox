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
    smaller groups made up of 11, 7, 4, 2, and 1 combination(s), which can be
    represented by the quadratic sequence n(n-1)/2 + 1, where n (the number of
    elements) is a positive integer.

    n   f(n)
    1   1       1 + 1 = 2
    2   2       2 + 2 = 4
    3   4       3 + 4 = 7
    4   7       4 + 7 = 11
    5   11      5 + 11 = 16
    ...

    This sequence can be generated with the following pseudo code:

    n = 1
    fn = 1
    while 1:
        fn = fn + n
        n += 1
"""

combo_list = []
letters = ["a", "b", "c", "d", "e"]

for i, letter in enumerate(letters):
    remaining = letters[i+1:]
    combo = letter
    combo_list.append(combo)
    while remaining:
        for sub_letter in remaining:
            combo = combo + sub_letter
            combo_list.append(combo)
        del remaining[0]

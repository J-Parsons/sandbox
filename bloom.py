""" This module provides simple implementations for several variants of Bloom
Filters. A Bloom Filter is a high-speed, space-efficient data structure that
supports probabilistic set-membership tests for a set of items. False positive
matches are possible, but false negatives are not, so a membership query will
return "possibly in the set" or "definitely not in the set". Elements can be
added to a Bloom Filter but not removed, similar to a blockchain.

A basic Bloom Filter uses k hash functions and an m-bit array to represent n
items. Each of the n items is hashed into k different numbers by each hash
function. The corresponding bits in the array (using the numbers are indices)
are then set to 1. Collisions between hash functions are possible, which
explains why membership is probabilistic and why items cannot be removed.

Different Bloom Filter implementations use different hashing strategies.
For example, instead of using 5 different hash functions, some implementations
choose to use double- or triple-hashing, slice the result from a single hash
function, or use a hash function to generate pseudo-random numbers.

Each Bloom Filter uses the variables n, p, m, and k. While these names are not
descriptive, they are consistent with online resources and journal articles.

p is the desired false positive rate.
k is the number of hash functions that will be used to map items to the filter.

Depending on the implementation, n and m take on different meanings.

n is the number of items stored (or desired to store) in the Bloom Filter.
    For a basic Bloom Filter, n represents the number of bits in the filter
    that are occupied (set to 1). For an Invertible Bloom Filter, n is the
    number of key-value pairs currently stored in the filter.

m is the amount of space used by a Bloom Filter. For a basic Bloom Filter,
    m represents the total number of bits available. For an Invertible Bloom
    Filter, m is the total number of key-value pairs that can be stored.

It is not recommended to store [n, m] items, as the false positive rate will
begin to increase rapidly. A full filter is useless and will report *any*
item as present.
"""

import heapq
from math import ceil, exp, log, pow
import random


class _BloomBase:
    def __init__(self, n=None, p=None, m=None, k=None):
        # Use the optimal values for n, p, m, and k if they are not provided.
        # At minimum, we'll need desired values for np, nm, or mp.
        # TODO I doubt this holds for specialized Bloom Filters
        if n and p:
            self.n = n
            self.p = p
            self.m = m or self._optimal_m()
            self.k = k or self._optimal_k()
        elif n and m:
            self.n = n
            self.m = m
            self.k = k or self._optimal_k()
            self.p = self._optimal_p()
        elif m and p:
            self.m = m
            self.p = p
            self.n = n or self._optimal_n()
            self.k = k or self._optimal_k()
        else:
            raise ValueError("You must provide at least np, nm, or mp")

    def _optimal_k(self):
            return round((self.m / self.n) * log(2))

    def _optimal_m(self):
        return ceil((self.n * log(self.p)) / log(1 / pow(2, log(2))))

    def _optimal_n(self):
        return ceil(self.m / -self.k / log(1 - exp(log(self.p) / self.k)))

    def _optimal_p(self):
        return pow(1 - exp(-self.k / (self.m / self.n)), self.k)

    def hashk(self, data):
        """ Returns k random numbers in [1, m] based on data.

        For simplicity, this method uses a single hash function to seed a
        pseudo-random number generator to mimic the use of multiple hash
        functions with minimal collision.

        Args:
            data (any): an object or variable that can be hashed
        """
        hashes = [0] * self.k
        random.seed(data)
        for i in range(self.k):
            hashes[i] = random.randrange(0, self.m)
        return hashes


class BloomFilter(_BloomBase):
    """ A basic Bloom Filter with constant lookup and insertions.

    For simplicity, this implementation uses Python's built-in bytearray
    instead of a proper bit array. For args, see the module-level docstring.
    """
    def __init__(self, n=None, p=None, m=None, k=None):
        super().__init__(n, p, m, k)
        self.bloom = bytearray(self.m)

    def insert(self, data):
        """ Adds a hashable object to the Bloom Filter.

        Args:
            data (any): a hashable object to store in the filter
        """
        for i in self.hashk(data):
            self.bloom[i] = 1

    def get(self, data):
        """ Checks whether a hashable object resides in the Bloom Filter.

        This has a small chance of returning a false positive.

        Args:
            data (any): a hashable object to search for in the filter
        """
        for i in self.hashk(data):
            if not self.bloom[i]:
                return False
        return True


class _Entry:
    def __init__(self):
        self.count = 0
        self.key_sum = 0
        self.value_sum = 0

    def __lt__(self, other):
        return self.count < other.count

    __slots__ = ('count', 'key_sum', 'value_sum')


class InvertibleBloomFilter(_BloomBase):
    """ A Bloom Filter with support for deletion and key-value storage.
    For details, see https://arxiv.org/pdf/1101.2245.pdf.
    For args, see the module-level docstring.
    """
    def __init__(self, n=None, p=None, m=None, k=None):
        super().__init__(n, p, m, k)
        self.bloom = [_Entry() for _ in range(self.m)]

    def insert(self, key, value):
        """ Inserts a key-value pair into the IBF.

        Insertion will always succeed assuming all keys are distinct.
        All data will be converted to strings before being inserted.

        Args:
            key (int): a key to store in the IBF
            value (int) a value referenced by key to store in the IBF
        """
        for i in self.hashk(key):
            self.bloom[i].count += 1
            self.bloom[i].key_sum ^= key
            self.bloom[i].value_sum ^= value

    def delete(self, key, value):
        """ Deletes a key-value pair from the IBF.

        Deletion will always succeed provided the key-value pair was present.

        Args:
            key (int): an key to store in the IBF
            value (int): a value referenced by key to store in the IBF
        """
        for i in self.hashk(key):
            self.bloom[i].count -= 1
            self.bloom[i].key_sum ^= key
            self.bloom[i].value_sum ^= value

    def get(self, key):  # getitem
        """ Probabilistically retrieves a value from the IBF.

        There is a low (but constant) chance that retrieval may fail,
        resulting in a ValueError.

        Args:
            key (int): a key to search for in the filter

        Returns:
            int: the value stored alongside the requested key
            None: if the key could not be found in the filter
        """
        for i in self.hashk(key):
            if self.bloom[i].count == 0:
                return None
            elif self.bloom[i].count == 1:
                if self.bloom[i].key_sum == key:
                    return self.bloom[i].value_sum
                else:
                    return None
        raise ValueError("key-value pair not found")

    def list(self):
        """ Lists all of the key-value pairs stored in the IBF.

        There is a low (but constant) chance that this method will return an
        incomplete list of items.

        Returns:
            (bool, list[(int, int)]): the list is complete if (True, [...]) is
                returned and incomplete if (False, [...]) is returned.
        """
        pairs = []
        bloom = [entry for entry in self.bloom if entry.count >= 1]
        heapq.heapify(bloom)

        # Walk through the priority queue, deleting the current entry and
        # updating the rest of the queue if it has a count of 1.
        while bloom:
            entry = heapq.heappop(bloom)
            if entry.count == 1:
                pairs.append((entry.key_sum, entry.value_sum))
                self.delete(entry.key_sum, entry.value_sum)

        # If there is an entry in the queue with a count greater than 0,
        # then we were only be able to extract a partial list.
        try:
            next(entry for entry in bloom if entry.count > 1)
        except StopIteration:
            return False, pairs
        return True, pairs

# TODO Spectral Bloom Filter, Fault Tolerant Invertible Bloom Filter, Tests

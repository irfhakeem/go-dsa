"""
Time Complexity: O(3^n)
Space Complexity: O(n)
"""
class Solution:
    def generate(self, store: list, n: int, curr: str):
        if len(curr) == n:
            store.append(curr)
            return

        for c in "abc":
            if not curr or curr[len(curr)-1] != c:
                self.generate(store, n, curr+c)

    def getHappyString(self, n: int, k: int) -> str:
        store = []

        self.generate(store, n, "")

        if len(store) < k:
            return ""

        # print(store)

        return store[k-1]

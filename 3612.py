class Solution:
    def proceesStr(self, s: str) -> None:
        result: str = ""
        for idx, value in enumerate(s):
            if value == "*":
                result = result[:-1]
            elif value == "#":
                result = result * 2
            elif value == "%":
                reverse: str = ""
                for i in range(
                    len(result) - 1, -1, -1
                ):  # starts from -1, increment of -1, until -1
                    reverse += result[i]
                result = reverse
            else:
                result += value
        print(result)


if __name__ == "__main__":
    solution = Solution()
    solution.proceesStr("a#b%*")

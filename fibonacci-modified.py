# Link : https://www.hackerrank.com/challenges/fibonacci-modified
# code snippet to solve the fibonacci modified problem on hackerrank

def recurse(current, previous, currentIndex, stopIndex):
    if currentIndex > stopIndex:
        return current
    else:
        value = (current * current) + previous
        currentIndex = currentIndex + 1
        return recurse(value, current, currentIndex, stopIndex)

def __main__():
    input_string = input()
    this = input_string.split()
    first = int(this[0])
    second = int(this[1])
    stopIndex = int(this[2])

    output = recurse(second, first, 3, stopIndex)
    print(output)

__main__()

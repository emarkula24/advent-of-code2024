
import os
import datetime
main_path = os.path.dirname(__file__)
file_path = os.path.join(main_path, 'input.txt')

def part1():
    start = datetime.datetime.now()
    total = 0
    with open(file_path) as f:
        for line in f:
            
            stack = [0]
            line = line.strip()
            line = [int(x) for x in line]
            
            for i, first_number in enumerate(line):
                for j in range(i + 1, len(line)):
                    second_number = line[j]
                    pair = str(first_number) +str(second_number)

                    larger = stack[-1]
                    if int(larger) < int(pair):
                        stack.append(pair)
                        
            total += int(stack.pop())
            
    print(total)
    end = datetime.datetime.now()
    duration = end - start
    print(f'{duration.total_seconds():.3f} milliseconds')

def part2():
    start = datetime.datetime.now()
    total = 0
    with open(file_path) as f:
        for line in f:
            stack = []
            line = line.strip()
            line = [int(x) for x in line]
            n = len(line)
            k = 12
            if k >= n:
                continue
            to_drop = n - k
            for number in line:
                while stack and to_drop > 0 and stack[-1] < number:
                    stack.pop()
                    to_drop -= 1
                stack.append(number)
            
            if to_drop > 0:
                stack = stack[:-to_drop]
                
            result = int("".join(str(x) for x in stack))
            
            total += int(result) 
    print(total)
    end = datetime.datetime.now()
    duration = end - start
    print(f'{duration.total_seconds():.3f} milliseconds')
part2()
part1()
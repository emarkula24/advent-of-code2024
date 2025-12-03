

def part1():
    with open ("input.txt") as f:
        f = f.read().split(',')
   
    
    total = 0
    for line in f:
        start_str, end_str = line.strip().split('-')
        start = int(start_str)
        end = int(end_str)
        if start_str[0] == 0 or end_str[0] == 0:
            continue
        for number in range(start, end + 1):
            number_str = str(number)
            l = len(str(number))
            mid = l // 2
            left, right = number_str[:mid], number_str[mid:]
            if left == right:
                total += number
            
    print(total)

def part2():
    with open ("input.txt") as f:
        f = f.read().split(',')
   
    
    total = 0
    for line in f:
        start_str, end_str = line.strip().split('-')
        start = int(start_str)
        end = int(end_str)
        if start_str[0] == 0 or end_str[0] == 0:
            continue
        for number in range(start, end + 1):
            number_str = str(number)
            l = len(str(number))
            mid = l // 2
            left, right = number_str[:mid], number_str[mid:]
            if left == right:
                total += number
                
            elif number_str in (number_str + number_str)[1:-1]:
                
                total += number
    print(total)




part1()
part2()
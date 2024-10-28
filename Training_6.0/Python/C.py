def find_intervals(table, x) -> list[list[int]]:
    intervals = []
    current_interval = []
    in_interval = False
    for y in range(len(table)):
        if table[y][x] == "#":
            if not in_interval:
                in_interval = True
                current_interval.append(y)
        else:
            if in_interval:
                current_interval.append(y)
                intervals.append(current_interval)
                current_interval = []
                in_interval = False
    else:
        if in_interval:
            current_interval.append(y + 1)
            intervals.append(current_interval)
    return intervals


def table_to_letter(table):
    status = 'find_first_intervals'
    # Сканируем слева направо
    for x in range(len(table[0])):

        # Проходим по столбцу
        intervals = find_intervals(table, x)

        if status == 'find_first_intervals':
            if len(intervals) == 1:
                status = 'find_i'
                check_intervals = intervals
            elif len(intervals) > 1:
                return 'X'
        elif status == 'find_i':
            if len(intervals) == 0:
                status = 'i_end'
            elif len(intervals) == 1:
                if intervals[0] == check_intervals[0]:
                    continue
                elif intervals[0][1] == check_intervals[0][1] and intervals[0][0] > check_intervals[0][0]:
                    status = 'find_l'
                    check_intervals = intervals
                elif intervals[0][0] > check_intervals[0][0] and intervals[0][1] < check_intervals[0][1]:
                    status = 'find_h'
                    left_intervals = check_intervals
                    check_intervals = intervals
                else:
                    return 'X'
            elif len(intervals) == 2:
                if intervals[0][0] == check_intervals[0][0] and intervals[1][1] < check_intervals[0][1]:
                    status = 'find_p'
                    check_intervals = intervals
                elif intervals[0][0] == check_intervals[0][0] and intervals[1][1] == check_intervals[0][1]:
                    status = 'find_c'
                    left_intervals = check_intervals
                    check_intervals = intervals
                else:
                    return 'X'
            else:
                return 'X'
        elif status == 'find_l':
            if len(intervals) == 0:
                status = 'l_end'
            elif len(intervals) == 1:
                if intervals[0] != check_intervals[0]:
                    return 'X'
            else:
                return 'X'
        elif status == 'find_h':
            if len(intervals) == 1:
                if intervals[0] == check_intervals[0]:
                    continue
                elif left_intervals[0] == intervals[0]:
                    status = 'right_h'
                    check_intervals = intervals
                else:
                    return 'X'
            else:
                return 'X'
        elif status == 'right_h':
            if len(intervals) == 0:
                status = 'h_end'
            elif len(intervals) == 1:
                if intervals[0] != check_intervals[0]:
                    return 'X'
            else:
                return 'X'
        elif status == 'find_p':
            if len(intervals) == 2:
                if intervals != check_intervals:
                    return 'X'
            elif len(intervals) == 1:
                if intervals[0][0] == check_intervals[0][0] and intervals[0][1] == check_intervals[1][1]:
                    status = 'p_right'
                    check_intervals = intervals
                else:
                    return 'X'
            else:
                return 'X'
        elif status == 'p_right':
            if len(intervals) == 0:
                status = 'p_end'
            elif len(intervals) == 1:
                if intervals[0] != check_intervals[0]:
                    return 'X'
            else:
                return 'X'
        elif status == 'find_c':
            if len(intervals) == 0:
                status = 'c_end'
            elif len(intervals) == 2:
                if intervals != check_intervals:
                    return 'X'
            elif len(intervals) == 1:
                if intervals == left_intervals:
                    status = 'find_o'
                    check_intervals = intervals
                else:
                    return 'X'
            else:
                return 'X'
        elif status == 'find_o':
            if len(intervals) == 0:
                status = 'o_end'
            elif len(intervals) == 1:
                if intervals == check_intervals:
                    continue
                else:
                    return 'X'
            else:
                return 'X'
        elif status == 'i_end' or status == 'l_end' or status == 'h_end' or status == 'p_end' or status == 'c_end' or status == 'o_end':
            if len(intervals) != 0:
                return 'X'
    else:
        if status == 'i_end' or status == 'find_i':
            return 'I'
        elif status == 'l_end' or status == 'find_l':
            return 'L'
        elif status == 'h_end' or status == 'right_h':
            return 'H'
        elif status == 'p_end' or status == 'p_right':
            return 'P'
        elif status == 'c_end' or status == 'find_c':
            return 'C'
        elif status == 'o_end' or status == 'find_o':
            return 'O'
        else:
            return 'X'


n = int(input())
table = [input() for _ in range(n)]
print(table_to_letter(table))

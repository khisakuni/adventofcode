const std = @import("std");
const input = @embedFile("input.txt");

pub fn main() !void {
    var it = std.mem.tokenizeScalar(u8, input, '\n');

    var firstList = std.ArrayList(i32).init(std.heap.page_allocator);
    var secondList = std.ArrayList(i32).init(std.heap.page_allocator);
    while (it.next()) |token| {
        var iter = std.mem.splitSequence(u8, token, "   ");

        const firstNum = try std.fmt.parseInt(i32, iter.next().?, 10);
        try firstList.append(firstNum);

        const secondNum = try std.fmt.parseInt(i32, iter.next().?, 10);
        try secondList.append(secondNum);
    }

    const f = try firstList.toOwnedSlice();
    defer std.heap.page_allocator.free(f);

    const s = try secondList.toOwnedSlice();
    defer std.heap.page_allocator.free(s);

    std.mem.sort(i32, f, {}, comptime std.sort.asc(i32));
    std.mem.sort(i32, s, {}, comptime std.sort.asc(i32));

    var map = std.AutoHashMap(i32, i8).init(std.heap.page_allocator);
    defer map.deinit();

    var total: i32 = 0;
    var index: usize = 0;
    for (f) |num| {
        const sNum = s[index];
        var diff = num - sNum;
        if (diff < 0) {
            diff *= -1;
        }
        total += diff;
        index += 1;

        const freq = map.get(sNum) orelse 0;
        try map.put(sNum, freq + 1);
    }

    var part2Total: i32 = 0;
    for (f) |num| {
        const freq = map.get(num) orelse 0;
        part2Total += num * freq;
    }

    std.debug.print("Part 1: {d}\nPart 2: {d}", .{ total, part2Total });
}

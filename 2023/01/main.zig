const std = @import("std");

pub fn main() !void {
    // Open input file.
    var file = try std.fs.cwd().openFile("input.txt", .{ .mode = .read_only });
    defer file.close();

    // Get file info.
    var file_stat = try file.stat();

    // Get file size.
    const file_size = file_stat.size;

    // Initialize allocator.
    const allocator = std.heap.page_allocator;

    // Read file contents.
    const buffer = try file.readToEndAlloc(allocator, file_size);
    defer allocator.free(buffer);

    var map = std.StringHashMap(u8).init(allocator);
    try map.put("1", 1);
    try map.put("one", 1);
    try map.put("2", 2);
    try map.put("two", 2);
    try map.put("3", 3);
    try map.put("three", 3);
    try map.put("4", 4);
    try map.put("four", 4);
    try map.put("5", 5);
    try map.put("five", 5);
    try map.put("6", 6);
    try map.put("six", 6);
    try map.put("7", 7);
    try map.put("seven", 7);
    try map.put("8", 8);
    try map.put("eight", 8);
    try map.put("9", 9);
    try map.put("nine", 9);
    defer map.deinit();

    const candidates = [_][]const u8{
        "1",
        "one",
        "2",
        "two",
        "3",
        "three",
        "4",
        "four",
        "5",
        "five",
        "6",
        "six",
        "7",
        "seven",
        "8",
        "eight",
        "9",
        "nine",
    };

    var line = std.ArrayList(u8).init(allocator);
    defer line.deinit();
    var lines = std.ArrayList(std.ArrayList(u8)).init(allocator);
    defer lines.deinit();

    for (buffer) |character| {
        if (character == '\n') {
            try lines.append(line);
            line = std.ArrayList(u8).init(allocator);
            continue;
        }

        try line.append(character);
    }

    var total: u32 = 0;
    for (lines.items) |l| {
        std.debug.print("line: {s}\n", .{l.items});
        var list = std.ArrayList([]const u8).init(allocator);
        defer list.deinit();
        for (0..l.items.len) |i| {
            for (candidates) |candidate| {
                if (l.items.len - i >= candidate.len and std.mem.eql(u8, candidate, l.items[i .. i + candidate.len])) {
                    std.debug.print("c: {s}: {d}\n", .{ candidate, l.items.len });
                    try list.append(candidate);
                }
            }
        }

        var first = map.get(list.items[0]).?;
        var last = map.get(list.items[list.items.len - 1]).?;
        total += first * 10 + last;
    }

    std.debug.print("Total {d}\n", .{total});
}

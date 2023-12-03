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
    defer map.deinit();

    try map.put("red", 12);
    try map.put("green", 13);
    try map.put("blue", 14);

    var iter = std.mem.split(u8, buffer, "\n");
    var total: i16 = 0;
    while (iter.next()) |line| {
        // std.debug.print("line: {s}\n", .{line});
        var line_iter = std.mem.splitAny(u8, line, ":;");

        var game_num_parts_iter = std.mem.splitBackwardsAny(u8, line_iter.next().?, " ");
        var num = try std.fmt.parseInt(i16, game_num_parts_iter.next().?, 10);

        std.debug.print(" >> {d}\n", .{num});

        var isValid = true;
        while (line_iter.next()) |parts| {
            var it = std.mem.splitAny(u8, parts, ",");
            while (it.next()) |part| {
                var i = std.mem.splitAny(u8, part, " ");
                var colorNum: i16 = 0;
                // var mins = std.StringHashMap(i16).init(allocator);
                // defer mins.deinit();

                while (i.next()) |p| {
                    if (std.mem.eql(u8, p, "")) {
                        continue;
                    }

                    var n = std.fmt.parseInt(i16, p, 10) catch {
                        const limit = map.get(p).?;
                        if (colorNum > limit) {
                            isValid = false;
                        }

                        std.debug.print(">>>>>>> {s} {d}, limit: {d}\n", .{ p, colorNum, limit });
                        continue;
                    };

                    colorNum = n;
                }
            }
        }

        if (isValid) {
            total += num;
        }
        isValid = true;
    }

    std.debug.print("total: {d}\n", .{total});
}

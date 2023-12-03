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

    var iter = std.mem.split(u8, buffer, "\n");
    var total: i32 = 0;
    while (iter.next()) |line| {
        if (std.mem.eql(u8, line, "")) {
            continue;
        }
        var line_iter = std.mem.splitAny(u8, line, ":;");

        var game_num_parts_iter = std.mem.splitBackwardsAny(u8, line_iter.next().?, " ");
        var num = try std.fmt.parseInt(i32, game_num_parts_iter.next().?, 10);

        std.debug.print("game: {d}\n", .{num});

        var mins = std.StringHashMap(i32).init(allocator);
        defer mins.deinit();

        while (line_iter.next()) |parts| {
            var it = std.mem.splitAny(u8, parts, ",");

            while (it.next()) |part| {
                var i = std.mem.splitAny(u8, part, " ");
                var colorNum: i32 = 0;

                while (i.next()) |p| {
                    if (std.mem.eql(u8, p, "")) {
                        continue;
                    }

                    var n = std.fmt.parseInt(i32, p, 10) catch {
                        const currentMin = mins.get(p) orelse 0;
                        if (colorNum > currentMin) {
                            try mins.put(p, colorNum);
                        }

                        continue;
                    };

                    colorNum = n;
                }
            }
        }

        var mins_iter = mins.iterator();
        var pow: i32 = 1;
        while (mins_iter.next()) |entry| {
            pow *= entry.value_ptr.*;
        }

        total += pow;
    }

    std.debug.print("total: {d}\n", .{total});
}

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

    var total: i32 = 0;
    var lines_iter = std.mem.splitAny(u8, buffer, "\n");
    while (lines_iter.next()) |line| {
        if (line.len == 0) {
            continue;
        }
        var line_iter = std.mem.splitAny(u8, line, "|");
        var win = line_iter.next().?;
        var have = line_iter.next().?;

        // std.debug.print("winning: {s}, have: {s}\n", .{ win, have });
        var win_set = std.AutoHashMap(u8, void).init(allocator);
        defer win_set.deinit();

        // Build up set of winning numbers.
        var i: usize = win.len;
        var num_seq_start: usize = 0;
        var is_num = false;
        while (i >= 0) {
            i -= 1;
            if (win[i] == ':') {
                break;
            }
            if (std.ascii.isDigit(win[i])) {
                if (!is_num) {
                    is_num = true;
                    num_seq_start = i;
                }
            } else {
                if (is_num) {
                    is_num = false;
                    const num = try std.fmt.parseInt(u8, win[i + 1 .. num_seq_start + 1], 10);
                    try win_set.put(num, undefined);
                }
            }
        }

        // Start checking numbers we have against winning numbers.
        i = have.len;
        is_num = false;
        var count: i16 = 0;
        while (i > 0) {
            i -= 1;
            if (std.ascii.isDigit(have[i])) {
                if (!is_num) {
                    is_num = true;
                    num_seq_start = i;
                }
            } else {
                if (is_num) {
                    is_num = false;
                    const num = try std.fmt.parseInt(u8, have[i + 1 .. num_seq_start + 1], 10);
                    // std.debug.print("have: {d}\n", .{num});
                    if (win_set.get(num)) |_| {
                        count += 1;
                    }
                }
            }
        }
        if (count > 0) {
            const score = std.math.pow(i32, 2, count - 1);
            total += score;
            // std.debug.print("count: {d}: {d}\n", .{ count, score });
        }
    }

    std.debug.print("total: {d}\n", .{total});
}

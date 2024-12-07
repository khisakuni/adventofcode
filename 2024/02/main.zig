const std = @import("std");
const input = @embedFile("input.txt");

pub fn main() !void {
    var it = std.mem.tokenizeScalar(u8, input, '\n');
    var safe: i32 = 0;
    outer: while (it.next()) |report| {
        var iter = std.mem.splitSequence(u8, report, " ");

        var list = std.ArrayList(i8).init(std.heap.page_allocator);
        defer list.deinit();

        while (iter.next()) |level| {
            const num = try std.fmt.parseInt(i8, level, 10);
            try list.append(num);
        }

        var skip: i32 = -1;
        skip: while (skip < list.items.len) {
            defer skip += 1;
            var start: usize = 0;
            if (skip == 0) {
                start = 1;
            }

            var prev: i8 = list.items[start];
            var asc: ?bool = null;

            inner: for (start + 1..list.items.len) |i| {
                if (i == skip) {
                    continue :inner;
                }

                const delta: i8 = list.items[i] - prev;
                if (asc == null) {
                    asc = delta > 0;
                }

                if ((delta == 0) or
                    (delta > 3) or
                    (delta < -3) or
                    (asc.? and delta < 0) or
                    (!asc.? and delta > 0))
                {
                    continue :skip;
                }

                prev = list.items[i];
            }
            safe += 1;
            continue :outer;
        }
    }

    std.debug.print("Answer: {d}\n", .{safe});
}

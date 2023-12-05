const std = @import("std");

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{ .mode = .read_only });
    defer file.close();

    var file_stat = try file.stat();

    const file_size = file_stat.size;

    const allocator = std.heap.page_allocator;

    const buffer = try file.readToEndAlloc(allocator, file_size);
    defer allocator.free(buffer);

    var seed_ranges = std.ArrayList(i64).init(allocator);
    defer seed_ranges.deinit();

    var seeds_delim = "seeds:";

    var transforms = std.ArrayList(std.ArrayList([3]i64)).init(allocator);
    defer transforms.deinit();

    var transform = std.ArrayList([3]i64).init(allocator);
    defer transform.deinit();

    var lines_iter = std.mem.splitAny(u8, buffer, "\n");
    while (lines_iter.next()) |line| {
        // Parse seeds.
        if (line.len > seeds_delim.len and std.mem.eql(u8, line[0..seeds_delim.len], seeds_delim)) {
            var seeds_words = line[seeds_delim.len + 1 .. line.len];
            var seeds_iter = std.mem.splitAny(u8, seeds_words, " ");
            while (seeds_iter.next()) |seed_word| {
                var seed = try std.fmt.parseUnsigned(i64, seed_word, 10);
                try seed_ranges.append(seed);
            }
            continue;
        }

        // New line between maps, or seed and map sections.
        if (line.len == 0) {
            if (transform.items.len > 0) {
                try transforms.append(transform);
                transform = std.ArrayList([3]i64).init(allocator);
            }
            continue;
        }

        if (!std.ascii.isDigit(line[0])) {
            continue;
        }

        var nums = std.mem.splitAny(u8, line, " ");
        var dest_start = try std.fmt.parseInt(i64, nums.next().?, 10);
        var src_start = try std.fmt.parseInt(i64, nums.next().?, 10);
        var offset = try std.fmt.parseInt(i64, nums.next().?, 10);
        var list = [3]i64{ dest_start, src_start, offset };
        try transform.append(list);
    }

    var deduped_ranges = std.ArrayList(i64).init(allocator);
    defer deduped_ranges.deinit();

    var min: i64 = -1;
    var i: usize = 0;
    var ranges = std.ArrayList([2]i64).init(allocator);
    defer ranges.deinit();

    while (i + 1 < seed_ranges.items.len) {
        try ranges.append([2]i64{ seed_ranges.items[i], seed_ranges.items[i] + seed_ranges.items[i + 1] });
        i += 2;
    }

    for (ranges.items) |range| {
        var current = std.ArrayList([2]i64).init(allocator);
        defer current.deinit();

        try current.append(range);

        for (transforms.items) |t| {
            // std.debug.print("transform: {d}\n", .{index});
            var next = std.ArrayList([2]i64).init(allocator);
            while (current.items.len > 0) {
                var current_range = current.pop();
                var added = false;

                for (t.items) |item| {
                    var dest_start = item[0];
                    var src_start = item[1];
                    var offset = item[2];

                    const start = current_range[0];
                    const end = current_range[1];

                    const diff = dest_start - src_start;

                    // Complete overlap.
                    if (start >= src_start and end <= src_start + offset) {
                        // std.debug.print("complete overlap\n", .{});
                        try next.append([2]i64{ start + diff, end + diff });
                        added = true;
                        break;
                    }

                    // Front overlap.
                    if (start > src_start and start < src_start + offset and end > src_start + offset) {
                        // std.debug.print("front overlap\n", .{});
                        // Add overlapping part to next.
                        try next.append([2]i64{ start + diff, src_start + offset + diff });
                        added = true;

                        // Add non-overlapping part for further processing.
                        try current.append([2]i64{ src_start + offset, end });
                        break;
                    }

                    // Back overlap.
                    if (start < src_start and end > src_start and end < src_start + offset) {
                        // std.debug.print("back overlap\n", .{});
                        // Add overlapping part to next.
                        try next.append([2]i64{ src_start + diff, end + diff });
                        added = true;

                        // Add non-overlapping part for further processing.
                        try current.append([2]i64{ start, src_start });
                        break;
                    }

                    // Split.
                    if (start < src_start and end > src_start + offset) {
                        // std.debug.print("split overlap\n", .{});
                        try next.append([2]i64{ src_start + diff, src_start + offset + diff });

                        try current.append([2]i64{ start, src_start });
                        try current.append([2]i64{ src_start + offset, end });
                    }
                }

                if (!added) {
                    try next.append(current_range);
                }
            }

            current = next;
        }

        for (current.items) |r| {
            if (min < 0) {
                min = r[0];
            }

            if (r[0] < min) {
                min = r[0];
            }
        }
    }

    std.debug.print("min: {d}\n", .{min});
}

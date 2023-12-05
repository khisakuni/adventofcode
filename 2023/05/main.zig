const std = @import("std");

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{ .mode = .read_only });
    defer file.close();

    var file_stat = try file.stat();

    const file_size = file_stat.size;

    const allocator = std.heap.page_allocator;

    const buffer = try file.readToEndAlloc(allocator, file_size);
    defer allocator.free(buffer);

    var seeds = std.ArrayList(i64).init(allocator);
    defer seeds.deinit();

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
                var seed = try std.fmt.parseInt(i64, seed_word, 10);
                try seeds.append(seed);
            }
            continue;
        }

        // New line between maps, or seed and map sections.
        if (line.len == 0) {
            if (transform.items.len > 0) {
                // std.debug.print("appending! {d}\n", .{transform.items.len});
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
        // std.debug.print("{d}, {d}, {d}\n", .{ dest_start, src_start, offset });
        var list = [3]i64{ dest_start, src_start, offset };
        try transform.append(list);
    }

    var min: i64 = seeds.items[0];
    for (seeds.items[1..]) |seed| {
        var next: i64 = seed;

        for (transforms.items) |t| {
            for (t.items) |item| {
                var dest_start = item[0];
                var src_start = item[1];
                var offset = item[2];

                if (next >= src_start and next <= src_start + offset) {
                    next = dest_start + (next - src_start);
                    break;
                }
            }
        }

        // std.debug.print("{d} -> {d}\n", .{ seed, next });
        if (next < min) {
            min = next;
        }
    }

    std.debug.print("min: {d}\n", .{min});
}

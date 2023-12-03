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

    var coords_to_id = std.AutoHashMap(i16, std.AutoHashMap(i16, i16)).init(allocator);
    defer coords_to_id.deinit();

    var sym_coords = std.AutoHashMap(i16, std.AutoHashMap(i16, void)).init(allocator);
    defer sym_coords.deinit();

    var nums = std.ArrayList(i16).init(allocator);
    defer nums.deinit();

    var is_num = false;
    var id: i16 = 0;
    var y: i16 = 0;
    var x: i16 = 0;
    var maxX: i16 = 0;
    var num_seq = std.ArrayList(u8).init(allocator);
    defer num_seq.deinit();

    for (buffer) |c| {
        // std.debug.print("x: {d}, max: {d}\n", .{ x, maxX });
        if (x - 1 > maxX) {
            maxX = x - 1;
        }
        if (isNumber(c)) {
            is_num = true;
            try num_seq.append(c);
            var entry = try coords_to_id.getOrPut(y);
            if (!entry.found_existing) {
                entry.value_ptr.* = std.AutoHashMap(i16, i16).init(allocator);
            }

            try entry.value_ptr.*.put(x, id);
        } else {
            if (is_num) {
                var num = try std.fmt.parseInt(i16, num_seq.items, 10);
                try nums.append(num);
                is_num = false;
                num_seq = std.ArrayList(u8).init(allocator);
                id += 1;
            }

            if (isSymbol(c)) {
                // std.debug.print(">>>> {c}, {d}:{d}\n", .{ c, y, x });
                var entry = try sym_coords.getOrPut(y);
                if (!entry.found_existing) {
                    entry.value_ptr.* = std.AutoHashMap(i16, void).init(allocator);
                }

                try entry.value_ptr.*.put(x, {});
            }

            if (isNewLine(c)) {
                y += 1;
                x = 0;
                continue;
            }
        }

        x += 1;
    }

    // for (0.., nums.items) |i, num| {
    //     std.debug.print("{d} -> {d}\n", .{ i, num });
    // }

    var sum: i32 = 0;
    var iter = sym_coords.iterator();
    var set = std.AutoHashMap(i16, void).init(allocator);
    defer set.deinit();
    while (iter.next()) |entry| {
        var sy = entry.key_ptr.*;
        var x_vals = entry.value_ptr.*;
        var x_iter = x_vals.iterator();
        while (x_iter.next()) |x_entry| {
            var sx = x_entry.key_ptr.*;
            // std.debug.print("y: {d} x:{d}\n", .{ sy, sx });

            // Top left
            if (sy > 0 and sx > 0) {
                if (coords_to_id.get(sy - 1)) |*e| {
                    if (e.get(sx - 1)) |num_id| {
                        // std.debug.print("found num: {d}:{d} -> {d}\n", .{ sy, sx, nums.items[@as(usize, @intCast(num_id))] });
                        try set.put(num_id, {});
                    }
                }
            }
            // Top
            if (sy > 0) {
                if (coords_to_id.get(sy - 1)) |*e| {
                    if (e.get(sx)) |num_id| {
                        // std.debug.print("found num: {d}:{d} -> {d}\n", .{ sy, sx, nums.items[@as(usize, @intCast(num_id))] });
                        try set.put(num_id, {});
                    }
                }
            }
            // Top right
            if (sy > 0 and sx < maxX) {
                if (coords_to_id.get(sy - 1)) |*e| {
                    if (e.get(sx + 1)) |num_id| {
                        // std.debug.print("found num: {d}:{d} -> {d}\n", .{ sy, sx, nums.items[@as(usize, @intCast(num_id))] });
                        try set.put(num_id, {});
                    }
                }
            }
            // Left
            if (sx > 0) {
                if (coords_to_id.get(sy)) |*e| {
                    if (e.get(sx - 1)) |num_id| {
                        // std.debug.print("found num: {d}:{d} -> {d}\n", .{ sy, sx, nums.items[@as(usize, @intCast(num_id))] });
                        try set.put(num_id, {});
                    }
                }
            }
            // Right
            if (sx < maxX) {
                if (coords_to_id.get(sy)) |*e| {
                    if (e.get(sx + 1)) |num_id| {
                        // std.debug.print("found num: {d}:{d} -> {d}\n", .{ sy, sx, nums.items[@as(usize, @intCast(num_id))] });
                        try set.put(num_id, {});
                    }
                }
            }
            // Bottom left
            if (sy < y and sx > 0) {
                if (coords_to_id.get(sy + 1)) |*e| {
                    if (e.get(sx - 1)) |num_id| {
                        // std.debug.print("found num: {d}:{d} -> {d}\n", .{ sy, sx, nums.items[@as(usize, @intCast(num_id))] });
                        try set.put(num_id, {});
                    }
                }
            }
            // Bottom
            if (sy < y) {
                if (coords_to_id.get(sy + 1)) |*e| {
                    if (e.get(sx)) |num_id| {
                        // std.debug.print("found num: {d}:{d} -> {d}\n", .{ sy, sx, nums.items[@as(usize, @intCast(num_id))] });
                        try set.put(num_id, {});
                    }
                }
            }
            // Bottom right
            if (sy < y and sx > 0) {
                if (coords_to_id.get(sy + 1)) |*e| {
                    if (e.get(sx + 1)) |num_id| {
                        // std.debug.print("found num: {d}:{d} -> {d}\n", .{ sy, sx, nums.items[@as(usize, @intCast(num_id))] });
                        // try set.put(nums.items[@as(usize, @intCast(num_id))], {});
                        try set.put(num_id, {});
                    }
                }
            }
        }
    }

    // 5 + 7 + 4 + 13 + 9 + 15 + 9
    var i = set.iterator();
    while (i.next()) |e| {

        // try set.put(nums.items[@as(usize, @intCast(num_id))], {});
        var n = nums.items[@as(usize, @intCast(e.key_ptr.*))];
        // std.debug.print("num!: {d}\n", .{n});
        sum += n;
    }

    std.debug.print("sum: {d}\n", .{sum});
}

fn isNumber(c: u8) bool {
    return c >= 48 and c <= 57;
}

fn isNewLine(c: u8) bool {
    return c == '\n';
}

fn isSymbol(c: u8) bool {
    return !isNumber(c) and !isNewLine(c) and c != '.';
}

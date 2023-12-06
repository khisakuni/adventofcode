const std = @import("std");

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{ .mode = .read_only });
    defer file.close();

    var file_stat = try file.stat();

    const file_size = file_stat.size;

    const allocator = std.heap.page_allocator;

    const buffer = try file.readToEndAlloc(allocator, file_size);
    defer allocator.free(buffer);

    var times = std.ArrayList(i32).init(allocator);
    defer times.deinit();

    var dists = std.ArrayList(i32).init(allocator);
    defer dists.deinit();

    var lines_iter = std.mem.splitAny(u8, buffer, "\n");
    var i: u8 = 0;
    while (lines_iter.next()) |line| {
        var line_iter = std.mem.splitAny(u8, line, " ");
        while (line_iter.next()) |word| {
            var num = std.fmt.parseInt(i32, word, 10) catch {
                continue;
            };
            if (i == 0) {
                try times.append(num);
            }
            if (i == 1) {
                try dists.append(num);
            }
        }
        i += 1;
    }

    var total: f32 = 1;
    for (0.., times.items) |index, time| {
        const a: f32 = -1;
        const b: f32 = @floatFromInt(time);
        const c: f32 = @floatFromInt(-1 * dists.items[index]);

        const res = try quadratic(a, b, c);
        const min = @ceil(res[0]);
        const max = @floor(res[1]);
        total *= max - min + 1;
    }
    std.debug.print("answer: {d}\n", .{total});
}

fn quadratic(a: f32, b: f32, c: f32) ![2]f32 {
    const d: f32 = b * b - 4 * a * c;
    const d_sqrt: f32 = @sqrt(d);
    const root1 = (-b + d_sqrt) / (2 * a);
    const root2 = (-b - d_sqrt) / (2 * a);
    return [2]f32{ root1, root2 };
}

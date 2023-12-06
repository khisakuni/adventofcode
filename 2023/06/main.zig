const std = @import("std");

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{ .mode = .read_only });
    defer file.close();

    var file_stat = try file.stat();

    const file_size = file_stat.size;

    const allocator = std.heap.page_allocator;

    const buffer = try file.readToEndAlloc(allocator, file_size);
    defer allocator.free(buffer);

    var times = std.ArrayList(u8).init(allocator);
    defer times.deinit();

    var dists = std.ArrayList(u8).init(allocator);
    defer dists.deinit();

    var lines_iter = std.mem.splitAny(u8, buffer, "\n");
    var i: u8 = 0;
    while (lines_iter.next()) |line| {
        for (line) |char| {
            if (std.ascii.isDigit(char)) {
                if (i == 0) {
                    try times.append(char);
                }
                if (i == 1) {
                    try dists.append(char);
                }
            }
        }
        i += 1;
    }

    const time = try std.fmt.parseInt(i64, times.items, 10);
    const dist = try std.fmt.parseInt(i64, dists.items, 10);

    const a: f64 = -1;
    const b: f64 = @floatFromInt(time);
    const c: f64 = @floatFromInt(-1 * dist);

    const res = try quadratic(a, b, c);
    const min = @ceil(res[0]);
    const max = @floor(res[1]);

    std.debug.print("answer: {d}\n", .{max - min + 1});
}

fn quadratic(a: f64, b: f64, c: f64) ![2]f64 {
    const d: f64 = b * b - 4 * a * c;
    const d_sqrt: f64 = @sqrt(d);
    const root1 = (-b + d_sqrt) / (2 * a);
    const root2 = (-b - d_sqrt) / (2 * a);
    return [2]f64{ root1, root2 };
}

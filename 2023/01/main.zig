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

  var map = std.AutoHashMap(u8, u8).init(allocator);
  try map.put('0', 0);
  try map.put('1', 1);
  try map.put('2', 2);
  try map.put('3', 3);
  try map.put('4', 4);
  try map.put('5', 5);
  try map.put('6', 6);
  try map.put('7', 7);
  try map.put('8', 8);
  try map.put('9', 9);


  var total: u32 = 0;
  var is_new_line = true;
  var first: u8 = 0;
  var last: u8 = 0;
  for (buffer) |character| {
    const num = map.get(character);
    if (num) |value| {
      if (is_new_line) {
        first = value;
        is_new_line = false;
      }
      last = value;
    }

    if (character == '\n') {
      is_new_line = true;
      const calibration_value = (first * 10) + last;
      total += calibration_value;
    }
  }

  std.debug.print("Total {d}\n", .{total});
}

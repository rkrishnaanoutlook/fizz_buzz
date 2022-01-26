const fizzbuzz = require("./fizzbuzz"); // or wherever you have your fizzbuzz file

test("fizzbuzz function exists", () => {
  expect(typeof fizzbuzz).toEqual("function");
});

test('for multiples of three print "fizz" instead of the number', () => {
  expect(fizzbuzz(2)).not.toEqual(expect.arrayContaining(["fizz"]));
  expect(fizzbuzz(3)).toEqual(expect.arrayContaining(["fizz"]));
  expect(fizzbuzz(6)).toEqual([1, 2, "fizz", 4, "buzz", "fizz"]);
});

test('for multiples of five print "buzz" instead of the number', () => {
  expect(fizzbuzz(3)).not.toContain("buzz");
  expect(fizzbuzz(5)).toContain("buzz");
  expect(fizzbuzz(10)).toEqual([
    1,
    2,
    "fizz",
    4,
    "buzz",
    "fizz",
    7,
    8,
    "fizz",
    "buzz",
  ]);
});

test('for multiples of both three and five print "fizzbuzz”', () => {
  expect(fizzbuzz(10)).not.toContain("fizzbuzz");
  expect(fizzbuzz(15)).toContain("fizzbuzz");
  expect(fizzbuzz(15)).toEqual([
    1,
    2,
    "fizz",
    4,
    "buzz",
    "fizz",
    7,
    8,
    "fizz",
    "buzz",
    11,
    "fizz",
    13,
    14,
    "fizzbuzz",
  ]);
});
